package main

import (
	"fmt"
	"math/rand"
)

func getName() string {
	var name string

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	fmt.Printf("Hello %s!\n", name)

	return name
}

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}

	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	// Create rows
	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

	// Generate columns
	for col := 0; col < cols; col++ {
		selected := map[int]bool{}

		for row := 0; row < rows; row++ {

			for true {
				randomIndex := getRandomNumber(0, len(reel)-1)

				_, exists := selected[randomIndex]

				if !exists {
					selected[randomIndex] = true

					result[row] = append(
						result[row],
						reel[randomIndex],
					)

					break
				}
			}
		}
	}

	return result
}

func printSpin(spin [][]string) {
	fmt.Println("\n----------------")

	for _, row := range spin {
		fmt.Print("| ")

		for _, symbol := range row {
			fmt.Printf("%s ", symbol)
		}

		fmt.Println("|")
	}

	fmt.Println("----------------")
}

func getBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf(
			"Enter your bet, or 0 to quit (balance = %d): ",
			balance,
		)

		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}

	return bet
}

func calculateWinnings(spin [][]string, bet uint) uint {
	if len(spin) == 0 {
		return 0
	}

	// Check each row
	for _, row := range spin {

		// Make sure the row has at least 3 symbols
		if len(row) < 3 {
			continue
		}

		if row[0] == row[1] && row[1] == row[2] {

			switch row[0] {
			case "A":
				return bet * 20
			case "B":
				return bet * 10
			case "C":
				return bet * 7
			case "D":
				return bet * 4
			}
		}
	}

	return 0
}

func main() {

	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 10,
		"D": 20,
	}

	symbolArr := generateSymbolArray(symbols)

	name := getName()

	fmt.Printf("\nWelcome to the casino, %s!\n", name)

	balance := uint(500)

	for balance > 0 {

		bet := getBet(balance)

		if bet == 0 {
			break
		}

		// Remove bet from balance
		balance -= bet

		// Spin
		spin := getSpin(symbolArr, 3, 3)

		// Display spin
		printSpin(spin)

		// Calculate winnings
		winnings := calculateWinnings(spin, bet)

		if winnings > 0 {
			fmt.Printf("You won %d!\n", winnings)
			balance += winnings
		} else {
			fmt.Println("You lost!")
		}

		fmt.Printf("Current balance: %d\n\n", balance)
	}

	fmt.Printf("You left with %d\n", balance)
}