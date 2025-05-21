package sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Payload struct {
	Event   string
	Date    string
	Id      string
	Payment string
}

func SendWebhook(data interface{}, url string, webhookId string) error {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling data:")
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// send the webhook request

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	// determine the status bases on the response code
	status := "failed"
	if resp.StatusCode == http.StatusOK {
		status = "delivered"
	}

	log.Println(status)

	if status == "failed" {
		return errors.New(status)
	}

	return nil
}
