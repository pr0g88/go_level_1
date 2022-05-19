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

func Valid() {

	s := "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
}
