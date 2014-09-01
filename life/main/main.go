package main

import (
	"github.com/rvignacio/life"

	"fmt"
)

func main() {
	fmt.Println("Welcome to the game of life!")

	//game := life.NewGame(50, 50)
	//game.Play()

	life.RPentomino.Play()
}
