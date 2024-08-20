package web

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/milindtheengineer/charge-maps-server/config"
	"github.com/milindtheengineer/charge-maps-server/geodata"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func StartRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	syncRtreeMap, err := geodata.FetchData(config.AppConfig.GeoJSONFilePath)
	if err != nil {
		logger.Panic().Msgf("screwed due to %v", err)
	}
	app := App{
		logger: logger,
		geoMap: syncRtreeMap,
	}
	r.Get("/health", HealthHandler)
	r.Post("/login", app.HandleLogin)
	r.Group(func(r chi.Router) {
		// r.Use(app.authMiddleware)
		r.Post("/locations", app.LocationHanlder)
	})

	// r.GET("/v1/user", authMiddleware(user.Crud))
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Panic().Msg(err.Error())
	}
}
