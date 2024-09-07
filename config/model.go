package config

type Config struct {
	Debug                bool
	GeoJSONFilePath      map[string]string
	SuperchargerFilePath string
	GoogleToken          string
	SigningKey           string
	Cors                 []string
}
