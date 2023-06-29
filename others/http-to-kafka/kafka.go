package main

import (
	"context"
	"crypto/rand"
	"reflect"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
	customerv1beta2 "github.com/kanapuli/http-to-kafka/hellofresh/stream/customer/v1beta2"
)

type KafkaClient struct {
	producer         *kafka.Producer
	consumer         *kafka.Consumer
	topic            string
	bootstrapServers string
	serializer       *protobuf.Serializer
	deserializer     *protobuf.Deserializer
}

func NewKafkaClient(c *Config) *KafkaClient {
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(c.SchemaRegistryUrl))
	if err != nil {
		log.Fatalln(err)
	}

	seConfig := protobuf.NewSerializerConfig()
	seConfig.AutoRegisterSchemas = true
	seConfig.NormalizeSchemas = false
	seConfig.UseLatestVersion = true

	serializer, err := protobuf.NewSerializer(client, serde.ValueSerde, seConfig)
	if err != nil {
		log.Fatalln(err)
	}

	serializer.SubjectNameStrategy = schemaSubjectNameStrategy
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", c.KafkaServer, c.KafkaPort),
	})
	if err != nil {
		log.Fatalln(err)
	}

	return &KafkaClient{
		producer:         p,
		topic:            c.Topic,
		bootstrapServers: fmt.Sprintf("%s:%s", c.KafkaServer, c.KafkaPort),
		serializer:       serializer,
	}
}

func schemaSubjectNameStrategy(topic string, serdeType serde.Type, schema schemaregistry.SchemaInfo) (string, error) {
	fmt.Printf("trying to resolve for topic: %s schema name: %s\n", topic, schema.Schema)
	// For hackaton only.. loook away!
	return "hellofresh/stream/customer/v1beta2/customer.proto", nil
}

func msgFactory(subject string, name string) (interface{}, error) {
	fmt.Printf("trying to construct message for subject: %s name: %s\n", subject, name)
	return &customerv1beta2.CustomerValue{}, nil
}

func NewKafkaConsumerClient(c *Config) *KafkaClient {
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig("http://schema-registry:8081"))
	if err != nil {
		log.Fatalln(err)
	}

	desConfig := protobuf.NewDeserializerConfig()
	deserializer, err := protobuf.NewDeserializer(client, serde.ValueSerde, desConfig)
	deserializer.SubjectNameStrategy = schemaSubjectNameStrategy
	if err != nil {
		log.Fatalln(err)
	}

	deserializer.MessageFactory = msgFactory
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", c.KafkaServer, c.KafkaPort),
		"group.id":          "my-group-1",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	return &KafkaClient{
		consumer:         consumer,
		topic:            c.Topic,
		bootstrapServers: fmt.Sprintf("%s:%s", c.KafkaServer, c.KafkaPort),
		deserializer:     deserializer,
	}
}

