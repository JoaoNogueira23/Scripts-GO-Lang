package queue

import (
	"context"
	"gateway-payments/sender"
	"log"
	"time"

	redisClient "gateway-payments/redis"
)

func ProcessWebhooks(ctx context.Context, webhookQueue chan redisClient.WebhookPayload) {

	for payload := range webhookQueue {
		go func(p redisClient.WebhookPayload) {

			backOffTime := time.Second
			maxBackofftime := time.Hour

			retries := 0
			maxRetries := 5

			for {
				err := sender.SendWebhook(p.Data, p.Url, p.WebhookId)

				if err == nil {
					break
				}

				retries++
				if retries >= maxRetries {
					log.Println("Max retries reached. |Giving up on webhook:", p.WebhookId)
					break
				}

				time.Sleep(backOffTime)
				backOffTime *= 2 // double the beckoff time for the next interation
				if backOffTime > maxBackofftime {
					backOffTime = maxBackofftime // cap the backoff time to maxBackoffTime
				}
			}
		}(payload)
	}
}
