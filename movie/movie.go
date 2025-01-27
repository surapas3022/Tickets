package movie

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	_ "embed"

	"github.com/surapas3022/Tickets/utils"
)

var moviesCache []Movie
var loadError error

//go:embed cinema.json
var cinemaJSON []byte // Embed the cinema.json file

// func init() {
// 	moviesCache, loadError = LoadMovies("cinema.json")
// }

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

	if moviesCache == nil {
		return nil, errors.New("moviesCache is empty, JSON might not have loaded properly")
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
