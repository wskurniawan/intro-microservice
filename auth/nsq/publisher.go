package nsq

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/wskurniawan/intro-microservice/auth/database"
	"log"
)

func SendEmail (auth database.Auth){
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Panic(err)
	}
	log.Println("JALAN BOS KU")
	authNSQ, _ := json.Marshal(auth)

	log.Println(string(authNSQ))
	err = p.Publish("My_NSQ_Topic", []byte(authNSQ))
	if err != nil {
		log.Panic(err)
	}
}
