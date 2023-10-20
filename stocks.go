package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type Stock struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
	Price  float64
}

type Values struct {
	Open float64 `json:"open"`
}

func SearchTicker(ticker string) []Stock {
	ApiKey := "a key"
	PolygonPath := "https://api.polygon.io"
	resp, err := http.Get(PolygonPath + "/v3/reference/tickers?" + ApiKey + "&ticker" + strings.ToUpper(ticker))

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)

	data := struct {
		Results []Stock `json:"results"`
	}{}

	json.Unmarshal(body, &data)
	return data.Results
}

func GetDailyValues(ticker string) Values {
	ApiKey := "apiKey=xP0hsWdLO8GhzlO7t30yYvYhyM5YTs8G"
	PolygonPath := "https://api.polygon.io"
	resp, err := http.Get(PolygonPath + "/v3/open-close/" + ApiKey + strings.ToUpper(ticker))

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)

	var data Values

	json.Unmarshal(body, &data)
	return data
}
