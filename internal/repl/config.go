package repl

import "github.com/DavoReds/pokego/internal/pokecache"

type Config struct {
	Next     *string
	Previous *string
	Cache    pokecache.Cache
}
