# Slots

A Go package for simulating slot machine games with configurable reels, symbols, and payout tables.

## Features

- **Configurable Slot Machine**: Create slot machines with custom reel configurations and payout tables
- **Realistic Spin Simulation**: Generates three-line displays (top, payline, bottom) for authentic slot machine experience
- **Flexible Payout System**: Support for various winning combinations including wildcards and symbol groups
- **Built-in Symbol Matching**: Advanced pattern matching for "any bar", "any red", "any white", "any blue", etc.
- **Probability Analysis**: Command-line tool for analyzing payout probabilities

## Package Structure

```text
github.com/rbrabson/slots/
├── slots.go        # Main SlotMachine type and core functionality
├── lookup.go       # Reel definitions and spin logic (LookupTable, Reel types)
├── payout.go       # Payout calculations and symbol matching (PayoutTable, PayoutAmount types)
├── cmd/
│   └── payout/     # Command-line tool for analyzing payout probabilities
│       └── main.go
└── go.mod
```

## Installation

```bash
go get github.com/rbrabson/slots
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/rbrabson/slots"
)

func main() {
    // Create a new slot machine with default configuration
    sm := slots.NewSlotMachine()
    
    // Spin with a bet of 1 coin
    result := sm.Spin(1)
    
    fmt.Printf("Top Line:    %v\n", result.TopLine)
    fmt.Printf("Payline:     %v\n", result.Payline)
    fmt.Printf("Bottom Line: %v\n", result.BottomLine)
    fmt.Printf("Bet: %d, Payout: %d\n", result.Bet, result.Payout)
    
    if result.Payout > 0 {
        fmt.Printf("Winner! %s\n", result.Message)
    }
}
```

## Core Types

### SlotMachine

The main type that orchestrates the slot machine functionality:

```go
type SlotMachine struct {
    LookupTable LookupTable
    PayoutTable PayoutTable
}
```

**Methods:**

- `NewSlotMachine(opts ...Option) *SlotMachine` - Create a new slot machine
- `Spin(bet int) *SpinResult` - Perform a spin with the given bet amount

### SpinResult

Contains the result of a single spin:

```go
type SpinResult struct {
    TopLine    []string  // Symbols above the payline
    Payline    []string  // Winning line symbols
    BottomLine []string  // Symbols below the payline
    Bet        int       // Amount bet
    Payout     int       // Amount won
    Message    string    // Win description
}
```

### Configuration Options

Customize your slot machine using functional options:

```go
// Custom lookup table (reel configuration)
sm := slots.NewSlotMachine(
    slots.WithLookupTable(customReels),
)

// Custom payout table
sm := slots.NewSlotMachine(
    slots.WithPayoutTable(customPayouts),
)

// Both custom tables
sm := slots.NewSlotMachine(
    slots.WithLookupTable(customReels),
    slots.WithPayoutTable(customPayouts),
)
```

## Payout System

The package supports sophisticated winning combinations:

### Exact Matches

```go
[]string{"red 7", "red 7", "red 7"}  // Three red 7s
```

### Wildcard Patterns

```go
[]string{"any", "any", "any"}              // Any three symbols
[]string{"any non-blank", "any", "any"}    // Any non-blank + any two
[]string{"any bar", "any bar", "any bar"}  // Any three bar symbols
[]string{"any red", "any red", "any red"}  // Any three red symbols
```

### Color Groups

- `"any red"` - Matches any symbol containing "red"
- `"any white"` - Matches any symbol containing "white"  
- `"any blue"` - Matches any symbol containing "blue"
- `"any bar"` - Matches any symbol containing "bar"

### Symbol Types

- `"any"` - Matches any symbol including blanks
- `"any non-blank"` - Matches any symbol except "blank"
- `"blank"` - Matches only blank symbols

## Default Configuration

The package comes with a realistic slot machine configuration featuring:

**Symbols:**

- Red, white, and blue 7s
- 1 bar, 2 bar, 3 bar
- Blank spaces

**Sample Payouts:**

- Mixed 7s (red, white, blue): 2400x bet
- Three red 7s: 1200x bet
- Three white 7s: 200x bet
- Three blue 7s: 150x bet
- Any three 7s: 50x bet
- Mixed bars (1, 2, 3): 50x bet
- And many more combinations...

## Command Line Tools

### Payout Analysis Tool

Analyze the probability and expected return of your slot machine:

```bash
cd cmd/payout
go run main.go
```

This tool calculates:

- Probability of each winning combination
- Expected return percentage
- Number of possible outcomes

## Examples

### Custom Reel Configuration

```go
customReels := slots.LookupTable{
    slots.Reel{"cherry", "lemon", "orange", "cherry"},  // Reel 1
    slots.Reel{"cherry", "lemon", "orange", "bell"},    // Reel 2  
    slots.Reel{"cherry", "lemon", "orange", "bar"},     // Reel 3
}

sm := slots.NewSlotMachine(
    slots.WithLookupTable(customReels),
)
```

### Custom Payout Table

```go
customPayouts := slots.PayoutTable{
    {
        Win:    []string{"cherry", "cherry", "cherry"},
        Bet:    1,
        Payout: 100,
    },
    {
        Win:    []string{"any", "any", "any"},
        Bet:    1, 
        Payout: 2,
    },
}

sm := slots.NewSlotMachine(
    slots.WithPayoutTable(customPayouts),
)
```

## Dependencies

- `github.com/joho/godotenv` - Environment variable management (used in CLI tools)

## License

See LICENSE file for details.
