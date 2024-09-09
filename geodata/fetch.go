package geodata

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/tidwall/rtree"
)

func FetchSuperchargerData(dataFile string) (*SyncRTree, error) {
	var data []Supercharger
	syncRTree := &SyncRTree{
		mu: sync.Mutex{},
		tr: rtree.RTree{},
	}
	file, err := os.Open(dataFile)
	if err != nil {
		return syncRTree, fmt.Errorf("FetchData: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return syncRTree, fmt.Errorf("FetchData: %w", err)
	}
	for _, elem := range data {
		if elem.Status == "OPEN" {
			address := elem.Address.Street + "\n" + elem.Address.City + " " + elem.Address.State + " " + elem.Address.Zip
			syncRTree.InsertPoint(elem.Gps.Longitude, elem.Gps.Latitude, "Tesla Supercharger", strings.TrimSpace(address), elem.StallCount, elem.PowerKilowatt, "", "supercharger")
		}
	}
	return syncRTree, nil
}

func FetchData(dataFile string, key string) (*SyncRTree, error) {
	var data Data
	syncRTree := &SyncRTree{
		mu: sync.Mutex{},
		tr: rtree.RTree{},
	}
	file, err := os.Open(dataFile)
	if err != nil {
		return syncRTree, fmt.Errorf("FetchData: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return syncRTree, fmt.Errorf("FetchData: %w", err)
	}
	for _, elem := range data.Elements {
		if elem.Type == "node" {
			address := elem.Tags["addr:housenumber"] + " " + elem.Tags["addr:street"] + "\n" + elem.Tags["addr:state"] + " " + elem.Tags["addr:postcode"]
			syncRTree.InsertPoint(elem.Lon, elem.Lat, elem.Tags["name"], strings.TrimSpace(address), 0, 0, elem.Tags["website"], elem.Tags["brand"])
		} else if elem.Type == "way" {
			address := elem.Tags["addr:housenumber"] + " " + elem.Tags["addr:street"] + "\n" + elem.Tags["addr:state"] + " " + elem.Tags["addr:postcode"]
			syncRTree.InsertPoint(elem.Center.Lon, elem.Center.Lat, elem.Tags["name"], strings.TrimSpace(address), 0, 0, elem.Tags["website"], elem.Tags["brand"])
		}
	}
	return syncRTree, nil
}
