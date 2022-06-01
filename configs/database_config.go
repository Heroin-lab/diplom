package configs

type DBurl struct {
	DatabaseURL string `toml:"database_url"`
}

func NewDBurl() *DBurl {
	return &DBurl{}
}
