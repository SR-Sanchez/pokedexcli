package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func newClient(timeOutInSeconds int) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Duration(timeOutInSeconds) * time.Second,
		},
	}
}

func (c *Client) makePokedexRequest(cfg * config, url string) (MapResult, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MapResult{}, err
	}

	/* client := &http.Client{
		Timeout: 10 * time.Second,
	} */
	res, err := c.httpClient.Do(req)
	if err != nil {
		return MapResult{}, err
	}
	defer res.Body.Close()

	var mapData MapResult
	if err := json.NewDecoder(res.Body).Decode(&mapData); err != nil {
		return MapResult{}, err
	}

	cfg.next = mapData.Next
	cfg.previous = mapData.Previous

	return mapData, nil	
}