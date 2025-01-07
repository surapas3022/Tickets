package main

import (
	"encoding/json"
	"fmt"

	"github.com/surapas3022/Tickets/movie"
	"github.com/surapas3022/Tickets/ticket"
)

// func init() {
// 	fmt.Println("init:main")
// }

func main() {

	mn, err := movie.FindNameJson(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Movie Found: %+v\n", *mn)

	ticket.BuyTicket(mn.Name, mn.Price)
	movie.Review(mn.Name, 9.6)

	b, err := json.Marshal(mn)
	fmt.Printf("type : %T \n", b)
	fmt.Printf("byte : %v \n", b)
	fmt.Printf("string : %s \n", b)
	fmt.Println(err)

	// movies, err := movie.LoadMovies("cinema.json")
	// if err != nil {
	// 	fmt.Println("Error loading movies:", err)
	// 	return
	// }

	// fmt.Println(movies)

	// imdbID := "tt02"
	// mn := movie.FindName(imdbID)
	// fmt.Println(mn)
	// ticket.BuyTicket(mn)
	// movie.Review(mn, 9.54345)

}
