package geodata

import "time"

type Bbox struct {
	MinLat string
	MaxLat string
	MinLon string
	MaxLon string
}

type Supercharger struct {
	ID         int    `json:"id"`
	LocationID string `json:"locationId"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Address    struct {
		Street    string `json:"street"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		CountryID int    `json:"countryId"`
		Country   string `json:"country"`
		RegionID  int    `json:"regionId"`
		Region    string `json:"region"`
	} `json:"address"`
	Gps struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"gps"`
	DateOpened      string `json:"dateOpened"`
	StallCount      int    `json:"stallCount"`
	Counted         bool   `json:"counted"`
	ElevationMeters int    `json:"elevationMeters"`
	PowerKilowatt   int    `json:"powerKilowatt"`
	SolarCanopy     bool   `json:"solarCanopy"`
	Battery         bool   `json:"battery"`
	OtherEVs        bool   `json:"otherEVs"`
	StatusDays      int    `json:"statusDays"`
	URLDiscuss      bool   `json:"urlDiscuss"`
	Stalls          struct {
		V3 int `json:"v3"`
	} `json:"stalls"`
	Plugs struct {
		Tpc  int `json:"tpc"`
		Nacs int `json:"nacs"`
	} `json:"plugs"`
	ParkingID    int    `json:"parkingId"`
	FacilityName string `json:"facilityName"`
	PlugshareID  int    `json:"plugshareId"`
	OsmID        int64  `json:"osmId"`
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
