package api

import (
	"fmt"
	"time"

	utils "gorabbit/src/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeAndSend(fromChannel *amqp.Channel, toChannel *amqp.Channel, q1 string, q2 string) {
	fmt.Println("Shovelling queue: ", q1)
	messages, err := fromChannel.Consume(
		q1,    // queue name
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // arguments
	)
	utils.FailOnError(err, "Error subscribing to "+q1)

	count := 0
	for message := range messages {
		count++
		fmt.Printf("[%s] #%d\n", q1, count)
		err = toChannel.Publish(
			message.Exchange, // exchange
			q2,               // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				ContentType:   message.ContentType,
				Body:          []byte(message.Body),
				Headers:       message.Headers,
				CorrelationId: message.CorrelationId,
				ReplyTo:       message.ReplyTo,
				Timestamp:     message.Timestamp,
				Type:          message.Type,
			})
		utils.FailOnError(err, "Failed to publish a message")
		message.Ack(true)
		fmt.Printf("[%s] #%d ok\n", q1, count)
		time.Sleep(10 * time.Second)
	}
}
