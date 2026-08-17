package store

//rey
type Config struct {
	// Add configuration fields for your store here, e.g., database connection string, etc.
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
