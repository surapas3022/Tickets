package movie

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	_ "embed"

	"github.com/surapas3022/Tickets/utils"
)

// Embedded JSON data
//
//go:embed cinema.json
var cinemaJSON []byte

var MoviesCache []Movie
var LoadError error

func init() {
	// MoviesCache, LoadError = LoadMovies("cinema.json")
	MoviesCache, LoadError = LoadMoviesFromBytes(cinemaJSON)
}

func FindName(imdbID string) string {
	switch imdbID {
	case "tt01":
		return "Avenger Infinities Wars"
	case "tt02":
		return "Iron man"
	}
	return "Not Found."
}

func FindNameJson(ID int) (*Movie, error) {
	if LoadError != nil {
		return nil, LoadError
	}

	if MoviesCache == nil {
		return nil, errors.New("MoviesCache is empty, JSON might not have loaded properly")
	}

	for _, movie := range MoviesCache {
		if movie.ID == ID {
			return &movie, nil
		}
	}

	return nil, errors.New("movie not found")
}

func Review(name string, rating float64) {
	fmt.Printf("I reviewed %s and it's rating is %.2f\n", name, utils.RoundToTwoDecimalPlaces(rating))
}

func LoadMovies(filename string) ([]Movie, error) {
	fmt.Println("Attempting to load file:", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var movieData MovieData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&movieData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	fmt.Println("Successfully loaded movies:", movieData)
	return movieData.Data, nil
}

func LoadMoviesFromBytes(data []byte) ([]Movie, error) {
	var movieData MovieData
	if err := json.Unmarshal(data, &movieData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	fmt.Println("Successfully loaded movies:", movieData)
	return movieData.Data, nil
}

type Movie struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	ImdbID string  `json:"imdbID"`
	Rating float64 `json:"rating"`
	Price  string  `json:"price"`
}

type MovieData struct {
	Data []Movie `json:"data"`
}
