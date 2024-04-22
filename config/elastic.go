package config

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func connectElastic() error {
	cert, _ := os.ReadFile("cert/http_ca.crt")
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Password: elasticPassword,
		Username: elasticUser,
		CACert:   cert,
	}
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	return nil
}
