package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AlperSeyman/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Encoding = Turning a Go struct into JSON or BSON (to send to MongoDB, or to an API).
// Decoding = Turning JSON/BSON from MongoDB into a Go struct.

const connectionString = ""
const dbName = "netflixDB"
const collectionName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB

func init() {

	// client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)
}

// MONGODB helpers

// get all records
func getAllMovies() []models.Netflix {

	filter := bson.M{} // get everthing

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var movies []models.Netflix             // holds all movies, var movies []primitive.M
	for cursor.Next(context.Background()) { // moves to the next document in the list (just like reading from a file)
		var movie models.Netflix     // holds one movie temporarly, var movie bson.M
		err := cursor.Decode(&movie) // decode document into struct
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie) // add a result
	}

	return movies

}

// get 1 record
func getOneMovie(movieID string) {

	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	// Filter to match the document want to update
	filter := bson.M{"_id": id} // find document by _id

	var movie models.Netflix
	err2 := collection.FindOne(context.Background(), filter).Decode(&movie)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(movie)
}

// insert 1 record
func insertOneMovie(movie models.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieID string) { // newTitle string

	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	// Filter to match the document want to update
	filter := bson.M{"_id": id} // find document by _id

	// What you want to update
	/* update := bson.M{"$set": bson.M{"watched": newTitle}} // change the movie name */
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal("Updated failed", err)
	}

	fmt.Printf("Updated %v document(s)\n", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(moveID string) {

	id, err := primitive.ObjectIDFromHex(moveID)
	if err != nil {
		log.Fatal(err)
	}

	// Filter to match the document want to delete
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v document(s)\n", result.DeletedCount)

}

// delete all records from mongodb
func deleteAllMovie() int64 {

	filter := bson.M{} // empty filter = match all records

	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents", result.DeletedCount)
	return result.DeletedCount
}

// Actual Controller - file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	vars := mux.Vars(r)
	movieID := vars["id"]
	getOneMovie(movieID)
	json.NewEncoder(w).Encode(movieID)
}

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// grab id from request
	vars := mux.Vars(r)
	movieID := vars["id"]
	updateOneMovie(movieID)
	json.NewEncoder(w).Encode(movieID)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	vars := mux.Vars(r)
	movieID := vars["id"]
	deleteOneMovie(movieID)
	json.NewEncoder(w).Encode(movieID)
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}
