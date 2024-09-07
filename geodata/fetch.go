package geodata

import (
	"encoding/json"
	"fmt"
	"os"
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
		syncRTree.InsertPoint(elem.Gps.Longitude, elem.Gps.Latitude, "supercharger", elem.Name)
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
			syncRTree.InsertPoint(elem.Lon, elem.Lat, elem.Tags["name"], elem.Tags["addr:city"])
		} else if elem.Type == "way" {
			syncRTree.InsertPoint(elem.Center.Lon, elem.Center.Lat, elem.Tags["name"], elem.Tags["addr:city"])
		}
	}
	return syncRTree, nil
}
