package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassward string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()
	return Config{
		PublicHost: getEnv("Public_Host", "https://localhost"),
		Port:       getEnv("Port", ":8080"),
		DBUser:     getEnv("DB_User", "root"),
		DBPassward: getEnv("DB_Passward", "mypassward"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_Host", "127.0.0.1"), getEnv("Port", "3306")),
		DBName:     getEnv("DB_Name", "ECOM"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
