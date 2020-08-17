package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type person struct {
	ID        primitive.ObjectID `json:"_id,omitempy" bson:"_id,omitempty"`
	firstName string             `json:"firstname,omitempy" bson:"firstname,omitempty"`
	lastName  string             `json:"lastname,omitempy" bson:"lastname,omitempty"`
}

var client *mongo.Client

// User is a struct that group all field into a single unit

// post is a struct that group all field into a single unit

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb//localhost:2000")
	r := mux.NewRouter()
	r.HandleFunc("/person", personEndPoint).Methods("POST")
	http.ListenAndServe(":5000", r)
}

func personEndPoint(respons http.ResponseWriter, request *http.Request) {
	respons.Header().Set("Content-Type", "application/json")
	var person person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("TheNativeDeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(respons).Encode(result)
}

func getPeopleEndpoint(response http.ResponseWriter, request http.Request) {
	respons.Header().set("content-type", "application/json")
	var people []person
	collection := client.Database("TheNativeDeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `s"}`))
		return
	}
	defer cursor.Next(ctx)
	for cursor.Next(ctx) {
		var person person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `s"}`))
		return
	}
	json.NewDecoder(response).Encode(people)

}
