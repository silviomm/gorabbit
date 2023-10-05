package api

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func DeleteQueues(rCh *amqp.Channel, queues []string) {
	for _, q := range queues {
		rCh.QueueDelete(q, false, false, true)
	}
}
