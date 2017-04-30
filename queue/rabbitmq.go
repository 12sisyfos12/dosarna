package rabbitmq

import (
	"log"

	amqp "github.com/streadway/amqp"
)

// PublishMsg - Publis a message on named queue
func Publish(ch *amqp.Channel, message string) error {
	body := message
	err := ch.Publish(
		"",    // exchange
		"dd2", //q.Name, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s on queue %s", body, "queue")
	if err != nil {
		msg := "Failed to publish a message"
		log.Fatalf("%s: %s", msg, err)
		return err
	}
	return nil
}

// NewChannel - Creates a rabbit Channel
func NewChannel(hostname string, queue string) (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@" + hostname + ":5672/")
	if err != nil {
		msg := "Failed to connect to RabbitMQ"
		log.Fatalf("%s: %s", msg, err)
		return nil, err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		msg := "Failed to open a channel"
		log.Fatalf("%s: %s", msg, err)
		return nil, err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
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
		return nil, err
	}
	return ch, nil
}
