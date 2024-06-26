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
	Callback    func(*config.Config, ...string) error
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
		"explore": {
			Name:        "explore <location_area_name>",
			Description: "List pokemon in the location area",
			Callback:    CommandExplore,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
	}
}

func CommandExit(cfg *config.Config, args ...string) error {
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *config.Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	commands := GetCommands()
	for _, command := range commands {
		fmt.Printf(" - %s: %s\n", command.Name, command.Description)
	}

	return nil
}

func CommandMap(cfg *config.Config, args ...string) error {

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

func CommandMapb(cfg *config.Config, args ...string) error {
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

func CommandExplore(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.PokeapiCleint.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemons in %s:\n", locationAreaName)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(" _", pokemon.Pokemon.Name)
	}

	return nil
}
