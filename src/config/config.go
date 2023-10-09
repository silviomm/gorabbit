package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".gorabbit")

	viper.SetDefault(Config_CurrentContext, "default")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file not found. Creating default config file at ", home)
		CreateContext("default", LocalRabbit())
		viper.SafeWriteConfig() // creates file only if it does not exists
	}

	CurrentContextName = viper.GetString(Config_CurrentContext)
	context := viper.Sub(CurrentContextName)
	if context == nil {
		log.Fatal("Error reading context", CurrentContextName)
	}

	err = context.Unmarshal(&CurrentContext)
	if err != nil {
		log.Fatal("Error loading context", CurrentContextName)
	}
}
