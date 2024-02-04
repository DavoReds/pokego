package repl

import (
	"github.com/DavoReds/pokego/internal/pokeapi/types"
	"github.com/DavoReds/pokego/internal/pokecache"
)

type State struct {
	Next     *string
	Previous *string
	Cache    pokecache.Cache
	Pokedex  map[string]types.Pokemon
}
