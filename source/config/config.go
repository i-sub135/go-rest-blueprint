package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")
var cfg Config

// LoadConfig loads config from a YAML file (optional) and env overrides.
// Call once at bootstrap.
func LoadConfig(path string) error {
	// try load file if present
	if path != "" {
		if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
			log.Printf("warning: config file not loaded: %v", err)
			// continue - env can still provide values
		}
	}

	// env provider: convert APP_ENV -> app.env
	if err := k.Load(env.Provider("", ".", func(s string) string {
		return strings.ToLower(strings.ReplaceAll(s, "_", "."))
	}), nil); err != nil {
		return err
	}

	// set defaults if not provided
	if k.String("app.name") == "" {
		k.Set("app.name", "github.com/i-sub135/go-rest-blueprint")
	}
	if k.String("app.env") == "" {
		k.Set("app.env", "local")
	}
	if k.Int("app.port") == 0 {
		k.Set("app.port", 8080)
	}
	if k.String("log.level") == "" {
		k.Set("log.level", "debug")
	}
	if k.String("db.dsn") == "" {
		k.Set("db.dsn", "host=localhost user=postgres password=postgres dbname=myapp port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		return err
	}
	return nil
}

func GetConfig() *Config  { return &cfg }
func Koanf() *koanf.Koanf { return k }

// ResetConfig resets the global config state - mainly for testing
func ResetConfig() {
	k = koanf.New(".")
	cfg = Config{}
}
