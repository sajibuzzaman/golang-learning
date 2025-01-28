package movieRouter

import (
	"github.com/gorilla/mux"
	movieController "github.com/sajibuzzaman/CrudApiMongoDB/controllers"
)

func MovieRouter() *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/movies", movieController.GetAllMovies).Methods("GET")
	apiRouter.HandleFunc("/movie", movieController.CreateMovie).Methods("POST")
	apiRouter.HandleFunc("/movie/{id}", movieController.GetSingleMovie).Methods("GET")
	apiRouter.HandleFunc("/movie/{id}", movieController.UpdateMovie).Methods("PUT")
	apiRouter.HandleFunc("/movie/{id}", movieController.DeleteMovie).Methods("DELETE")
	apiRouter.HandleFunc("/deleteAllMovies", movieController.DeleteAllMovies).Methods("DELETE")

	return apiRouter
}
