package movieController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	movieModel "github.com/sajibuzzaman/CrudApiMongoDB/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://sajibbutterfly:txNP2JeLuwjZgX0i@cluster0.amfct.mongodb.net"
const dbName = "Netflix"
const colName = "movie"

var collection *mongo.Collection

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func GetSingleMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	params := mux.Vars(r)

	movie, err := getSingleMovie(params["id"])
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Movie not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to fetch movie: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie movieModel.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}

	if err := createMovie(&movie); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	movieId := params["id"]

	var movie primitive.M
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}

	updatedMovie := updateMovie(movieId, movie)

	json.NewEncoder(w).Encode(updatedMovie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	movieId := params["id"]

	deleteMovie(movieId)

	json.NewEncoder(w).Encode(movieId)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()

	json.NewEncoder(w).Encode(count)
}

// *** Service ***//
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie primitive.M
		if err := cur.Decode(&movie); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies

}

func getSingleMovie(movieId string) (movieModel.Movie, error) {
	id, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		return movieModel.Movie{}, fmt.Errorf("invalid movie ID format: %v", err)
	}

	filter := bson.M{"_id": id}
	var movie movieModel.Movie
	if err := collection.FindOne(context.Background(), filter).Decode(&movie); err != nil {
		return movieModel.Movie{}, err
	}

	return movie, nil
}

func createMovie(movie *movieModel.Movie) error {
	res, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	movie.ID = res.InsertedID.(bson.ObjectID)
	fmt.Println("Inserted 1 movie in db successfully", res.InsertedID)
	return nil

}

func updateMovie(movieId string, movie primitive.M) primitive.M {
	id, _ := bson.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": movie}

	var updatedMovie primitive.M
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&updatedMovie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated 1 movie in db successfully")
	return updatedMovie
}

func deleteMovie(movieId string) {
	id, _ := bson.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted 1 movie in db successfully", res.DeletedCount)
}

func deleteAllMovies() int64 {
	res, err := collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v movies in db successfully", res.DeletedCount)
	return res.DeletedCount
}
