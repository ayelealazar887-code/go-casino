package main

import (
	"fmt"
)

func getName() string {

	var name string = ""
	fmt.Println("Enter your name")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return  ""
	}
	fmt.Printf("hello %s\n", name)
	return  name
}


func getBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance = %s)", balance)
		fmt.Scan (&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
		   break
		}
	}

	return  bet
}


func main() {
	
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 10,
		"D": 20,
	}
	multiplier := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 7,
		"D": 4,
	}
	balance := uint(500)
	getName()

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}

		balance-=bet
	}

	fmt.Printf("You left with %d\n", balance)
}