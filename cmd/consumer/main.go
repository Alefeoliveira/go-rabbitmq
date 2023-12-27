package main

import (
	"fmt"

	"github.com/Alefeoliveira/fcutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "minhafila")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) // é um acknowledge ele não faz um requeue na mensagem na fila já foi processada
	}
}
