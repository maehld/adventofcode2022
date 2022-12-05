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
	topElfCalories := 0

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	calories := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if calories > topElfCalories {
				topElfCalories = calories
			}

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
		topElfCalories = calories
	}

	fmt.Printf("The most calories carried by an elf are %d calories", topElfCalories)
}
