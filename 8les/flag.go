package main

import (
	"flag"
	"fmt"

	"./config"
)

type Config struct {
	Port, DB_url, Some_app_id, Jaeger_url, Sentry_url, Kafka_broker, Some_app_key string
}

func main() {
	a := config.NewConfig()
	fmt.Println(a)
	var port = flag.String("port", "8080", "Port number")
	var db_url = flag.String("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "URL")
	var jaeger_url = flag.String("jaeger_url", "http://jaeger:16686", "URL")
	var some_app_id = flag.String("app_id", "testid", "ID")
	var sentry_url = flag.String("centry_url", " http://sentry:9000", "Sentry_URL")
	var kafka_broker = flag.String("kafka", "kafka:9092", "Kafka broker")
	var some_app_key = flag.String("key", "testkey", "Key")
	flag.Parse()

	//validation

	config := Config{
		Port:         *port,
		DB_url:       *db_url,
		Some_app_id:  *some_app_id,
		Jaeger_url:   *jaeger_url,
		Sentry_url:   *sentry_url,
		Kafka_broker: *kafka_broker,
		Some_app_key: *some_app_key,
	}
	fmt.Println(config)
}
