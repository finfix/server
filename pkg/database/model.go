package database

type RepoConfig struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"SERVER_DB_USER"`
	Password string `env:"SERVER_DB_PASSWORD"`
}
