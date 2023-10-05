package models

type RabbitMqConfig struct {
	Host      string
	User      string
	Password  string
	Port      string
	AdminPort string
}

func LocalRabbit() RabbitMqConfig {
	return RabbitMqConfig{
		Host:      "localhost",
		User:      "admin",
		Password:  "admin",
		Port:      "5672",
		AdminPort: "15672",
	}
}
