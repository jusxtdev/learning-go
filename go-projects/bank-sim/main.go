package main

import (
	"bank_simulator/account"
	"fmt"
	"log"
)

func main() {
	account := account.Account{}

	for {
		choice := print_menu()

		switch choice {
		case 1:
			var holder string
			fmt.Printf("Enter Name : ")
			_, err := fmt.Scanln(&holder)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println()
			account.Init_account(holder)
		case 2:
			var amt float64
			fmt.Printf("Enter amount to deposit : ")
			_, err := fmt.Scanln(&amt)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println()
			account.Deposit(amt)
		case 3:
			var amt float64
			fmt.Printf("Enter amount to withdraw: ")
			_, err := fmt.Scanln(&amt)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println()
			account.Withdraw(amt)
		case 4:
			account.Display()
		}
	}

}

func print_menu() int {
	var choice int
	menu := "1. Create Account\n2. Deposit Amount \n3. Withdraw Amount\n4. Display Account"

	fmt.Println(menu)
	fmt.Printf("Enter Choice : ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	return choice
}
