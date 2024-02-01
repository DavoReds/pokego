package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type MapResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetMapAreas(url string) (MapResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return MapResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return MapResponse{}, err
	}

	response := MapResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return MapResponse{}, err
	}

	return response, nil
}
