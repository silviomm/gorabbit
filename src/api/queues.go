package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	config "gorabbit/src/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateQueue(rCh *amqp.Channel, name string) (amqp.Queue, error) {
	return rCh.QueueDeclare(name, false, false, false, false, nil)
}

func DeleteQueues(rCh *amqp.Channel, queues []string) {
	for _, q := range queues {
		rCh.QueueDelete(q, false, false, true)
	}
}

func GetQueues(config config.RabbitMqConfig) []string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+config.Host+":"+config.AdminPort+"/api/queues", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(config.User, config.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)

	var arr []map[string]interface{}
	_ = json.Unmarshal([]byte(bodyText), &arr)
	var result []string
	for _, v := range arr {
		result = append(result, v["name"].(string))
	}
	return result
}
