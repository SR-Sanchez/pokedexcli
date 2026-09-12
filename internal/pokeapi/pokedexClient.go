package pokedexClient

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/SR-Sanchez/pokedexcli/internal/pokecache"
	"io"
	"errors"
	"fmt"
)

type Client struct {
	httpClient    http.Client
	cache         *pokecache.Cache
}

func NewClient(timeOutInSeconds int) Client {
	// cache interval is intentionally separate from the HTTP timeout -
	// they represent different concerns
	cache := pokecache.NewCache(time.Duration(5) * time.Second)
	return Client{
		httpClient: http.Client{
			Timeout: time.Duration(timeOutInSeconds) * time.Second,
		},
		cache: cache,
	}
}

func fetchAndUnmarshal[T any](url string, c *Client) (T, error) {
	var result T
	
	// 1. Check cache first
	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &result); err != nil {
			return result, err
		}
		return result, nil
	}
	
	// 2. Make HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}
	
	res, err := c.httpClient.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	
	// 3. Handle non-200 HTTP statuses
	if res.StatusCode == http.StatusNotFound {
		return result, errors.New("resource not found")
	}
	if res.StatusCode > 299 {
		return result, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	
	// 4. Read response body
	// read the entire response body into memory as raw bytes
	// (res.Body is a stream and can only be read once)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	
	// 5. Save raw bytes to cache
	c.cache.Add(url, data)
	
	// 6. Unmarshal JSON into generic result type T
	// unmarshal from our in-memory bytes, not res.Body (already drained)
	if err := json.Unmarshal(data, &result); err != nil {
		return result, err
	}
	
	return result, nil
}

func (c *Client) MakePokedexRequest(url string) (MapResult, error) {
	return fetchAndUnmarshal[MapResult] (url, c)
}

func (c *Client) ListPokemonsInArea(location string) (PokemonAreaResult, error) {
	return fetchAndUnmarshal[PokemonAreaResult] (location, c)
}

func (c *Client) CatchPokemon(pokemon string) (PokemonInfo, error) {
	return fetchAndUnmarshal[PokemonInfo] (pokemon, c)
}