package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/diazharizky/go-app-core/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	config.Global.SetDefault("mongodb.host", "localhost")
	config.Global.SetDefault("mongodb.port", 27017)
}

func GetClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf(
		"mongodb://%s:%d",
		config.Global.GetString("mongodb.host"),
		config.Global.GetInt("mongodb.port"),
	)

	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error unable to get MongoDB client: %v", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("error unable to establish MongoDB client connection: %v", err)
	}

	return client, nil
}
