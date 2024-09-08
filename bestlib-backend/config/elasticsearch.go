package config

import (
    "github.com/elastic/go-elasticsearch/v8"
    "log"
)

var ElasticClient *elasticsearch.Client

func InitElasticClient() {
    var err error
    ElasticClient, err = elasticsearch.NewClient(elasticsearch.Config{
        Addresses: []string{"http://localhost:9200"},
    })
    if err != nil {
        log.Fatalf("Error creating the Elasticsearch client: %s", err)
    }
    log.Println("Elasticsearch client created")
}
