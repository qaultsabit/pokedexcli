package config

import (
	"time"

	"github.com/qaultsabit/pokedexcli/internal/pokeapi"
)

type Config struct {
	PokeapiCleint       pokeapi.Client
	NextLoactionAreaURL *string
	PrevLocationAreaURL *string
	CaughtPokemon       map[string]pokeapi.Pokemon
}

func NewConfig() Config {
	return Config{
		PokeapiCleint: pokeapi.NewClient(time.Minute),
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}
}
