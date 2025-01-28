package main

import (
	"fmt"
	"log"
	"net/http"

	movieRouter "github.com/sajibuzzaman/CrudApiMongoDB/routes"
)

func main() {
	fmt.Println("Welcome to CRUD API")
	r := movieRouter.MovieRouter()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("⚙️ Server is running at port : 4000")
}
