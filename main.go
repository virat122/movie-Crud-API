package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id`
	Isbn     string    `json:"isbn"`
	Tittle   string    `json:"tittle"`
	Director *Director `json:"director"`
}

type Director struct {
	FristName string `json : "fristname"`
	LastName  string `json :"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") ///not get

	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // delete using append
		}
	}

	//returning remaining movie
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			////returning movie  by id
			json.NewEncoder(w).Encode(item)
			return

		}
	}

}
func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)

	newMovie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(newMovie)

}
func udpateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for index, item := range movies {
		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var updtmovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updtmovie)
			updtmovie.ID = param["id"]

			movies = append(movies, updtmovie)
			json.NewEncoder(w).Encode(updtmovie)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "47836", Tittle: "3 Ediot", Director: &Director{FristName: "jhon", LastName: "doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "47836", Tittle: " chichore", Director: &Director{FristName: "aman", LastName: "jhagale"}})
	movies = append(movies, Movie{ID: "3", Isbn: "47836", Tittle: "run", Director: &Director{FristName: "dogen", LastName: "puffet"}})
	movies = append(movies, Movie{ID: "4", Isbn: "47836", Tittle: "DON", Director: &Director{FristName: "deg", LastName: "fatahk"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", udpateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at 8080 ")
	log.Fatal(http.ListenAndServe(":8080", r)) //creating server

}
