package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/rupam-shil/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchList"

//Most important
var collection *mongo.Collection

//connect with Mong

func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to MongoDB
	client,err :=mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo db connected")

	collection = client.Database(dbName).Collection(colName)

	//if collection is ready
	fmt.Println("Collection instace is ready")
}

// MONGODB helpers - file

// insert one record

func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(),movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db with id:", inserted.InsertedID)
}