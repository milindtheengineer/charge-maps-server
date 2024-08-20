package web

import (
	"encoding/json"
	"net/http"

	"github.com/milindtheengineer/charge-maps-server/database"
	"github.com/milindtheengineer/charge-maps-server/geodata"
	"github.com/rs/zerolog"
)

type App struct {
	db     *database.DBConn
	geoMap map[string]*geodata.SyncRTree
	logger zerolog.Logger
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}

func (a *App) LocationHanlder(w http.ResponseWriter, r *http.Request) {
	var bbox geodata.Bbox
	if err := json.NewDecoder(r.Body).Decode(&bbox); err != nil {
		a.logger.Error().Msgf("LocationHanlder: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := []geodata.LocationData{}
	targetData, err := a.geoMap["target"].SearchPoint(bbox)
	if err != nil {
		a.logger.Error().Msgf("LocationHanlder: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	superchargerdata, err := a.geoMap["supercharger"].SearchPoint(bbox)
	if err != nil {
		a.logger.Error().Msgf("LocationHanlder: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data = append(data, targetData...)
	data = append(data, superchargerdata...)
	body, err := json.Marshal(data)
	if err != nil {
		a.logger.Error().Msgf("LocationHanlder: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(body)
}
