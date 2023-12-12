package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func filereader() string {
	inputFile, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("file err")
		return ""
	}
	inputString := string(inputFile)
	return inputString
}

func main() {
	input := filereader()
	// p1(input)
	p1(input)
	p2(input)

}

func p1(puzzle string) {
	MAX_RED, MAX_BLUE, MAX_GREEN := 12, 14, 13

	type Round struct {
		red   int
		green int
		blue  int
	}
	type Game struct {
		id     int
		rounds []Round
	}

	// split to lines
	puzzles := strings.Split(puzzle, "\n")
	var games []Game
	for _, val := range puzzles {
		gamestage := strings.Split(val, ":")
		var id int
		// parse game id
		fmt.Sscanf(gamestage[0], "Game %d", &id)
		// parse rounds
		game := Game{}
		game.id = id
		for _, value := range strings.Split(gamestage[1], ";") {
			// parse rounds
			r := Round{}
			for _, ball := range strings.Split(value, ",") {
				num, color := 0, ""
				fmt.Sscanf(ball, "%d %s", &num, &color)
				switch color {
				case "red":
					r.red = num
				case "blue":
					r.blue = num
				case "green":
					r.green = num
				default:
					fmt.Println("ERR!!")
					return
				}
			}
			game.rounds = append(game.rounds, r)

		}
		games = append(games, game)

	}

	for _, game := range games {
		fmt.Print("gid: ", game.id, "\n")
		for _, r := range game.rounds {
			fmt.Printf(" r: %d, b: %d, g: %d\n", r.red, r.blue, r.green)
		}
		fmt.Println()
	}

	//  start simulate
	possibleIdSum := 0
	for _, game := range games {
		isValid := true
		for _, r := range game.rounds {
			// impossible case handler
			if r.red > MAX_RED || r.blue > MAX_BLUE || r.green > MAX_GREEN {
				isValid = false
			}
		}

		if isValid {
			possibleIdSum += game.id
		}
	}
	fmt.Println(possibleIdSum)

}

func p2(puzzle string) {

	type Round struct {
		red   int
		green int
		blue  int
	}
	type Game struct {
		id     int
		rounds []Round
	}

	// split to lines
	puzzles := strings.Split(puzzle, "\n")
	var games []Game
	for _, val := range puzzles {
		gamestage := strings.Split(val, ":")
		var id int
		// parse game id
		fmt.Sscanf(gamestage[0], "Game %d", &id)
		// parse rounds
		game := Game{}
		game.id = id
		for _, value := range strings.Split(gamestage[1], ";") {
			// parse rounds
			r := Round{}
			for _, ball := range strings.Split(value, ",") {
				num, color := 0, ""
				fmt.Sscanf(ball, "%d %s", &num, &color)
				switch color {
				case "red":
					r.red = num
				case "blue":
					r.blue = num
				case "green":
					r.green = num
				default:
					fmt.Println("ERR!!")
					return
				}
			}
			game.rounds = append(game.rounds, r)

		}
		games = append(games, game)

	}

	for _, game := range games {
		fmt.Print("gid: ", game.id, "\n")
		for _, r := range game.rounds {
			fmt.Printf(" r: %d, b: %d, g: %d\n", r.red, r.blue, r.green)
		}
		fmt.Println()
	}

	//  start simulate
	possibleIdSum := 0
	for _, game := range games {
		max_r, max_g, max_b := 0, 0, 0
		for _, r := range game.rounds {
			if r.red > max_r {
				max_r = r.red
			}
			if r.blue > max_b {
				max_b = r.blue
			}
			if r.green > max_g {
				max_g = r.green
			}

		}
		possibleIdSum += max_b * max_g * max_r

	}
	fmt.Println(possibleIdSum)

}

// https://discord.com/channels/453113677730545666/453113677730545668/1183772210532454473
func tmp() {
	rounds := 0
	coin := 1000000
	coingain := 0
	cdtime := 0
	// set rand seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for coin > 0 && rounds < 10000000 {
		rounds++
		if cdtime > 0 {
			cdtime--
			continue

		} else {
			coin--
			dice := rng.Intn(6) + 1
			coin += dice
			coingain += dice
			cdtime = dice
		}
	}

	fmt.Println("ratio: ", float64(coingain)/float64(rounds), "coin: ", coin, " coingain: ", coingain, "rounds: ", rounds)
}
