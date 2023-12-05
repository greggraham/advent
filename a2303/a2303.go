package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
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
	var symbols [][]int
	lineNum := 0
	symRe := regexp.MustCompile(`[^.0-9]`)
	for scanner.Scan() {
		line := scanner.Text()
		symbols = append(symbols, symRe.FindAllStringIndex(line, -1)[0])
	}
	for scanner.Scan() {
		line := scanner.Text()

	}
}

func part2(scanner *bufio.Scanner) {
}
