package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/qaultsabit/pokedexcli/pkg/command"
	"github.com/qaultsabit/pokedexcli/pkg/config"
)

func Start(cfg *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		inputs := cleanInput(input)
		if len(inputs) == 0 {
			continue
		}

		commandName := inputs[0]
		args := []string{}
		if len(inputs) > 1 {
			args = inputs[1:]
		}

		availableCommands := command.GetCommands()
		cmd, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		if err := cmd.Callback(cfg, args...); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
