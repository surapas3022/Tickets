package movie

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/surapas3022/Tickets/utils"

	_ "embed"
)

var moviesCache []Movie
var loadError error

func init() {
	moviesCache, loadError = LoadMovies("cinema.json")
	moviesCache, loadError = LoadMoviesFromEmbed()
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
	if loadError != nil {
		return nil, loadError
	}

	for _, movie := range moviesCache {
		if movie.ID == ID {
			return &movie, nil
		}
	}

	return nil, errors.New("movie not found")
}

func Review(name string, rating float64) {
	fmt.Printf("I reviewed %s and it's rating is %.2f\n", name, utils.RoundToTwoDecimalPlaces(rating))
}

// LoadMovies loads movies from a JSON file
func LoadMovies(filename string) ([]Movie, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Unmarshal into a MovieData struct
	var movieData MovieData
	if err := json.NewDecoder(file).Decode(&movieData); err != nil {
		return nil, err
	}

	// Return the movies slice
	return movieData.Data, nil
}

var cinemaJSON []byte

func LoadMoviesFromEmbed() ([]Movie, error) {
	var movieData MovieData
	if err := json.Unmarshal(cinemaJSON, &movieData); err != nil {
		return nil, err
	}
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
