package main

import (
	"log"

	queue "github.com/12sisyfos12/dosarna/queue"
)

func failOnError(err error, msg string) {
}

func main() {

	err := queue.PublishMsg("udkmac004", "myqueue", "My message")
	if err != nil {
		msg := "Failed to send messagel"
		log.Fatalf("%s: %s", msg, err)
	}

}
