package config

import "github.com/qaultsabit/pokedexcli/internal/pokeapi"

type Config struct {
	PokeapiCleint       pokeapi.Client
	NextLoactionAreaURL *string
	PrevLocationAreaURL *string
}

func NewConfig() Config {
	return Config{
		PokeapiCleint: pokeapi.NewClient(),
	}
}
