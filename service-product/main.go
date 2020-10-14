package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/wskurniawan/intro-microservice/service-product/config"
	"github.com/wskurniawan/intro-microservice/service-product/handler"
	"log"
	"net/http"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Panic(err)
		return
	}

	router := mux.NewRouter()

	router.Handle("/add-product", http.HandlerFunc(handler.AddMenuHandler))

	fmt.Printf("Server listen on :%s", cfg.Port)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router))
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
