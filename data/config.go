package data

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	NftPortKey string `mapstructure:"NFT_PORT_KEY"`
	InfoLog    *log.Logger
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
