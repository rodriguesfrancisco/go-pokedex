package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(page *pageConfig) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area"
	if page.Next != "" {
		fullUrl = page.Next
	}

	data, ok := page.Cache.Get(fullUrl)
	if ok {
		response := LocationAreaResponse{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}

		for _, area := range response.Results {
			fmt.Println(area.Name)
		}
		page.Next = response.Next
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
		return fmt.Errorf("call to location-area endpoint return status code %d", res.StatusCode)
	}

	page.Cache.Add(fullUrl, body)

	response := LocationAreaResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	page.Next = response.Next

	return nil
}

func commandMapb(page *pageConfig) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area"
	if page.Previous != "" {
		fullUrl = page.Previous
	}

	data, ok := page.Cache.Get(fullUrl)
	if ok {
		response := LocationAreaResponse{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return err
		}

		for _, area := range response.Results {
			fmt.Println(area.Name)
		}
		page.Previous = response.Previous
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
		return fmt.Errorf("call to location-area endpoint return status code %d", res.StatusCode)
	}

	page.Cache.Add(fullUrl, body)

	response := LocationAreaResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	page.Previous = response.Previous

	return nil
}
