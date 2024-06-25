package main

import (
	"github.com/qaultsabit/pokedexcli/pkg/config"
	"github.com/qaultsabit/pokedexcli/pkg/repl"
)

func main() {
	cfg := config.NewConfig()
	repl.Start(&cfg)
}
