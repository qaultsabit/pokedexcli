package command

import (
	"errors"
	"fmt"
	"math/rand"
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
		"catch": {
			Name:        "catch <pokemon_name>",
			Description: "Attemp to catch a pokemon and add it to your pokedex",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>",
			Description: "Display information about a caought pokemon",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "inspect",
			Description: "Display all pokemons in your pokedex",
			Callback:    CommandPokedex,
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
		fmt.Printf(" - %s\t: %s\n", command.Name, command.Description)
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
	pokemonName := args[0]

	locationArea, err := cfg.PokeapiCleint.GetLocationArea(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemons in %s:\n", pokemonName)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(" _", pokemon.Pokemon.Name)
	}

	return nil
}

func CommandCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.PokeapiCleint.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.CaughtPokemon[pokemonName] = pokemon

	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}

func CommandInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.CaughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	return nil
}

func CommandPokedex(cfg *config.Config, args ...string) error {
	if len(cfg.CaughtPokemon) == 0 {
		return errors.New("there are no pokemon in your pokedex")
	}
	fmt.Println("Pokemons in your Pokedex:")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Println(" -", pokemon.Name)
	}

	return nil
}
