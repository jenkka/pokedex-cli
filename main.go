package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type CLICommand struct {
	name        string
	description string
	callback    func() error
}

var inputToCliCommand = map[string]CLICommand{
	"help": {
		name:        "help",
		description: "Display a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func commandHelp() error {
	msg := `Welcome to the Pokedex!

Usage:
    help: Displays a help message
    exit: Exit the Pokedex`
	fmt.Println(msg)
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye...")
	os.Exit(0)
	return nil
}

func displayPrompt() {
	fmt.Print("Pokedex>> ")
}

func handleInvalidCommand(invalidInput string) {
	fmt.Printf("Invalid command: %s\n", invalidInput)
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	return strings.Fields(output)
}

func readUserInput(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return cleanInput(scanner.Text())
}

func validateUserInput(input []string) error {
	if len(input) == 0 {
		return errors.New("")
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		displayPrompt()
		userInput := readUserInput(scanner)
		inputError := validateUserInput(userInput)
		if inputError != nil {
			fmt.Print(inputError)
			continue
		}

		commandInput := userInput[0]

		cliCommand, cliCommandExists := inputToCliCommand[commandInput]
		if !cliCommandExists {
			handleInvalidCommand(commandInput)
			continue
		}

		cliCommandErr := cliCommand.callback()
		if cliCommandErr != nil {
			fmt.Printf("The following error occurred while running "+
				"the command '%s': %s", commandInput, cliCommandErr)
			continue
		}
	}
}

