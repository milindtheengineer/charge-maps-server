package geodata

import "time"

type Bbox struct {
	MinLat string
	MaxLat string
	MinLon string
	MaxLon string
}

type LocationData struct {
	Name string
	Lat  float64
	Lon  float64
	City string
}

type Data struct {
	Version   float64 `json:"version"`
	Generator string  `json:"generator"`
	Osm3S     struct {
		TimestampOsmBase   time.Time `json:"timestamp_osm_base"`
		TimestampAreasBase time.Time `json:"timestamp_areas_base"`
		Copyright          string    `json:"copyright"`
	} `json:"osm3s"`
	Elements []struct {
		Type   string            `json:"type"`
		ID     int64             `json:"id"`
		Lat    float64           `json:"lat,omitempty"`
		Lon    float64           `json:"lon,omitempty"`
		Tags   map[string]string `json:"tags"`
		Center struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"center,omitempty"`
		Nodes []interface{} `json:"nodes,omitempty"`
	} `json:"elements"`
}
