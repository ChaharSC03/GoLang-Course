package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ChaharSC03/mongodbapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

const dbName = "Netflix"
const collectionName = "watchList"

// connection

var collect *mongo.Collection

// connect with mondoDB

func init() {

	// client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
	fmt.Println("mongodb connection success")

	collect = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection istance is ready to go")

}

// mongodb helpers - files seperate for proper way

// insert 1 record

func insertOneMovie(movie model.Netflix) {

	inserted, err := collect.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie with id:", inserted.InsertedID)

}

// update one records

func updateOneMovie(movieid string) {

	id, _ := primitive.ObjectIDFromHex(movieid)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collect.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("the record is updated successfully count:", result.ModifiedCount)
}

// delete one records

func deleteOneMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCount, err := collect.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("movie got deleted with delete count", deleteCount)

}

// delete all records from db
func deleteAllMovie() int {
	filter := bson.D{{}}
	deleteResult, err := collect.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("number of movies deleted: ", deleteResult.DeletedCount)

	return int(deleteResult.DeletedCount)
}

// show all records from database

func showAllMovies() []primitive.M {

	cour, err := collect.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cour.Next(context.Background()) {

		var movie bson.M
		err := cour.Decode(&movie)

		if err != nil {

			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cour.Close(context.Background())
	return movies
}

// actual controllers - file

func GetAllMyMoviesController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allMovies := showAllMovies()

	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovieController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatchedController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteMovieController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}
func DeleteAllMoviesController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()

	json.NewEncoder(w).Encode(count)

}
