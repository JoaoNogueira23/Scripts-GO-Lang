package redis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

type WebhookPayload struct {
	Url       string `json:"url"`
	WebhookId string `json:"webhookId"`
	Data      struct {
		Id      string `json:"id"`
		Payment string `json:"payment"`
		Event   string `json:"event"`
		Date    string `json:"date"`
	} `json:"data"`
}

func Subscribe(ctx context.Context, rdb *redis.Client, webhookQueue chan WebhookPayload) error {
	// Subscribe to the "webhooks" channel in redis
	pubsub := rdb.Subscribe(ctx, "payments")

	// Ensure that the pubSub connection is closed when the function exits
	defer func(pubSub *redis.PubSub) {
		err := pubSub.Close()
		if err != nil {
			log.Println("Error closing pubsub:", err)
		}
	}(pubsub)

	var payload WebhookPayload

	// infinite loop to receive messages from the "webhooks" channel
	for {
		// receivbe a message from the channel
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			return err
		}

		// unmarshal the message into the payload struct
		err = json.Unmarshal([]byte(msg.Payload), &payload)
		if err != nil {
			log.Println("Error unmarshalling payload:", err)
			continue // continue with the next message if there's an error
		}

		webhookQueue <- payload // Sending the payload to the channel
	}

}
