package slots

import (
	"fmt"
	"log/slog"
)

const (
	DummyGuildID = "000000000000000000"
)

// SlotMachine represents a slot machine with a lookup table, payout table, and symbol table.
type SlotMachine struct {
	LookupTable LookupTable
	PayoutTable PayoutTable
}

// Option is a function that can be used to modify the slot machine's configuration.
type Option func(sm *SlotMachine)

// newSlotMachine creates a new instance of the SlotMachine with initialized lookup table, payout table, and symbol table.
func NewSlotMachine(opts ...Option) *SlotMachine {
	slotMachine := &SlotMachine{
		LookupTable: defaultLookupTable,
		PayoutTable: defaultPayoutTable,
	}
	slotMachine.Apply(opts)

	return slotMachine
}

// Apply applies the given Option(s) to the SlotMachine.
func (sm *SlotMachine) Apply(opts []Option) {
	for _, opt := range opts {
		opt(sm)
	}
}

// WithLookupTable sets the lookup table for the slot machine.
func WithLookupTable(lookupTable LookupTable) Option {
	return func(sm *SlotMachine) {
		sm.LookupTable = lookupTable
	}
}

// WithPayoutTable sets the payout table for the slot machine.
func WithPayoutTable(payoutTable PayoutTable) Option {
	return func(sm *SlotMachine) {
		sm.PayoutTable = payoutTable
	}
}

// Spin represents the result of a spin in the slot machine game, including the winning index and the symbols displayed.
// The spin contains multiple rows of symbols, with the winning row indicated by Payline. THe multiple rows are used
// to create the multiple display lines.
type SpinResult struct {
	TopLine    []string
	Payline    []string
	BottomLine []string
	Bet        int
	Payout     int
	Message    string
}

// String returns a string representation of the Spin.
func (s *SpinResult) String() string {
	return fmt.Sprintf("Spin{Payline: %v, TopLine: %v, BottomLine: %v, Bet: %d, Payout: %d}", s.Payline, s.TopLine, s.BottomLine, s.Bet, s.Payout)
}

// Spin simulates a spin of the slot machine with the given bet amount and returns the result of the spin,
// including the payline, previous line, next line, bet amount, and payout amount.
func (sm *SlotMachine) Spin(bet int) *SpinResult {
	paylineIndices, payline := sm.LookupTable.GetPaylineSpin()
	previousIndices, previousLine := sm.LookupTable.GetPreviousSpin(paylineIndices)
	_, nextLine := sm.LookupTable.GetNextSpin(paylineIndices, previousIndices)
	payoutAmount := sm.PayoutTable.GetPayoutAmount(bet, payline)

	spinResult := &SpinResult{
		Payline:    payline,
		BottomLine: previousLine,
		TopLine:    nextLine,
		Bet:        bet,
		Payout:     payoutAmount,
	}

	slog.Debug("slot machine spin result",
		slog.Int("bet", spinResult.Bet),
		slog.Int("payout", spinResult.Payout),
	)

	return spinResult
}
