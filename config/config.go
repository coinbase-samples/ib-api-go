package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port       string `mapstructure:"PORT"`
	ClientId   string `mapstructure:"COGNITO_APP_CLIENT_ID"`
	UserPoolId string `mapstructure:"COGNITO_USER_POOL_ID"`
	Env        string `mapstructure:"ENV_NAME"`
	Region     string `mapstructure:"REGION"`

	InfoLog *log.Logger
}

func Setup(app *AppConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Missing env file %v\n", err)
	}

	err = viper.Unmarshal(&app)
	if err != nil {
		fmt.Printf("Cannot parse env file %v\n", err)
	}
}
