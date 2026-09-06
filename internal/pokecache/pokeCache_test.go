package pokecache

import (
	"testing"
	"time"
	"fmt"
)

func TestAddGet(t *testing.T) {
	const interval = time.Second * 5
	cases := []struct {
		key   string
		val   []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area/",
			val: []byte("testdata"),	
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
			val: []byte("more test data"),
		},
	}

	for idx, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", idx), func (t *testing.T){
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key")
				return
			}
			if string(val) != string(c.val){
				t.Errorf("Expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime * 2
	cache := NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area/", []byte("testdata"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area/")
	if !ok {
		t.Errorf("Expected to find key")
		return
	}
	
	time.Sleep(waitTime)
	
	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area/")
	if ok {
		t.Errorf("Expected to not find key")
	}
}