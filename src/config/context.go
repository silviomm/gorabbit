package config

import (
	"log"

	"github.com/spf13/viper"
)

var CurrentContextName string
var CurrentContext RabbitMqConfig

func GetCurrentContext() RabbitMqConfig {
	return CurrentContext
}

func GetCurrentContextName() string {
	return CurrentContextName
}

func GetContext(name string) RabbitMqConfig {
	contextSub := viper.Sub(name)
	if contextSub == nil {
		log.Fatalf("Context %s not found", name)
	}
	var context RabbitMqConfig
	err := contextSub.Unmarshal(&context)
	if err != nil {
		log.Fatal("Error loading context", CurrentContextName)
	}

	return context
}

func SetDefaultContext() {
	viper.Set("default.Host", LocalRabbit().Host)
	viper.Set("default.Port", LocalRabbit().Port)
	viper.Set("default.AdminPort", LocalRabbit().AdminPort)
	viper.Set("default.User", LocalRabbit().User)
	viper.Set("default.Password", LocalRabbit().Password)
	// viper.Set("default", LocalRabbit()) // viper cant recognise it from the beggining.
	viper.WriteConfig()
}
