package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		digits := []rune{' ', ' '}
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if digits[0] == ' ' {
					digits[0] = c
				}
				digits[1] = c
			}
		}
		dstr := string(digits)
		fmt.Println(line, dstr)
		value, err := strconv.Atoi(dstr)
		if err != nil {
			log.Fatal("ParseInt failed on", dstr)
		}
		sum += value
	}
	fmt.Println(sum)
}

func part2(scanner *bufio.Scanner) {
	sum := 0
	lines := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		line = strings.ReplaceAll(line, "one", "o1e")
		line = strings.ReplaceAll(line, "two", "t2o")
		line = strings.ReplaceAll(line, "three", "t3e")
		line = strings.ReplaceAll(line, "four", "f4r")
		line = strings.ReplaceAll(line, "five", "f5e")
		line = strings.ReplaceAll(line, "six", "s6x")
		line = strings.ReplaceAll(line, "seven", "s7n")
		line = strings.ReplaceAll(line, "eight", "e8t")
		line = strings.ReplaceAll(line, "nine", "n9e")

		digits := []rune{' ', ' '}
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if digits[0] == ' ' {
					digits[0] = c
				}
				digits[1] = c
			}
		}
		dstr := string(digits)
		fmt.Println(line, dstr)
		value, err := strconv.Atoi(dstr)
		if err != nil {
			log.Fatal("ParseInt failed on", dstr)
		}
		sum += value
		fmt.Println(value, sum)
		lines++
	}
	fmt.Println(sum)
	fmt.Println(lines)
}
