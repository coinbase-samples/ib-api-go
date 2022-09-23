package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port             string `mapstructure:"PORT"`
	ClientId         string `mapstructure:"COGNITO_APP_CLIENT_ID"`
	UserPoolId       string `mapstructure:"COGNITO_USER_POOL_ID"`
	GrpcPort         string `mapstructure:"GRPC_PORT"`
	OrderGrpcPort    string `mapstructure:"ORDER_GRPC_PORT"`
	Env              string `mapstructure:"ENV_NAME"`
	LogLevel         string `mapstructure:"LOG_LEVEL"`
	NetworkName      string `mapstructure:"ORDER_MGR_HOSTNAME"`
	DatabaseEndpoint string `mapstructure:"DB_ENDPOINT"`
	ProfileTableName string `mapstructure:"PROFILE_TABLE"`
	BalanceTableName string `mapstructure:"BALANCE_TABLE"`
}

func Setup(app *AppConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.AllowEmptyEnv(true)
	// set defaults
	viper.SetDefault("LOG_LEVEL", "warning")
	viper.SetDefault("PORT", "8443")
	viper.SetDefault("GRPC_PORT", "8449")
	viper.SetDefault("ORDER_GRPC_PORT", "8444")
	viper.SetDefault("ENV_NAME", "local")
	viper.SetDefault("ORDER_MGR_HOSTNAME", "localhost")
	viper.SetDefault("DB_ENDPOINT", "http://localhost:4566")
	viper.SetDefault("PROFILE_TABLE", "Profile")
	viper.SetDefault("BALANCE_TABLE", "Balance")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Missing env file %v\n", err)
	}

	err = viper.Unmarshal(&app)
	if err != nil {
		fmt.Printf("Cannot parse env file %v\n", err)
	}
}
