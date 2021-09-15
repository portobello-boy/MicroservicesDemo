package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

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
	log.Print("Inserting into DB")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var eventBody map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&eventBody)
	if err != nil {
		log.Print(err)
		return
	}

	delete(eventBody, "_id")

	res, err := client.db.Collection("events").InsertOne(ctx, eventBody)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(res)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	return
}

func (client *MongoDBClient) ReadAll(w http.ResponseWriter, r *http.Request) {
	log.Print("Reading from DB")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cur, err := client.db.Collection("events").Find(ctx, bson.D{})
	if err != nil {
		log.Print(err)
		return
	}
	defer cur.Close(ctx)

	var events []interface{}
	if err = cur.All(ctx, &events); err != nil {
		log.Print(err)
		return
	}

	var resp []primitive.M
	for _, e := range events {
		resp = append(resp, e.(primitive.D).Map())
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	log.Printf("Found: %+v\n", events)
}

func (client *MongoDBClient) Read(w http.ResponseWriter, r *http.Request) {
	log.Print("Reading from DB")

	id := chi.URLParam(r, "id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var event interface{}
	err = client.db.Collection("events").FindOne(ctx, bson.M{"_id": objId}).Decode(&event)
	if err == mongo.ErrNoDocuments {
		log.Print("Nothing found")
		return
	} else if err != nil {
		log.Print(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event.(primitive.D).Map())

	log.Printf("Found: %+v\n", event)
}

func (client *MongoDBClient) Update(w http.ResponseWriter, r *http.Request) {
	log.Print("Updating in DB")

	var eventBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&eventBody)
	if err != nil {
		log.Print(err)
		return
	}

	var id primitive.ObjectID
	if eventBody["_id"] != nil {
		id, err = primitive.ObjectIDFromHex(eventBody["_id"].(string))
		if err != nil {
			log.Print(err)
			return
		}
	} else {
		log.Print("ObjectID not provided")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	delete(eventBody, "_id")

	res, err := client.db.Collection("events").ReplaceOne(ctx, bson.M{"_id": id}, eventBody)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	log.Print(res)
}

func (client *MongoDBClient) Delete(w http.ResponseWriter, r *http.Request) {
	log.Print("Deleting from DB")

	id := chi.URLParam(r, "id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := client.db.Collection("events").DeleteOne(ctx, bson.M{"_id": objId})
	if err == mongo.ErrNoDocuments {
		log.Print("Nothing found")
		return
	} else if err != nil {
		log.Print(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	log.Printf("Deleted Objects: %d\n", res.DeletedCount)
}
