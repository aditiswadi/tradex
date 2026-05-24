package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppName    string
	AppEnv string
	ServerPort string
	DBUrl      string
	JWTSecret  string
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found", err)
	}
}

func loadYaml() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("config.yaml not found", err)
	}
}

func enableEnvOverride() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func Load() Config {
	loadEnv()
	loadYaml()
	enableEnvOverride()

	return Config{
		AppName: viper.GetString("app.name"),
		AppEnv:  viper.GetString("app.env"),
		ServerPort: viper.GetString("server.port"),
		DBUrl: viper.GetString("database.url"),
		JWTSecret: viper.GetString("jwt.secret"),
	}
}
