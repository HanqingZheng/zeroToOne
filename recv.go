package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main(){
	conn, err := amqp.Dial("amqp://admin:123456@192.168.3.109:5672/thost")
	if err != nil {
		panic(err)
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()

	     defer ch.Close()

	     q, err := ch.QueueDeclare(
		         "hello", // name
		         false,   // durable
		         false,   // delete when unused
		         false,   // exclusive
		         false,   // no-wait
		         nil,     // arguments
		     )


	     msgs, err := ch.Consume(
		         q.Name, // queue
		         "",     // consumer
		         true,   // auto-ack
		         false,  // exclusive
		         false,  // no-local
		         false,  // no-wait
		         nil,    // args
		     )


	     forever := make(chan bool)

   go func() {
		       for d := range msgs {

			          log.Printf("Received a message: %s", d.Body)
			         }
		  }()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	 <-forever
}



