package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Pokedex > ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Youre text was: %v \n", input)
}

/*
var cliName string = "Pokedex"
// printPrompt displays repl prompt at the beg of each loop
func printPrompt() {
	fmt.Print(cliName, "> ")
}

// informs user about invalid inputs
func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}



type PDex interface {
	commandHelp()
	commandExit()
}

type cliCommand struct {
		name string
		description string
		callback func() error
}

func commandHelp(string) string {

}

return map[string]cliCommand {
	"help": {
		name: "help",
		desription: "Display a help message",
		callback: commandHelp,
	},
	"exit": {
		name: "exit",
		description: "Exit the Pokedex"
		callback: commandExit,
	},
}
*/
