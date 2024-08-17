package web

import (
	"net/http"

	"github.com/milindtheengineer/charge-maps-server/database"
	"github.com/rs/zerolog"
)

type App struct {
	db     *database.DBConn
	logger zerolog.Logger
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
