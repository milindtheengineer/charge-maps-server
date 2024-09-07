package geodata

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/tidwall/rtree"
)

type SyncRTree struct {
	mu sync.Mutex
	tr rtree.RTree
}

func (s *SyncRTree) SearchPoint(bbox Bbox) ([]LocationData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var locationDataList []LocationData
	minLon, err := strconv.ParseFloat(bbox.MinLon, 64)
	if err != nil {
		return nil, fmt.Errorf("SearchPoint: %v", err)
	}
	minLat, err := strconv.ParseFloat(bbox.MinLat, 64)
	if err != nil {
		return nil, fmt.Errorf("SearchPoint: %v", err)
	}
	maxLon, err := strconv.ParseFloat(bbox.MaxLon, 64)
	if err != nil {
		return nil, fmt.Errorf("SearchPoint: %v", err)
	}
	maxLat, err := strconv.ParseFloat(bbox.MaxLat, 64)
	if err != nil {
		return nil, fmt.Errorf("SearchPoint: %v", err)
	}
	s.tr.Search([2]float64{minLon, minLat}, [2]float64{maxLon, maxLat},
		func(min, max [2]float64, data interface{}) bool {
			locationDataList = append(locationDataList, data.(LocationData))
			return true
		},
	)
	return locationDataList, nil
}

func (s *SyncRTree) InsertPoint(lon float64, lat float64, name string, address string, numberOfChargingStalls int, power int, website string, shortName string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tr.Insert([2]float64{lon, lat}, [2]float64{lon, lat}, LocationData{Name: name, Lon: lon, Lat: lat, Address: address, NumberOfChargingStalls: numberOfChargingStalls, Power: power, Website: website, ShortName: shortName})
}
