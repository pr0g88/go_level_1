package conf

import (
	"flag"
	"fmt"
	"net"
	"net/url"
)

type Conf struct {
	port         int
	db_url       string
	jaeger_url   string
	sentry_url   string
	kafka_broker string
	some_app_id  string
	some_app_key string
}

func Flag() Conf {
	port := flag.Int("port", 8080, "port connection")
	dbUrl := flag.String("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "")
	jaegerUrl := flag.String("jaeger_url", "http://jaeger:16686", "")
	sentryUrl := flag.String("sentry_url", "http://sentry:9000", "")
	kafkaBroker := flag.String("kafka_broker", "kafka:9092", "")
	someAppId := flag.String("some_app_id", "testid", "test id to connect")
	someAppKey := flag.String("some_app_key", "testkey", "")

	return Conf{
		*port, *dbUrl, *jaegerUrl, *sentryUrl, *kafkaBroker, *someAppId, *someAppKey,
	}

}

func Valid(s *Conf) {
	_, err1 := url.Parse(s.db_url)
	if err1 != nil {
		panic(err1)
	}

	_, err2 := url.Parse(s.sentry_url)
	if err2 != nil {
		panic(err2)
	}

	_, err3 := url.Parse(s.jaeger_url)
	if err3 != nil {
		panic(err3)
	}
}
