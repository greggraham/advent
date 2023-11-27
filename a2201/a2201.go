package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	part2(scanner)
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}

func part1(scanner *bufio.Scanner) {
	maxCalories := 0
	currentCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		} else {
			number, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += int(number)
		}
	}
	fmt.Println(maxCalories)
}

func part2(scanner *bufio.Scanner) {
	maxCalories := []int{0, 0, 0}
	currentCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			maxCalories = append(maxCalories, currentCalories)
			sort.Ints(maxCalories)
			maxCalories = maxCalories[1:]
			currentCalories = 0
		} else {
			number, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += int(number)
		}
	}
	sum := 0
	for _, c := range maxCalories {
		sum += c
	}
	fmt.Println(sum)
}
