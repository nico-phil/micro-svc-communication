package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	Port = "8080"
)

type Config struct{}

func main() {
	app := Config{}

	log.Printf("starting broker service on port %s\n", Port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
