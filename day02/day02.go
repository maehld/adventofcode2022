package main

import (
	"C"
	"bufio"
	"fmt"
	"log"
	"os"
)

type shape string

const (
	rock     = "Rock"
	paper    = "Paper"
	scissors = "Scissors"
)

func score(opponentShape shape, youShape shape) int {
	if opponentShape == rock {
		if youShape == rock {
			return 3 + 1
		} else if youShape == paper {
			return 6 + 2
		} else if youShape == scissors {
			return 0 + 3
		}
	} else if opponentShape == paper {
		if youShape == rock {
			return 0 + 1
		} else if youShape == paper {
			return 3 + 2
		} else if youShape == scissors {
			return 6 + 3
		}
	} else if opponentShape == scissors {
		if youShape == rock {
			return 6 + 1
		} else if youShape == paper {
			return 0 + 2
		} else if youShape == scissors {
			return 3 + 3
		}
	}

	panic("invalid shape combination")
}

func parseShapesUsingPart1Rules(text string) (opponentShape shape, youShape shape) {
	if text[0] == 'A' {
		opponentShape = rock
	} else if text[0] == 'B' {
		opponentShape = paper
	} else {
		opponentShape = scissors
	}

	if text[2] == 'X' {
		youShape = rock
	} else if text[2] == 'Y' {
		youShape = paper
	} else {
		youShape = scissors
	}

	return opponentShape, youShape
}

func parseShapesUsingPart2Rules(text string) (opponentShape shape, youShape shape) {
	if text[0] == 'A' {
		opponentShape = rock
	} else if text[0] == 'B' {
		opponentShape = paper
	} else {
		opponentShape = scissors
	}

	if text[2] == 'X' {
		if opponentShape == rock {
			youShape = scissors
		} else if opponentShape == paper {
			youShape = rock
		} else {
			youShape = paper
		}
	} else if text[2] == 'Y' {
		youShape = opponentShape
	} else {
		if opponentShape == rock {
			youShape = paper
		} else if opponentShape == paper {
			youShape = scissors
		} else {
			youShape = rock
		}
	}

	return opponentShape, youShape
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
		opponentShape, youShape := parseShapesUsingPart2Rules(text)
		totalPoints += score(opponentShape, youShape)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You scored a total of %d points", totalPoints)
}
