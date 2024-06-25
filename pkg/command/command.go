package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/qaultsabit/pokedexcli/pkg/config"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Lists the next page of location areas",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Lists the previous page of location areas",
			Callback:    CommandMapb,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
	}
}

func CommandExit(cfg *config.Config) error {
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *config.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	commands := GetCommands()
	for _, command := range commands {
		fmt.Printf(" - %s: %s\n", command.Name, command.Description)
	}

	return nil
}

func CommandMap(cfg *config.Config) error {

	resp, err := cfg.PokeapiCleint.ListLocationAreas(cfg.NextLoactionAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(" -", area.Name)
	}

	cfg.NextLoactionAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return err
}

func CommandMapb(cfg *config.Config) error {
	if cfg.PrevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.PokeapiCleint.ListLocationAreas(cfg.PrevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(" -", area.Name)
	}

	cfg.NextLoactionAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return err
}
