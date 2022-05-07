package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ae-lexs/ae_albums_api/client"
	"github.com/ae-lexs/ae_albums_api/config"
	"github.com/ae-lexs/ae_albums_api/route"
)

func main() {
	config := config.Get()
	postgresClient := client.GetPostgres(
		config.DBName,
		config.DBHost,
		config.DBPassword,
		config.DBUser,
		config.DBPort,
	)
	ginRouter := route.GetGinRouter(postgresClient)
	server := http.Server{
		Addr:         fmt.Sprintf(":%v", config.ServerPort),
		Handler:      ginRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("Server running on port %v", config.ServerPort)

		if err := server.ListenAndServe(); err != nil {
			log.Printf("Error running the server: %v", err)

			os.Exit(1)
		}
	}()

	channel := make(chan os.Signal, 1)

	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, os.Kill)

	signalChannel := <-channel

	log.Printf("Got signal: %v", signalChannel)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(timeoutContext)
}
