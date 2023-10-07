package main

import "fmt"

func main() {
	for {
		fmt.Println("Type in a command or q to exit: ")
		var input string
		fmt.Scanln(&input)
		if input == "q" {
			break
		} else {
			continue
		}
	}
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

func commandHelp() string {
	
}

return map[string]cliCommand{
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