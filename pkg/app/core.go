package app

import (
	"context"
	"fmt"
	"time"

	"github.com/diazharizky/go-app-core/pkg/redix"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

type Core struct {
	Info Info

	TracerProvider *tracesdk.TracerProvider

	MongoClient *mongo.Client
	RDB         *gorm.DB
	Redix       *redix.Redix
}

func (c Core) Close() error {
	if c.MongoClient != nil {
		fmt.Println("Closing MongoDB connection...")

		if err := c.MongoClient.Disconnect(context.TODO()); err != nil {
			return fmt.Errorf("error unable to close MongoDB connection: %v", err)
		}
	}

	if c.RDB != nil {
		fmt.Println("Closing RDB connection...")

		db, err := c.RDB.DB()
		if err != nil {
			return fmt.Errorf("error unable to close RDB connection: %v", err)
		}

		if err = db.Close(); err != nil {
			return fmt.Errorf("error unable to close RDB connection: %v", err)
		}
	}

	if c.Redix != nil {
		fmt.Println("Closing Redis connection...")

		if err := c.Redix.Close(); err != nil {
			return fmt.Errorf("error unable to close Redis connection: %v", err)
		}
	}

	if c.TracerProvider != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := c.TracerProvider.Shutdown(ctx); err != nil {
			return fmt.Errorf("error unable to shutdown TracerProvider: %v", err)
		}
	}

	return nil
}
