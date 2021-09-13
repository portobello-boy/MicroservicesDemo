package server

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBClient struct {
	client *mongo.Client
	Close  context.CancelFunc
}

func CreateMongoDBClient(uri string) *MongoDBClient {
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

	return &MongoDBClient{client, cancel}
}

func (mongo_db *MongoDBClient) Create(w http.ResponseWriter, r *http.Request) {
	// mongo_db.client.Database("calendar").Collection("event").InsertOne()
}

func (mongo_db *MongoDBClient) Read(w http.ResponseWriter, r *http.Request) {
	// mongo_db.client.Database("calendar").Collection("event").Find()
}

func (mongo_db *MongoDBClient) Update(w http.ResponseWriter, r *http.Request) {
	// mongo_db.client.Database("calendar").Collection("event").ReplaceOne()
}

func (mongo_db *MongoDBClient) Delete(w http.ResponseWriter, r *http.Request) {
	// mongo_db.client.Database("calendar").Collection("event").DeleteOne()
}
