package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// PublishMsg - Publis a message on named queue
func PublishMsg(hostname string, queue string, message string) error {
	conn, err := amqp.Dial("amqp://guest:guest@" + hostname + ":5672/")
	if err != nil {
		msg := "Failed to connect to RabbitMQ"
		log.Fatalf("%s: %s", msg, err)
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		msg := "Failed to open a channel"
		log.Fatalf("%s: %s", msg, err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		msg := "Failed to declare a queue"
		log.Fatalf("%s: %s", msg, err)
		return err
	}

	body := message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s on queue %s", body, queue)
	if err != nil {
		msg := "Failed to publish a message"
		log.Fatalf("%s: %s", msg, err)
		return err
	}
	return nil
}
