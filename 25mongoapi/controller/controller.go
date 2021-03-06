package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rupam-shil/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//update One
func updateOnerecord(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched":true}}
	result,err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result.ModifiedCount)
}

//delete 1 record
func deleteOneRecord(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"id": id}

	deleteCount, err := collection.DeleteOne(context.Background(),filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The deleted count is ",deleteCount)

}
//delete all record from  mongodb
func deleteAllMovie() int64{

	deleteResult, err := collection.DeleteMany(context.Background(),bson.D{{}},nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The deleted count is ",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//get all record from mongodb
func getAllMovies()([]primitive.M){
	cursor, err := collection.Find(context.Background(),bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies

}
// Actual controller  
func GetMyAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOnerecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func DeleteAllMovie(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	 json.NewEncoder(w).Encode(count)

}

