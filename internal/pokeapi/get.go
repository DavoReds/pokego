package pokeapi

import (
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
