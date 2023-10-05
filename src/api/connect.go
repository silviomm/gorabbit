package api

import (
	"context"
	"fmt"
	config "gorabbit/src/config"
	utils "gorabbit/src/utils"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectAndGetRabbitChannel(config config.RabbitMqConfig) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	// defer ch.Close()

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Println(context, cancel)
	return conn, ch
}
