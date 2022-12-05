package main

import (
	"C" // to satisfy our antivirus software :(
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	topElfPosition := 0
	topElfCalories := 0

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := 1
	calories := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if calories > topElfCalories {
				topElfPosition = position
				topElfCalories = calories
			}
			position++
			calories = 0
		} else {
			number, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}

			calories += number
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if calories > topElfCalories {
		topElfPosition = position
		topElfCalories = calories
	}

	fmt.Printf("Elf #%d is carrying the most with %d calories in total", topElfPosition, topElfCalories)
}
