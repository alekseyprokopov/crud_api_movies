package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var movies []Movie = []Movie{}

func main() {
	movie1 := Movie{Id: "1", Title: "movie1", Director: &Director{
		FirstName: "ahalay",
		LastName:  "mahalay",
	}}

	movie2 := Movie{Id: "2", Title: "movie2", Director: &Director{
		FirstName: "firstName2",
		LastName:  "lastName2",
	}}
	movies = append(movies, movie1, movie2)

	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	router.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)
	router.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)

	http.ListenAndServe(":3000", router)
	log.Println("starting serve...")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	result, _ := json.Marshal(movies)
	w.Write(result)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if params["id"] == item.Id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if params["id"] == item.Id {
			movies = append(movies[:i], movies[i+1:]...)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	params := mux.Vars(r)
	for i, item := range movies {
		if params["id"] == item.Id {
			movie.Id = params["id"]
			movies = append(movies[:i], movies[i+1:]...)
			movies = append(movies, movie)
			return
		}
	}
}
