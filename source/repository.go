package source

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Select(domain string) ([]FDNS, error)
}

type repo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRepo(host string, port int, user string, pass string) (Repository, error) {
	connectionString := fmt.Sprintf("mongodb://%v:%v@%v:%v", user, pass, host, port)
	options := options.Client().ApplyURI(connectionString)
	client, err := mongo.NewClient(options)

	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	collection := client.Database("rapid_seven").Collection("dns_discovered")

	return repo{client, collection}, nil
}

func (r repo) Select(domain string) ([]FDNS, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"domain": domain}
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	data := []FDNS{}
	if err = cur.All(ctx, &data); err != nil {
		return nil, err
	}

	return data, nil
}
