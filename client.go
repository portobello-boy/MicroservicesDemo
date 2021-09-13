package main

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoDBClient struct {
	client *mongo.Client
	Close  context.CancelFunc
}

func createMongoDBClient(uri string) *mongoDBClient {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return &mongoDBClient{client, cancel}
}

func (mongo_db *mongoDBClient) create(w http.ResponseWriter, r *http.Request) {
	// mongo_db.client.Database("calendar").Collection("event").InsertOne()
}

func (mongo_db *mongoDBClient) read(w http.ResponseWriter, r *http.Request) {
	// mongo_db.client.Database("calendar").Collection("event").?
}

func (mongo_db *mongoDBClient) update(w http.ResponseWriter, r *http.Request) {
	// mongo_db.client.Database("calendar").Collection("event").ReplaceOne()
}

func (mongo_db *mongoDBClient) delete(w http.ResponseWriter, r *http.Request) {
	// mongo_db.client.Database("calendar").Collection("event").DeleteOne()
}
