package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DATABASE   = "conferences"
	COLLECTION = "conferences"
)

type MongoClient struct {
	client *mongo.Client
}

func New(mongoServerUrl string) (*MongoClient, error) {
	log.Printf("Connecting to mongo server: %s ...", mongoServerUrl)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoServerUrl))
	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("Could not ping MongoDB server at %s: %w", mongoServerUrl, err)
	}

	log.Printf("Connected to mongo server: %s", mongoServerUrl)
	return &MongoClient{client}, nil
}

func (client *MongoClient) GetConference(uniqueName string) (*ConferenceStorageModel, error) {
	collection := client.client.Database(DATABASE).Collection(COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	filter := &bson.M{"uniqueName": uniqueName}
	result := collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		return nil, fmt.Errorf("Could not find conference with uniqueName %s: %w", uniqueName, err)
	}
	var conference ConferenceStorageModel
	err = result.Decode(&conference)
	if err != nil {
		return nil, fmt.Errorf("Could not decode conference with uniqueName %s: %w", uniqueName, err)
	}
	return &conference, nil
}
func (client *MongoClient) UpdateConference(conference *ConferenceStorageModel) error {
	return nil
}

func (client *MongoClient) CreateConference(conference *ConferenceStorageModel) error {
	collection := client.client.Database(DATABASE).Collection(COLLECTION)
	marshalled, err := bson.Marshal(conference)
	if err != nil {
		return fmt.Errorf("Could not marshall conference %v, %w", *conference, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = collection.InsertOne(ctx, marshalled)
	if err != nil {
		return fmt.Errorf("Could not store conference in MongoDB %v: %w", *conference, err)
	}
	return nil
}
