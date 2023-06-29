package main

type Config struct {
	KafkaServer       string `envconfig:"KAFKA_SERVER"`
	KafkaPort         string `envconfig:"KAFKA_PORT"`
	SchemaRegistryUrl string `envconfig:"SCHEMA_REGISTRY_URL"`
	Topic             string `envconfig:"KAFKA_TOPIC"`
	HttpPort          string `envconfig:"HTTP_PORT" default:"7070"`
}
