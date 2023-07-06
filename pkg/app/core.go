package app

import (
	"context"
	"fmt"

	"github.com/diazharizky/go-app-core/pkg/redix"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Core struct {
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

	return nil
}