type ConsumedMessage struct {
	Topic     string       `json:"topic"`
	Partition int32        `json:"partition"`
	Offset    kafka.Offset `json:"offset"`
	Key       []byte       `json:"key"`
	Value     []byte       `json:"value"`
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := len(charset)
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	for i := 0; i < length; i++ {
		randomBytes[i] = charset[int(randomBytes[i])%charsetLength]
	}

	return string(randomBytes)
}
func (k *KafkaClient) consumeHandler(w http.ResponseWriter, r *http.Request) {

	consumerGroupID := fmt.Sprintf("group-%s", generateRandomString(5))
	log.Printf("consumerGroupID: %s\n", string(consumerGroupID))

	log.Printf("k.bootstrapServers: %s\n", string(k.bootstrapServers))

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": k.bootstrapServers,
		"group.id":          consumerGroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Printf("Failed to create consumer: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		resp := fmt.Sprintf("Failed to create Kafka consumer: %v", err)
		fmt.Fprintf(w, `{"error": "%v"}`, resp)
		return
	}

	c.SubscribeTopics([]string{k.topic}, nil)

	defer func() {
		c.Close()
	}()
	//consumedMessages := []ConsumedMessage{}
	customers := []Customer{}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second) // Set a timeout of 30 seconds
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			// Request canceled or timed out
			log.Println("Request canceled or timed out.")
			responseJSON, err := json.Marshal(customers)

			if err != nil {
				log.Printf("Failed to marshal response to JSON: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("Content-Type", "application/json")
				resp := fmt.Sprintf("Failed to marshal response to JSON: %v", err)
				fmt.Fprintf(w, `{"error": "%v"}`, resp)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
			return
		default:
			//log.Printf("Are you working...")
			ev := c.Poll(100)
			if ev == nil {
				log.Printf("cv is null")
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				log.Printf("Received message: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s\n",
					*e.TopicPartition.Topic, e.TopicPartition.Partition, e.TopicPartition.Offset, string(e.Key), string(e.Value))
				consumedMsg := ConsumedMessage{
					Topic:     *e.TopicPartition.Topic,
					Partition: e.TopicPartition.Partition,
					Offset:    e.TopicPartition.Offset,
					Key:       e.Key,
					Value:     e.Value,
				}
				log.Printf("decoding %s\n", string(consumedMsg.Value))
				if err != nil {
					fmt.Println("Error decoding Base64:", err)
					continue
				}
				// var customer customerv1beta2.CustomerValue

				msg, err := k.deserializer.Deserialize(consumedMsg.Topic, consumedMsg.Value)
				if err != nil {
					fmt.Println("Error deserializing proto message:", err)
					continue
				}

				log.Printf("decoded %+v\n", msg)
				log.Printf("decoded %+v\n", reflect.TypeOf(msg).String())

				customer := msg.(*customerv1beta2.CustomerValue)
				customers = append(customers, Customer{
					Name:          customer.Name,
					Email:         customer.Email,
					Height:        customer.Height,
					FavouriteFood: customer.FavouriteFood,
				})
			case kafka.Error:
				log.Printf("Consumer error: %v (%v)\n", e.Code(), e)
			}
		}
	}

}

func (k *KafkaClient) publishHandler(w http.ResponseWriter, r *http.Request) {

	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)
	defer r.Body.Close()
	log.Printf("...........")
	log.Printf("customer:  %s", customer)
	msg, err := k.prepareKafkaMessage(customer)
	log.Printf("msg:  %s", msg)
	if err != nil {
		log.Printf("error marshaling proto message: %v\n", err)
	}

	kMsg := &kafka.Message{
		Key:   []byte(customer.Email),
		Value: msg,
		TopicPartition: kafka.TopicPartition{
			Topic:     &k.topic,
			Partition: kafka.PartitionAny,
		},
	}

	err = k.produce(kMsg)
	if err != nil {
		log.Printf("error producing kafka message: %v", err)
		w.WriteHeader(500)
		w.Header().Add("Content-type", "application/json")
		resp := fmt.Sprintf("error producing kafka messaga: %v", err)
		fmt.Fprintf(w, `{"msg": "%v"}`, resp)
		return
	}

	w.WriteHeader(200)
	w.Header().Add("Content-type", "application/json")
	resp := fmt.Sprintf("message produced successfully to kafka for the customer: %v", customer.Name)
	fmt.Fprintf(w, `{"msg": "%v"}`, resp)
}

func (k *KafkaClient) produce(msg *kafka.Message) error {
	return k.producer.Produce(msg, nil)
}

// Converts the Customer object from the http payload
// to protobuf object
func (k *KafkaClient) prepareKafkaMessage(customer Customer) ([]byte, error) {
	msg := &customerv1beta2.CustomerValue{
		Name:          customer.Name,
		Email:         customer.Email,
		Height:        customer.Height,
		FavouriteFood: customer.FavouriteFood,
	}

	return k.serializer.Serialize(k.topic, msg)
}
