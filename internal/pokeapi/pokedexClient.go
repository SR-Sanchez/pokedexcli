package pokedexClient

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/SR-Sanchez/pokedexcli/internal/pokecache"
	"io"
)

type Client struct {
	httpClient    http.Client
	cache         *pokecache.Cache
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type MapResult struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
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

func (c *Client) MakePokedexRequest(url string) (MapResult, error) {
	var mapData MapResult

	// check cache first - avoid a network call if we already have this URL
	val, ok := c.cache.Get(url)
	if ok {
		if err := json.Unmarshal(val, &mapData); err != nil {
			return MapResult{}, err
		}
		return mapData, nil // cache hit - skip the HTTP request entirely
	}

	// cache miss - make the real HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MapResult{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return MapResult{}, err
	}
	defer res.Body.Close() // release the connection once we're done reading

	// read the entire response body into memory as raw bytes
	// (res.Body is a stream and can only be read once)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return MapResult{}, err
	}

	// store the raw bytes in the cache for future requests to this URL
	c.cache.Add(url, data)

	// unmarshal from our in-memory bytes, not res.Body (already drained)
	if err := json.Unmarshal(data, &mapData); err != nil {
		return MapResult{}, err
	}

	return mapData, nil
}