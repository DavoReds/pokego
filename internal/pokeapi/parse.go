package pokeapi

import (
	"encoding/json"
)

// Parses a given JSON payload into a given variable. Returns an error if it
// failed to do so.
func Parse[T any](body []byte, result *T) error {
	if err := json.Unmarshal(body, result); err != nil {
		return err
	}

	return nil
}
