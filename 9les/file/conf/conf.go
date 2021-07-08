package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Port         string `json:"port"`
	DB_url       string `json:"db_url"`
	Some_app_id  string `json:"some_app_id"`
	Jaeger_url   string `json:"jaeger_url"`
	Sentry_url   string `json:"sentry_url"`
	Kafka_broker string `json:"kafka_broker"`
	Some_app_key string `json:"some_app_key"`
}

func (a Config) IsValid() error {
	if a.DB_url == "" {
		fmt.Println("DB URL field is empty")
	}
	return fmt.Errorf("error3")
}
func NewConfig() *Config {
	f, err := os.Open("conf/conf.json")
	if err != nil {
		log.Fatalf("error1", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	app := Config{}
	err = json.NewDecoder(f).Decode(&app)
	if err != nil {
		log.Printf("error2")
	}
	//fmt.Println(app)

	var port = flag.String("port", app.Port, "Port number")
	var db_url = flag.String("db_url", app.DB_url, "URL")
	var jaeger_url = flag.String("jaeger_url", app.Jaeger_url, "URL")
	var some_app_id = flag.String("some_app_id", app.Some_app_id, "ID")
	var sentry_url = flag.String("sentry_url", app.Sentry_url, "Sentry_URL")
	var kafka_broker = flag.String("kafka_broker", app.Kafka_broker, "Kafka broker")
	var some_app_key = flag.String("some_app_key", app.Some_app_key, "Key")
	flag.Parse()
	a := &Config{
		Port:         *port,
		DB_url:       *db_url,
		Some_app_id:  *some_app_id,
		Jaeger_url:   *jaeger_url,
		Sentry_url:   *sentry_url,
		Kafka_broker: *kafka_broker,
		Some_app_key: *some_app_key,
	}

	er := a.IsValid()
	if er != nil {
		log.Println("error4")
	}
	return a

}
