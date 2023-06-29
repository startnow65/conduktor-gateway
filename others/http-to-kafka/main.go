package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/kelseyhightower/envconfig"

	"github.com/hashicorp/vault/api"
)

type ForgetCustomer struct {
	Email string `json:"email"`
}

func deleteuserkey(w http.ResponseWriter, r *http.Request) {

	var customer ForgetCustomer
	json.NewDecoder(r.Body).Decode(&customer)
	defer r.Body.Close()
	log.Printf("...........")
	log.Printf("customer:  %s", customer)

	// Create a new Vault API client
	client, err := api.NewClient(&api.Config{
		Address: "http://vault:8200", // Replace with your Vault server address
	})
	if err != nil {
		panic(err)
	}

	// Set the authentication token if required
	client.SetToken("root") // Replace with your authentication token

	// Delete the KV pair
	secretPath := "/kv" // Replace with the path to your KV pair
	err = client.KVv1(secretPath).Delete(context.Background(), "/"+customer.Email)
	// resp, err := client.Logical().Delete(secretPath + "/" + customer.Email)
	if err != nil {
		panic(err)
	}

	// Check the response status
	// if resp == nil {
	// 	log.Printf("delete KV pair")
	// }

	fmt.Println("KV pair deleted successfully")

}

func main() {

	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatalln(err.Error())
	}

	kClient := NewKafkaClient(&c)
	kConsumerClient := NewKafkaConsumerClient(&c)

	// Handle produce responses and log it concurrently
	go func() {
		for e := range kClient.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message successfully to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/publish", kClient.publishHandler)
	mux.HandleFunc("/consume", kConsumerClient.consumeHandler)
	mux.HandleFunc("/delete", deleteuserkey)

	// Define a middleware function to enable CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// If it's a preflight request, just return a success status code
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}

	// Wrap your mux with the CORS middleware
	handler := corsMiddleware(mux)

	log.Printf("Starting web server using the host:port 127.0.0.1:%v", c.HttpPort)
	if err = http.ListenAndServe(fmt.Sprintf(":%s", c.HttpPort), handler); err != nil {
	}

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server close \n")
		os.Exit(1)
	} else if err != nil {
		log.Printf("error starting the server: %s\n", err.Error())
		os.Exit(1)
	}
}
