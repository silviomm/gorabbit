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

func ContextExists(name string) bool {
	contextSub := viper.Sub(name)
	if contextSub == nil {
		return false
	}
	return true
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

func SetCurrentContext(name string) {
	viper.Set(Config_CurrentContext, name)
	viper.WriteConfig()

	CurrentContextName = name
	context := viper.Sub(CurrentContextName)
	if context == nil {
		log.Fatal("Error reading context", CurrentContextName)
	}

	err := context.Unmarshal(&CurrentContext)
	if err != nil {
		log.Fatal("Error loading context", CurrentContextName)
	}
}

func CreateContext(name string, c RabbitMqConfig) {
	viper.Set(Config_CurrentContext, name)
	viper.Set(name+".Host", c.Host)
	viper.Set(name+".Port", c.Port)
	viper.Set(name+".AdminPort", c.AdminPort)
	viper.Set(name+".User", c.User)
	viper.Set(name+".Password", c.Password)
	viper.WriteConfig()
}
