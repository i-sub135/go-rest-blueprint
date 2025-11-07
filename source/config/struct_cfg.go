package config

type Config struct {
	App struct {
		Name    string `koanf:"name"`
		Mode    string `koanf:"mode"`
		Port    int    `koanf:"port"`
		Version string `koanf:"version"`
	} `koanf:"app"`
	DB struct {
		DSN string `koanf:"dsn"`
	} `koanf:"db"`
	Log struct {
		Level         string `koanf:"level"`
		PrettyConsole bool   `koanf:"pretty_console"`
	} `koanf:"log"`
}
