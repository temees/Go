package config

type Config struct {
	Port, DB_url, Some_app_id, Jaeger_url, Sentry_url, Kafka_broker, Some_app_key string
}

func NewConfig() *Config {

	return &Config{}
}
