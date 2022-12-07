package main

import (
	"C"
	"bufio"
	"fmt"
	"log"
	"os"
)

func splitRucksack(rucksack string) (c1 string, c2 string) {
	return rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
}

func findDuplicateItem(c1 string, c2 string) rune {
	for _, c1Item := range c1 {
		for _, c2Item := range c2 {
			if c1Item == c2Item {
				return c1Item
			}
		}
	}

	return ' '
}

func getPriority(item rune) int {
	if item > 90 {
		return int(item) - 96
	} else {
		return int(item) - 38
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPriority := 0

	for scanner.Scan() {
		text := scanner.Text()
		c1, c2 := splitRucksack(text)
		item := findDuplicateItem(c1, c2)
		priority := getPriority(item)
		totalPriority += priority
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The total priority of all mispacked items is %d\n", totalPriority)
}
