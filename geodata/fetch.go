package geodata

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/tidwall/rtree"
)

func FetchData(dataFile string) (map[string]*SyncRTree, error) {
	var data Data
	rtreeMap := make(map[string]*SyncRTree)
	// Do this properly:
	rtreeMap["target"] = &SyncRTree{
		mu: sync.Mutex{},
		tr: rtree.RTree{},
	}
	rtreeMap["supercharger"] = &SyncRTree{
		mu: sync.Mutex{},
		tr: rtree.RTree{},
	}
	file, err := os.Open(dataFile)
	if err != nil {
		return rtreeMap, fmt.Errorf("FetchData: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return rtreeMap, fmt.Errorf("FetchData: %w", err)
	}
	for _, elem := range data.Elements {
		if elem.Tags["amenity"] == "charging_station" && elem.Tags["operator"] == "Tesla, Inc." {
			if elem.Type == "node" {
				rtreeMap["supercharger"].InsertPoint(elem.Lon, elem.Lat, "supercharger", elem.Tags["addr:city"])
			} else if elem.Type == "way" {
				rtreeMap["supercharger"].InsertPoint(elem.Center.Lon, elem.Center.Lat, "supercharger", elem.Tags["addr:city"])
			}
		} else if elem.Tags["brand"] == "Target" {
			if elem.Type == "node" {
				rtreeMap["target"].InsertPoint(elem.Lon, elem.Lat, "target", elem.Tags["addr:city"])
			} else if elem.Type == "way" {
				rtreeMap["target"].InsertPoint(elem.Center.Lon, elem.Center.Lat, "target", elem.Tags["addr:city"])
			}
		}
	}
	return rtreeMap, nil
}
