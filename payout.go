package slots

import (
	"strconv"
	"strings"
)

const (
	PAYOUT_FILE_NAME = "payout"
	Any              = "any"
	AnySeven         = "any 7"
	AnyBar           = "any bar"
	Blank            = "blank"
	MatchingNonBlank = "matching non-blank"
	AnyRed           = "any red"
	AnyWhite         = "any white"
	AnyBlue          = "any blue"
)

var (
	defaultPayoutTable = PayoutTable{
		{
			Win: []string{
				"red 7",
				"white 7",
				"blue 7",
			},
			Bet:     1,
			Payout:  2400,
			Message: "Jackpot!",
		},
		{
			Win: []string{
				"red 7",
				"red 7",
				"red 7",
			},
			Bet:     1,
			Payout:  1200,
			Message: "Three red 7s!",
		},
		{
			Win: []string{
				"white 7",
				"white 7",
				"white 7",
			},
			Bet:     1,
			Payout:  200,
			Message: "Three white 7s!",
		},
		{
			Win: []string{
				"blue 7",
				"blue 7",
				"blue 7",
			},
			Bet:     1,
			Payout:  150,
			Message: "Three blue 7s!",
		},
		{
			Win: []string{
				"any 7",
				"any 7",
				"any 7",
			},
			Bet:     1,
			Payout:  50,
			Message: "Three 7s!",
		},
		{
			Win: []string{
				"1 bar",
				"2 bar",
				"3 bar",
			},
			Bet:     1,
			Payout:  50,
			Message: "Bar 1, bar 2, bar 3!",
		},
		{
			Win: []string{
				"3 bar",
				"3 bar",
				"3 bar",
			},
			Bet:     1,
			Payout:  40,
			Message: "Three 3 bars!",
		},
		{
			Win: []string{
				"2 bar",
				"2 bar",
				"2 bar",
			},
			Bet:     1,
			Payout:  25,
			Message: "Three 2 bars!",
		},
		{
			Win: []string{
				"any red",
				"any white",
				"any blue",
			},
			Bet:     1,
			Payout:  20,
			Message: "Red, White, and Blue!",
		},
		{
			Win: []string{
				"1 bar",
				"1 bar",
				"1 bar",
			},
			Bet:     1,
			Payout:  10,
			Message: "All 1 bars!",
		},
		{
			Win: []string{
				"any bar",
				"any bar",
				"any bar",
			},
			Bet:     1,
			Payout:  5,
			Message: "All bars!",
		},
		{
			Win: []string{
				"any red",
				"any red",
				"any red",
			},
			Bet:     1,
			Payout:  2,
			Message: "All red!",
		},
		{
			Win: []string{
				"any white",
				"any white",
				"any white",
			},
			Bet:     1,
			Payout:  2,
			Message: "All white!",
		},
		{
			Win: []string{
				"any blue",
				"any blue",
				"any blue",
			},
			Bet:     1,
			Payout:  2,
			Message: "All blue!",
		},
		{
			Win: []string{
				"matching non-blank",
				"matching non-blank",
				"any",
			},
			Bet:     1,
			Payout:  1.5,
			Message: "Two consecutive non-blanks!",
		},
		{
			Win: []string{
				"any",
				"matching non-blank",
				"matching non-blank",
			},
			Bet:     1,
			Payout:  1.5,
			Message: "Two consecutive non-blanks!",
		},
		{
			Win: []string{
				"blank",
				"blank",
				"blank",
			},
			Bet:     1,
			Payout:  1,
			Message: "All blanks!",
		},
	}
)

// Payout defines a winning combination and the payout amounts for different bets.
type Payout struct {
	Win    []string `json:"win" bson:"win"`
	Bet    int      `json:"bet" bson:"bet"`
	Payout float64  `json:"payout" bson:"payout"`
}

// String returns a string representation of the Payout.
func (p *Payout) String() string {
	sb := strings.Builder{}
	sb.WriteString("Payout{")
	sb.WriteString("Win: [")
	for i, slot := range p.Win {
		sb.WriteString(slot)
		if i < len(p.Win)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	sb.WriteString(", Payouts: [")
	sb.WriteString("Bet: " + strconv.Itoa(p.Bet))
	sb.WriteString(", Payout: " + strconv.FormatFloat(p.Payout, 'f', -1, 64))
	sb.WriteString("]")
	sb.WriteString("}")

	return sb.String()
}

// PayoutAmount defines a winning combination and the payout amounts for different bets.
type PayoutAmount struct {
	Win     []string `json:"win" bson:"win"`
	Bet     int      `json:"bet" bson:"bet"`
	Payout  float64  `json:"payout" bson:"payout"`
	Message string   `json:"message,omitempty" bson:"message,omitempty"`
}

// String returns a string representation of the PayoutAmount.
func (p *PayoutAmount) String() string {
	sb := strings.Builder{}
	sb.WriteString("PayoutAmount{")
	sb.WriteString("Win: [")
	sb.WriteString(strings.Join(p.Win, ", "))
	sb.WriteString("]")
	sb.WriteString(", Bet: " + strconv.Itoa(p.Bet))
	sb.WriteString(", Payout: " + strconv.FormatFloat(p.Payout, 'f', -1, 64))
	sb.WriteString("]")
	sb.WriteString("}")

	return sb.String()
}

// PayoutTable defines a table of payouts for a specific guild.
type PayoutTable []PayoutAmount

// String returns a string representation of the PayoutTable.
func (pt PayoutTable) String() string {
	sb := strings.Builder{}
	sb.WriteString("PayoutTable{")
	sb.WriteString("Payouts: [")
	for i, payout := range pt {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(payout.String())
	}
	sb.WriteString("]")
	sb.WriteString("}")
	return sb.String()
}

func (pt *PayoutTable) GetPayoutAmount(bet int, spin []string) (int, string) {
	for _, payout := range *pt {
		if payout.Bet == bet {
			amount := payout.GetPayoutAmount(bet, spin)
			if amount > 0 {
				return amount, payout.Message
			}
		}
	}
	return 0, ""
}

// GetPayoutAmount returns the payout amount for a given bet and spin result.
func (payout *PayoutAmount) GetPayoutAmount(bet int, spin []string) int {
	match := true
	symbols := make([]string, len(spin))
	for i, symbol := range spin {
		if !symbolMatch(symbol, payout.Win[i]) {
			match = false
			break
		}
		symbols[i] = symbol
	}
	if !match {
		return 0
	}
	if strings.Contains(payout.Win[0], "matching") && strings.Contains(payout.Win[1], "matching") {
		if symbols[0] != symbols[1] || symbols[0] == symbols[2] {
			return 0
		}
	}
	if strings.Contains(payout.Win[1], "matching") && strings.Contains(payout.Win[2], "matching") {
		if symbols[1] != symbols[2] || symbols[1] == symbols[0] {
			return 0
		}
	}

	return int(float64(bet) * payout.Payout)
}

// symbolMatch checks if a symbol matches a payout condition.
func symbolMatch(symbol string, match string) bool {
	if symbol == match {
		return true
	}
	switch match {
	case Any:
		return true
	case AnySeven:
		return strings.Contains(symbol, "7")
	case AnyBar:
		return strings.Contains(symbol, "bar")
	case AnyRed:
		return strings.Contains(symbol, "red") || symbol == "1 bar"
	case AnyWhite:
		return strings.Contains(symbol, "white") || symbol == "2 bar"
	case AnyBlue:
		return strings.Contains(symbol, "blue") || symbol == "3 bar"
	case MatchingNonBlank:
		return symbol != "blank"
	default:
		return false
	}
}
