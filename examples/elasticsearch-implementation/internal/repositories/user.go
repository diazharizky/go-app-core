package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/diazharizky/go-app-core/examples/elasticsearch-implementation/internal/models"
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/google/uuid"
)

type userRepository struct {
	esClient *es.TypedClient
}

const usersIndex = "users"

func NewUserRepository(esClient *es.TypedClient) userRepository {
	return userRepository{esClient}
}

func (r userRepository) List(ctx context.Context) ([]models.User, error) {
	client := r.esClient

	resp, err := client.
		Search().
		Index(usersIndex).
		Request(
			&search.Request{
				Query: &types.Query{
					MatchAll: &types.MatchAllQuery{},
				},
			},
		).
		Do(context.TODO())

	if err != nil {
		return nil, err
	}

	users := make([]models.User, len(resp.Hits.Hits))

	var user models.User
	for i, doc := range resp.Hits.Hits {
		if err = json.Unmarshal(doc.Source_, &user); err != nil {
			fmt.Printf("json unmarshal error: %v", err)
		}

		users[i] = user
	}

	return users, nil
}

func (r userRepository) Create(ctx context.Context, newUser models.User) error {
	if newUser.ID == uuid.Nil {
		newUser.ID = uuid.New()
	}

	client := r.esClient

	_, err := client.Index(usersIndex).Request(newUser).Do(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
