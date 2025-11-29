package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaExploreResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(page *pageConfig) error {
	exploreAreaName := page.Parameters[0]
	if exploreAreaName == "" {
		err := errors.New("you need to send an location area name to explore")
		fmt.Println(err)
		return err
	}

	fullUrl := "https://pokeapi.co/api/v2/location-area/" + exploreAreaName

	data, ok := page.Cache.Get(fullUrl)
	if ok {
		response := LocationAreaExploreResponse{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}

		for _, pokemonEncounter := range response.PokemonEncounters {
			fmt.Println(pokemonEncounter.Pokemon.Name)
		}
		return nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("call to location-area name endpoint return status code %d", res.StatusCode)
	}

	page.Cache.Add(fullUrl, body)

	response := LocationAreaExploreResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	for _, pokemonEncounter := range response.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}

	return nil
}
