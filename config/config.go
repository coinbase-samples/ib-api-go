/**
 * Copyright 2022 - Present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// TODO move to util pkg
type BaseConfig struct {
	Env       string `mapstructure:"ENV_NAME"`
	LogLevel  string `mapstructure:"LOG_LEVEL"`
	LogToFile string `mapstructure:"LOG_TO_FILE"`
	Region    string `mapstructure:"REGION"`
	Port      string `mapstructure:"PORT"`
	GrpcPort  string `mapstructure:"GRPC_PORT"`
}

func UnmarshalBase(b *BaseConfig) {
	err := viper.Unmarshal(&b)
	if err != nil {
		fmt.Printf("Cannot parse env file %v\n", err)
	}
}

type AppConfig struct {
	BaseConfig
	ClientId   string `mapstructure:"COGNITO_APP_CLIENT_ID"`
	UserPoolId string `mapstructure:"COGNITO_USER_POOL_ID"`

	OrderGrpcPort    string `mapstructure:"ORDER_GRPC_PORT"`
	OrderRouteId     string `mapstructure:"ORDER_MGR_ROUTE_ID"`
	OrderMgrHostname string `mapstructure:"ORDER_MGR_HOSTNAME"`

	UserGrpcPort    string `mapstructure:"USER_GRPC_PORT"`
	UserMgrHostname string `mapstructure:"USER_MGR_HOSTNAME"`
	UserRouteId     string `mapstructure:"USER_MGR_ROUTE_ID"`

	DatabaseEndpoint  string `mapstructure:"DB_ENDPOINT"`
	BalanceTableName  string `mapstructure:"BALANCE_TABLE"`
	AssetTableName    string `mapstructure:"ASSET_TABLE"`
	ActivityTableName string `mapstructure:"ACTIVITY_TABLE"`

	RedisEndpoint string `mapstructure:"REDIS_ENDPOINT_ADDRESS"`
	RedisPort     string `mapstructure:"REDIS_ENDPOINT_PORT"`

	ExternalHostName    string `mapstructure:"EXTERNAL_HOST_NAME"`
	ExternalApiHostName string `mapstructure:"EXTERNAL_API_HOST_NAME"`
}

func (a AppConfig) IsLocalEnv() bool {
	return a.BaseConfig.Env == "local"
}

func Setup(app *AppConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.AllowEmptyEnv(true)
	// set defaults
	viper.SetDefault("ENV_NAME", "local")
	viper.SetDefault("LOG_LEVEL", "warning")
	viper.SetDefault("LOG_TO_FILE", "false")
	viper.SetDefault("REGION", "us-east-1")
	viper.SetDefault("PORT", "8443")
	viper.SetDefault("GRPC_PORT", "8449")

	viper.SetDefault("ORDER_GRPC_PORT", "8444")
	viper.SetDefault("ORDER_MGR_HOSTNAME", "localhost")
	viper.SetDefault("ORDER_MGR_ROUTE_ID", "ordermgr")

	viper.SetDefault("USER_GRPC_PORT", "8451")
	viper.SetDefault("USER_MGR_HOSTNAME", "localhost")
	viper.SetDefault("USER_MGR_ROUTE_ID", "usermgr")

	viper.SetDefault("DB_ENDPOINT", "http://localhost:4566")
	viper.SetDefault("BALANCE_TABLE", "Balance")
	viper.SetDefault("ASSET_TABLE", "Asset")
	viper.SetDefault("ACTIVITY_TABLE", "Activity")

	viper.SetDefault("REDIS_ENDPOINT_ADDRESS", "127.0.0.1")
	viper.SetDefault("REDIS_ENDPOINT_PORT", "7000")

	viper.SetDefault("EXTERNAL_HOST_NAME", "http://localhost:4200")
	viper.SetDefault("EXTERNAL_API_HOST_NAME", "https://localhost:4200")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Missing env file %v\n", err)
	}

	err = viper.Unmarshal(&app)
	if err != nil {
		fmt.Printf("Cannot parse env file %v\n", err)
	}
	UnmarshalBase(&app.BaseConfig)
}
