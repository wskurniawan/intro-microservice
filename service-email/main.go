package main

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
	"net/smtp"
	"sync"
)

func main() {
	//send("hello there")
	nsqConsumer()
}

func send(username string, token string) {
	from := "digtalent.tester@gmail.com"
	pass := "digtalentjaya123"
	to := "fadhlangaming@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		"Hello " + username + " here's your token" + token

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}

func nsqConsumer(){
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("My_NSQ_Topic", "My_NSQ_Channel", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}
	//c.MaxInFlight defaults to 1

	type Auth struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var auth Auth

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("NSQ message received:")
		err = json.Unmarshal(message.Body, &auth)
		log.Println("ERROR : ",err)
		log.Println(auth)
		send(auth.Username,auth.Password)
		return nil
	}))

	err = c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
	wg.Wait()
}