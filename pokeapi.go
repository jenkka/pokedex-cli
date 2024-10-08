package main

import (
    "net/http"
    "encoding/json"
    "io"
    "fmt"
    "errors"
)

var apiUrl string = "https://pokeapi.co/api/v2/"
var locationAreaUrl string = apiUrl + "location-area/"

type PokeAPI struct {
    next *string
    previous *string
}

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewPokeAPI() PokeAPI {
    return PokeAPI{next: nil, previous: nil}
}

func (api *PokeAPI) Map() error {
    var url string
    if api.next != nil {
        url = *(api.next)
    } else {
        url = locationAreaUrl
    }

    res, err := http.Get(url)
    if err != nil {
        return err
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        return err 
    }

    if res.StatusCode != http.StatusOK {
        return fmt.Errorf("Response failed with status code: %d and\n"+
                                 "body: %s\n", res.StatusCode, body)
    }

    locationAreas := LocationAreas{}
    err = json.Unmarshal(body, &locationAreas)
    if err != nil {
        return err
    }

    for _, area := range locationAreas.Results {
        fmt.Println(area.Name)
    }
    api.next = locationAreas.Next
    api.previous = locationAreas.Previous

    return nil
}

func (api *PokeAPI) MapB() error {
    var url string
    if api.previous != nil {
        url = *(api.previous)
    } else {
        return errors.New("No previous locations.\n")
    }

    res, err := http.Get(url)
    if err != nil {
        return err
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        return err 
    }

    if res.StatusCode != http.StatusOK {
        return fmt.Errorf("Response failed with status code: %d and\n"+
                                 "body: %s\n", res.StatusCode, body)
    }

    locationAreas := LocationAreas{}
    err = json.Unmarshal(body, &locationAreas)
    if err != nil {
        return err
    }

    for _, area := range locationAreas.Results {
        fmt.Println(area.Name)
    }
    api.next = locationAreas.Next
    api.previous = locationAreas.Previous

    return nil
}

