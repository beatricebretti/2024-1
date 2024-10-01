package main

import (
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://dist:WZM8AxqFNzzh@dist049.inf.santiago.usm.cl:5672/")
	failOnError(err, "Error en conectarse con RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Error en abrir canal")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"dosh_bank", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Error en declarar cola")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Error en registrar consumidor")

	forever := make(chan bool)

	totalAmount := 0

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			amount, err := strconv.Atoi(string(d.Body))
			if err != nil {
				log.Printf("Error parsing message body: %v", err)
				continue
			}
			totalAmount += amount
			log.Printf("Total accumulated amount: %d", totalAmount)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
