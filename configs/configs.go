package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Auth struct {
		SecretKey string
	}
	Server struct {
		ServerHost string
		ServerPort string
	}
	Mongo struct {
		MongoName string
		MongoHost string
		MongoPort string
	}
	Postgres struct {
		PostgresUser string
		PostgresPass string
		PostgresHost string
		PostgresPort string
	}
	Redis struct {
		RedisHost string
		RedisPort string
		Bucket    int
		Expiry    int
		Threshold int
	}
}

func NewConfig() *Config {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.Auth.SecretKey = os.Getenv("SECRET_KEY")

	cfg.Server.ServerHost = os.Getenv("SERVER_ADDRESS")
	cfg.Server.ServerPort = os.Getenv("SERVER_PORT")

	cfg.Postgres.PostgresUser = os.Getenv("POSTGRES_USER")
	cfg.Postgres.PostgresPass = os.Getenv("POSTGRES_PASS")
	cfg.Postgres.PostgresHost = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.PostgresPort = os.Getenv("POSTGRES_PORT")

	cfg.Mongo.MongoName = os.Getenv("MONGO_NAME")
	cfg.Mongo.MongoHost = os.Getenv("MONGO_HOST")
	cfg.Mongo.MongoPort = os.Getenv("MONGO_PORT")

	cfg.Redis.RedisHost = os.Getenv("REDIS_HOST")
	cfg.Redis.RedisPort = os.Getenv("REDIS_PORT")
	cfg.Redis.Bucket, _ = strconv.Atoi(os.Getenv("BUCKET"))
	cfg.Redis.Expiry, _ = strconv.Atoi(os.Getenv("EXPIRY"))
	cfg.Redis.Threshold, _ = strconv.Atoi(os.Getenv("THRESHOLD"))

	log.Printf("configs: %+v\n", *cfg)
	return cfg
}
