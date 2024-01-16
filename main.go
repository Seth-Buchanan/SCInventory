package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

func main() {
	start_menu();
}

func start_menu() {
	for {
		CallClear()
		prompt := promptui.Select{
			Label: "Choose your option!",
			Items: []string{"Add Item", "Remove Item", "Quit"},
		}
		_, result, err := prompt.Run()

		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		CallClear()
		switch result {
		case "Add Item":
			add_track()
		case "Remove Item":
			remove_track()
		case "Quit":
			os.Exit(0)
		}
	}
}

func add_track() {
	var name string
	var quantity int
	var bin_id int
	CallClear()
	for {
		var err error
		fmt.Print("What bin are you adding to?   (q to quit)\nBin: ")
		bin_id, err = number_prompt()
		if bin_id == -1 {	// exit inputed
			return
		}
		if err != nil {
			CallClear()
			fmt.Println(err)
		} else {
			quantity, name = item_quantity(bin_id)
			if quantity == -1  { // bin not found
				CallClear()
				fmt.Printf("A bin with the id %d was not found\n", bin_id)
			} else { // all is good
				CallClear()
				break
			}
			
		}

	}
	for {
		fmt.Printf("Current quantity of %s: %d\nHow much would you like to add?   (q to quit)\n#: ", name, quantity)
		inputNumber, err := number_prompt()
		if inputNumber == -1 { // exit inputed
			return
		}
		if err != nil {
			CallClear()
			fmt.Println(err)
		} else {
			CallClear()
			if yes_no_prompt(fmt.Sprintf("Add %d items to bin %d? (New total %d)", inputNumber, bin_id, quantity-inputNumber)) {
				add_quantity(bin_id, inputNumber)
			} 
			return
		}
	}
 }

func remove_track() {
	var name string
	var quantity int
	var bin_id int
	CallClear()
	for {
		var err error
		fmt.Print("What bin are you taking from?   (q to quit)\nBin: ")
		bin_id, err = number_prompt()
		if bin_id == -1 {	// exit inputed
			return
		}
		if err != nil {
			CallClear()
			fmt.Println(err)
		} else {
			quantity, name = item_quantity(bin_id)
			if quantity == -1  { // bin not found
				CallClear()
				fmt.Printf("A bin with the id %d was not found\n", bin_id)
			} else { // all is good
				CallClear()
				break
			}
			
		}

	}
	for {
		fmt.Printf("Current quantity of %s: %d\nHow much would you like to remove?   (q to quit)\n#: ", name, quantity)
		inputNumber, err := number_prompt()
		if inputNumber == -1 { // exit inputed
			return
		}
		if err != nil {
			CallClear()
			fmt.Println(err)
		} else {
			CallClear()
			if yes_no_prompt(fmt.Sprintf("Remove %d items to bin %d? (New total %d)", inputNumber, bin_id, quantity-inputNumber)) {
				subtract_quantity(bin_id, inputNumber)
				// Check if quantity is less than limit
			} 
			return
		}
	}
 }


