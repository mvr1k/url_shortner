package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitializeConfigs() {
	fmt.Println("initializing config")
	err := readConfigFile()
	if err != nil {
		panic("Error while initializing configs shutting down")
	}
	setConfigValuesForUse()
	fmt.Println("config initialized :) ")
}

func readConfigFile() error {
	viper.SetConfigName("config")        // name of config file (without extension)
	viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./app/config/") // path to look for the config file in
	return viper.ReadInConfig()
}

func setConfigValuesForUse() {
	//setting server details
	ServerHost = viper.GetString("server_host")
	ServerPort = viper.GetInt("server_port")
}
