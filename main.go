package main

import (
	"context"
	"log"
	"net/http"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	pubsubName := "pubsub"
	topicName := "topic"
	type publishData struct {
		Data string `json:"data"`
	}

	ctx := context.Background()
	go func() {
		for {
			publishEventData := publishData{Data: "Hello, World!"}
			if err := client.PublishEvent(ctx, pubsubName, topicName, publishEventData); err != nil {
				log.Fatalf("Error publishing event: %v", err)
			}
			log.Printf("Published event to topic %s: %v", topicName, publishEventData)
			time.Sleep(1 * time.Second)
		}
	}()
	if err := r.Run(":8000"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
