package repl

import (
	"github.com/DavoReds/pokego/internal/pokeapi/responses"
	"github.com/DavoReds/pokego/internal/pokecache"
)

type State struct {
	Next     *string
	Previous *string
	Cache    pokecache.Cache
	Pokedex  map[string]responses.Pokemon
}
