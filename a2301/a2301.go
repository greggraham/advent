package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	part1(scanner)
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
		} else {
			sum += value
		}
	}
	fmt.Println(sum)
}

func part2(scanner *bufio.Scanner) {

}
