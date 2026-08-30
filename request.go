package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func makePokedexRequest(cfg * config, url string) MapResult {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MapResult{}
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return MapResult{}
	}
	defer res.Body.Close()

	var mapData MapResult
	if err := json.NewDecoder(res.Body).Decode(&mapData); err != nil {
		return MapResult{}
	}

	cfg.next = mapData.Next
	cfg.previous = mapData.Previous

	return mapData	
}