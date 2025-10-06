package slots

import (
	"math/rand"
	"strings"
)

const (
	NUM_SPINS = 3
)

var (
	defaultLookupTable = LookupTable{
		Reel{
			"2 bar",
			"2 bar",
			"2 bar",
			"blank",
			"blank",
			"3 bar",
			"3 bar",
			"blank",
			"blank",
			"blank",
			"white 7",
			"white 7",
			"white 7",
			"white 7",
			"white 7",
			"white 7",
			"blank",
			"blank",
			"blank",
			"1 bar",
			"1 bar",
			"1 bar",
			"blank",
			"blank",
			"blank",
			"blue 7",
			"blue 7",
			"blue 7",
			"blue 7",
			"blue 7",
			"blue 7",
			"blank",
			"blank",
			"blank",
			"2 bar",
			"2 bar",
			"blank",
			"blank",
			"3 bar",
			"blank",
			"blank",
			"blank",
			"blank",
			"blank",
			"red 7",
			"blank",
			"blank",
			"blank",
			"blank",
			"blank",
			"3 bar",
			"3 bar",
			"3 bar",
			"blank",
			"blank",
			"2 bar",
			"2 bar",
			"blank",
			"blank",
			"1 bar",
			"1 bar",
			"1 bar",
			"blank",
			"blank",
		},
		Reel{
			"2 bar",
			"2 bar",
			"blank",
			"blank",
			"3 bar",
			"3 bar",
			"blank",
			"blank",
			"blank",
			"white 7",
			"blank",
			"blank",
			"blank",
			"1 bar",
			"1 bar",
			"1 bar",
			"1 bar",
			"blank",
			"blank",
			"blank",
			"blue 7",
			"blue 7",
			"blue 7",
			"blue 7",
			"blue 7",
			"blue 7",
			"blue 7",
			"blank",
			"blank",
			"blank",
			"2 bar",
			"2 bar",
			"blank",
			"blank",
			"3 bar",
			"3 bar",
			"blank",
			"blank",
			"blank",
			"blank",
			"blank",
			"red 7",
			"red 7",
			"red 7",
			"blank",
			"blank",
			"blank",
			"blank",
			"blank",
			"3 bar",
			"3 bar",
			"3 bar",
			"blank",
			"blank",
			"2 bar",
			"2 bar",
			"blank",
			"blank",
			"1 bar",
			"1 bar",
			"1 bar",
			"1 bar",
			"blank",
			"blank",
		},
		Reel{
			"2 bar",
			"2 bar",
			"2 bar",
			"blank",
			"blank",
			"3 bar",
			"blank",
			"blank",
			"blank",
			"white 7",
			"white 7",
			"white 7",
			"white 7",
			"white 7",
			"white 7",
			"white 7",
			"blank",
			"blank",
			"blank",
			"1 bar",
			"1 bar",
			"1 bar",
			"1 bar",
			"1 bar",
			"blank",
			"blank",
			"blank",
			"blue 7",
			"blank",
			"blank",
			"blank",
			"2 bar",
			"2 bar",
			"2 bar",
			"blank",
			"blank",
			"3 bar",
			"blank",
			"blank",
			"blank",
			"blank",
			"blank",
			"red 7",
			"blank",
			"blank",
			"blank",
			"blank",
			"blank",
			"3 bar",
			"3 bar",
			"3 bar",
			"blank",
			"blank",
			"2 bar",
			"2 bar",
			"2 bar",
			"blank",
			"blank",
			"1 bar",
			"1 bar",
			"1 bar",
			"1 bar",
			"blank",
			"blank",
		},
	}
)

// Reel represents a single reel of symbols in the slot machine.
type Reel []string

// String returns a string representation of the Reel.
func (r Reel) String() string {
	sb := strings.Builder{}
	sb.WriteString("Reel{")
	sb.WriteString(strings.Join([]string(r), ","))
	sb.WriteString("}")
	return sb.String()
}

// LookupTable represents the lookup table for a guild, containing the reels of slot symbols.
// The lookup table is used to determine the outcome of spins in the slot machine game.
type LookupTable []Reel

// String returns a string representation of the LookupTable.
func (lt LookupTable) String() string {
	sb := strings.Builder{}
	sb.WriteString("LookupTable{")
	sb.WriteString(", Reels: [")
	for i, reel := range lt {
		sb.WriteString(reel.String())
		if i < len(lt)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	sb.WriteString("}")
	return sb.String()
}

// GetPaylineSpin selects a random symbol from each reel to create the current spin.
// It returns the indices of the selected symbols and the symbols themselves.
func (lt LookupTable) GetPaylineSpin() ([]int, []string) {
	currentIndices := make([]int, 0, len(lt))
	for _, reel := range lt {
		randIndex := rand.Int31n(int32(len(reel)))
		currentIndices = append(currentIndices, int(randIndex))
	}
	currentSpin := make([]string, 0, 3)
	for i, reel := range lt {
		currentSpin = append(currentSpin, reel[currentIndices[i]])
	}
	return currentIndices, currentSpin
}

// GetPreviousSpin determines the previous spin based on the current indices.
// It returns the indices of the previous symbols and the symbols themselves.
// The previous symbol for each reel is the first symbol that is different from the current symbol,
func (lt LookupTable) GetPreviousSpin(currentIndices []int) ([]int, []string) {
	previousSpin := make([]string, 0, 3)
	previousIndices := make([]int, 0, len(lt))
	for i, reel := range lt {
		previousIndex := lt.GetPreviousIndex(reel, currentIndices[i])
		previousSpin = append(previousSpin, reel[previousIndex])
		previousIndices = append(previousIndices, previousIndex)
	}
	return previousIndices, previousSpin
}

// GetPreviousIndex finds the index of the previous symbol in the reel that is different from the current symbol.
// It wraps around to the end of the reel if necessary.
func (lt LookupTable) GetPreviousIndex(reel Reel, currentIndex int) int {
	currentSymbol := reel[currentIndex]
	previousIndex := currentIndex
	for {
		previousIndex--
		if previousIndex < 0 {
			previousIndex = len(reel) - 1
		}
		if reel[previousIndex] != currentSymbol {
			break
		}
	}
	return previousIndex
}

// GetNextSpin determines the next spin based on the current indices.
// It returns the indices of the next symbols and the symbols themselves.
// The next symbol for each reel is the first symbol that is different from the current symbol.
func (lt LookupTable) GetNextSpin(currentIndices []int, previousIndices []int) ([]int, []string) {
	nextSpin := make([]string, 0, 3)
	nextIndices := make([]int, 0, len(lt))
	for i, reel := range lt {
		nextIndex := lt.GetNextIndex(reel, currentIndices[i], previousIndices[i])
		nextSpin = append(nextSpin, reel[nextIndex])
		nextIndices = append(nextIndices, nextIndex)
	}
	return nextIndices, nextSpin
}

// GetNextIndex finds the index of the next symbol in the reel that is different from the current symbol.
// It wraps around to the beginning of the reel if necessary.
func (lt LookupTable) GetNextIndex(reel Reel, currentIndex int, previousIndex int) int {
	currentSymbol := reel[currentIndex]
	previousSymbol := reel[previousIndex]
	nextIndex := currentIndex
	for {
		nextIndex++
		if nextIndex > len(reel)-1 {
			nextIndex = 0
		}
		if reel[nextIndex] != currentSymbol && reel[nextIndex] != previousSymbol {
			break
		}
	}
	return nextIndex
}
