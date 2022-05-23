package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configs struct {
	ServiceHost string
	HTTPPort    string
	Version     string
	DBHost      string
	DBPort      string
	DBUsername  string
	DBName      string
	DBPassword  string
	DBSSLMode   string
}

func InitConfig() (cfg *Configs, err error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()

	if err != nil {
		return cfg, fmt.Errorf("fatal error config file: %w ", err)
	}

	if err := godotenv.Load(); err != nil {
		return cfg, fmt.Errorf("error loading env variables: %s", err.Error())
	}

	cfg = &Configs{

		ServiceHost: viper.GetString("app.server"),
		HTTPPort:    viper.GetString("app.port"),

		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		DBUsername: viper.GetString("db.username"),
		DBName:     viper.GetString("db.dbname"),
		DBSSLMode:  viper.GetString("db.sslmode"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
	return
}
