package server

import (
	"context"
	"log"
	"net/http"

	"github.com/portobello-boy/MicroservicesDemo/CRUD/structures"

	"github.com/go-chi/chi"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBClient struct {
	db    *mongo.Database
	Close context.CancelFunc
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

	return &MongoDBClient{client.Database("calendar"), cancel}
}

func (client *MongoDBClient) Create(w http.ResponseWriter, r *http.Request) {
	// client.db.Collection("events").InsertOne()
}

func (client *MongoDBClient) Read(w http.ResponseWriter, r *http.Request) {
	log.Print("Reading from DB")

	id := chi.URLParam(r, "id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//handle error
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var event structures.Event
	err = client.db.Collection("events").FindOne(ctx, bson.M{"_id": objId}).Decode(&event)
	if err == mongo.ErrNoDocuments {
		//handle the no record found
	} else if err != nil {
		//handle errors
	}

	log.Printf("Found: %+v\n", event)
}

func (client *MongoDBClient) Update(w http.ResponseWriter, r *http.Request) {
	// client.db.Collection("events").ReplaceOne()
}

func (client *MongoDBClient) Delete(w http.ResponseWriter, r *http.Request) {
	// client.db.Collection("events").DeleteOne()
}
