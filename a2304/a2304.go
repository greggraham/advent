package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if len(os.Args) == 2 && os.Args[1] == "2" {
		part2(scanner)
	} else {
		part1(scanner)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}

func part1(scanner *bufio.Scanner) {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		winning, yours := getNumbers(line)
		sum += scorePoints(winning, yours)
	}
	fmt.Println(sum)
}

func getNumbers(line string) (winning map[int]struct{}, yours []int) {
	winning = make(map[int]struct{})
	endOfWinning := false
	numBarRe := regexp.MustCompile(`(Card +\d+:)|\d+|[|]`)
	res := numBarRe.FindAllString(line, -1)
	for _, token := range res {
		if token == "|" {
			endOfWinning = true
		} else if strings.HasPrefix(token, "Card") {
			continue
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal(err)
			}
			if endOfWinning {
				yours = append(yours, num)
			} else {
				winning[num] = struct{}{}
			}
		}
	}
	return winning, yours
}

func scorePoints(winning map[int]struct{}, yours []int) int {
	points := 0
	for _, num := range yours {
		_, ok := winning[num]
		if ok {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	fmt.Println(winning, yours, points)
	return points
}

func part2(scanner *bufio.Scanner) {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	fmt.Println(sum)
}
