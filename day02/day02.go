package main

import (
	"C"
	"bufio"
	"fmt"
	"log"
	"os"
)

type Shape int

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

func score(you Shape, opponent Shape) int {
	score := int(you)

	if you == opponent {
		score += 3
	} else if you == Rock && opponent == Scissors {
		score += 6
	} else if you == Paper && opponent == Rock {
		score += 6
	} else if you == Scissors && opponent == Paper {
		score += 6
	}

	return score
}

func parseShapesUsingPart1Rules(text string) (you Shape, opponent Shape) {
	if text[0] == 'A' {
		opponent = Rock
	} else if text[0] == 'B' {
		opponent = Paper
	} else {
		opponent = Scissors
	}

	if text[2] == 'X' {
		you = Rock
	} else if text[2] == 'Y' {
		you = Paper
	} else {
		you = Scissors
	}

	return you, opponent
}

func parseShapesUsingPart2Rules(text string) (you Shape, opponent Shape) {
	if text[0] == 'A' {
		opponent = Rock
	} else if text[0] == 'B' {
		opponent = Paper
	} else {
		opponent = Scissors
	}

	if text[2] == 'X' {
		if opponent == Rock {
			you = Scissors
		} else if opponent == Paper {
			you = Rock
		} else {
			you = Paper
		}
	} else if text[2] == 'Y' {
		you = opponent
	} else {
		if opponent == Rock {
			you = Paper
		} else if opponent == Paper {
			you = Scissors
		} else {
			you = Rock
		}
	}

	return you, opponent
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPoints := 0

	for scanner.Scan() {
		text := scanner.Text()
		you, opponent := parseShapesUsingPart2Rules(text)
		totalPoints += score(you, opponent)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You scored a total of %d points\n", totalPoints)
}
