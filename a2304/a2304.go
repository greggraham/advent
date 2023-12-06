package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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
	numBarRe := regexp.MustCompile(`(Card +\d+:)|\d+|[|]`)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		winning, yours := getNumbers(line, numBarRe)
		sum += scorePoints(winning, yours)
	}
	fmt.Println(sum)
}

func getNumbers(line string, numBarRe *regexp.Regexp) (winning map[int]struct{}, yours []int) {
	winning = make(map[int]struct{})
	endOfWinning := false
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
	return points
}

func countWins(winning map[int]struct{}, yours []int) int {
	wins := 0
	for _, num := range yours {
		_, ok := winning[num]
		if ok {
			wins++
		}
	}
	return wins
}

type CardRec struct {
	CardNum int
	Wins    int
	Copies  int
}

func part2(scanner *bufio.Scanner) {
	numBarRe := regexp.MustCompile(`(Card +\d+:)|\d+|[|]`)
	var wins []int
	var cards []int
	cardNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		winning, yours := getNumbers(line, numBarRe)
		wins = append(wins, countWins(winning, yours))
		cards = append(cards, cardNum)
		cardNum++
	}
	maxCardNum := cardNum - 1

	for i := 0; i < len(cards); i++ {
		cardNum = cards[i]
		w := wins[cardNum-1]
		fmt.Printf("Card %d Wins %d Length %d\n", cardNum, w, len(cards))
		for j := 1; j <= w && cardNum+j <= maxCardNum; j++ {
			cards = append(cards, cardNum+j)
		}
		sort.Ints(cards)
	}

	fmt.Println(len(cards))
}
