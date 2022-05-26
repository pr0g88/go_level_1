package config

import (
	"encoding/json"
	"net/url"
	"os"
)

type Conf struct {
	Port        int    `json:"port"`
	DbURL       string `json:"db_url"`
	JaegerURL   string `json:"jaeger_url"`
	SentryURL   string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppID   string `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

func ReadJson() {
	file, err := os.Open("file.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	app := Conf{}

	err1 := json.NewDecoder(file).Decode(app)
	if err1 != nil {
		panic(err1)
	}

}

func Valid(s *Conf) {
	_, err1 := url.Parse(s.DbURL)
	if err1 != nil {
		panic(err1)
	}

	_, err2 := url.Parse(s.SentryURL)
	if err2 != nil {
		panic(err2)
	}

	_, err3 := url.Parse(s.JaegerURL)
	if err3 != nil {
		panic(err3)
	}
}
