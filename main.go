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
	CallClear()
	prompt := promptui.Select{
		Label: "Choose your option!",
		Items: []string{"Add Item", "Remove Item", "Edit Params", "Quit"},
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
	case "Edit Params":
		admin_track()
	case "Quit":
		os.Exit(0)
	}
}

func add_track() {
	var bin string = "B132"
	var Quantity int = 69


	CallClear()
	for { 
		fmt.Print("What bin are you adding to?\nBin: ")
		bin := string_prompt()
		if bin == "exit" {
			break
		}

		// if bin (in database) {
		//     Quanitity = bin quanitity
		//     
		// }
		CallClear()
	}
	var AmmoType string = "420mm"
	// AmmoType = bin -> ammotype in sql table
	for {
		fmt.Printf("Current quantity of %s: %d.\nHow much would you like to add?\n#: ", AmmoType, Quantity)
		inputNumber, err := number_prompt()
		if inputNumber == -1 { // exit inputed
			break
		}
		if err != nil {
			CallClear()
			fmt.Println(err)
		} else {
			CallClear()
			if yes_no_prompt(fmt.Sprintf("Add %d items to %s", inputNumber, bin)) {
				// Add inputNumber to Ammotype in sql database
			}
			break
		}
	}

	start_menu()
 }

func remove_track() {
	var input int
	var AmmoType string = "420mm"
	var Quantity int = 69

	CallClear()
	fmt.Print("What bin are you removing from?\nBin #: ")
	fmt.Scan(&input)
	CallClear()
	fmt.Printf("Current quantity of %v: %v.\nHow much would you like to remove?\n", AmmoType, Quantity)
	fmt.Scan(&input)
	start_menu()
}

func admin_track() {
	fmt.Println("Admining")
}
