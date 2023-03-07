package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerHost   string
	ServerPort   string
	ClientID     string
	ClientSecret string
	RedirectPath string
	PostgresUser string
	PostgresPass string
	PostgresHost string
	PostgresPort string
	MongoName    string
	MongoHost    string
	MongoPort    string
	RedisHost    string
	RedisPort    string
}

// envDefault:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
func NewConfig() *Config {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg.ServerHost = os.Getenv("SERVER_ADDRESS")
	cfg.ServerPort = os.Getenv("SERVER_PORT")

	cfg.ClientID = os.Getenv("CLIENT_ID")
	cfg.ClientSecret = os.Getenv("CLIENT_SECRET")
	cfg.RedirectPath = os.Getenv("REDIRECT_PATH")

	cfg.PostgresUser = os.Getenv("POSTGRES_USER")
	cfg.PostgresPass = os.Getenv("POSTGRES_PASS")
	cfg.PostgresHost = os.Getenv("POSTGRES_HOST")
	cfg.PostgresPort = os.Getenv("POSTGRES_PORT")

	cfg.MongoName = os.Getenv("MONGO_NAME")
	cfg.MongoHost = os.Getenv("MONGO_HOST")
	cfg.MongoPort = os.Getenv("MONGO_PORT")

	cfg.RedisHost = os.Getenv("REDIS_HOST")
	cfg.RedisPort = os.Getenv("REDIS_PORT")

	log.Printf("configs: %+v\n", *cfg)
	return cfg
}
