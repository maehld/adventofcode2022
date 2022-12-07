package main

import (
	"C" // to satisfy our antivirus software :(
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
import "sort"

func main() {
	caloriesSums := make([]int, 0)

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
			caloriesSums = append(caloriesSums, calories)
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

	sort.Ints(caloriesSums)

	n := 3
	topNCalories := 0
	for _, caloriesSum := range caloriesSums[len(caloriesSums)-3:] {
		topNCalories += caloriesSum
	}

	fmt.Printf("The most calories carried by the top %d elves are %d calories\n", n, topNCalories)
}
