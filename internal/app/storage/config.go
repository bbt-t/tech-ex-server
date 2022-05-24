package storage

type Config struct {
	DBUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DBUrl: "host=localhost dbname=postgres user=iluname password=123 sslmode=disable",
	}
}
