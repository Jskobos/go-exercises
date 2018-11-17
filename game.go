package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Game struct; allows Game to communicate with main routing
type Game struct {
	GameChannel chan string
	questions   int
	score       int
}

// NewGame instantiates a game that listens on the channel.
func NewGame() *Game {
	return &Game{
		GameChannel: make(chan string),
		questions:   0,
		score:       0,
	}
}

// Play initiates the game loop
func Play(g *Game) {
	fmt.Println("Starting game")
	csvFile, _ := os.Open("problems.csv")
	r := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := r.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		question := line[0]
		answer, _ := strconv.Atoi(line[1])
		g.questions++

		fmt.Println(question)

		var input int
		fmt.Scan(&input)
		var result string
		if input == answer {
			result = "Correct!"
			g.score++
		} else {
			result = "Incorrect, answer was" + strconv.Itoa(answer)
		}
		g.GameChannel <- result
	}
}

// EndGame tallies the score at the end of the game
func EndGame(g *Game) {
	fmt.Println("You answered",
		strconv.Itoa(g.score),
		"out of",
		strconv.Itoa(g.questions),
		"questions correctly!")
}
