package elasticsearch

import (
	"fmt"

	"github.com/diazharizky/go-app-core/config"
	es "github.com/elastic/go-elasticsearch/v8"
)

func init() {
	config.Global.SetDefault("elasticsearch.hosts", []string{"http://localhost:9200"})
}

func GetClient() (*es.TypedClient, error) {
	esConfig := es.Config{
		Addresses: config.Global.GetStringSlice("elasticsearch.hosts"),
	}

	client, err := es.NewTypedClient(esConfig)
	if err != nil {
		return nil,
			fmt.Errorf("error unable to get Elasticsearch client: %v", err)
	}

	return client, nil
}
