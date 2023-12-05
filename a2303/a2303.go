package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	// Locate symbols
	var symbols [][]int
	var lines []string
	symRe := regexp.MustCompile(`[^.0-9]`)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		symIndices := symRe.FindAllStringIndex(line, -1)
		var thisLineSym []int
		for _, elem := range symIndices {
			thisLineSym = append(thisLineSym, elem[0])
		}
		symbols = append(symbols, thisLineSym)
	}

	// Find part numbers
	numRe := regexp.MustCompile(`\d+`)
	sum := 0
	for lineNum, line := range lines {
		numbers := numRe.FindAllString(line, -1)
		indices := numRe.FindAllStringIndex(line, -1)

		// Loop through the indices of each number in the current line
		for i, pos := range indices {
			skip := false
			// Check for symbols on the line above
			if lineNum > 0 {
				for _, si := range symbols[lineNum-1] {
					if si >= pos[0]-1 && si <= pos[1] {
						num, err := strconv.Atoi(numbers[i])
						if err != nil {
							log.Fatal(err)
						}
						sum += num
						skip = true
						break
					}
				}
			}
			if skip {
				continue
			}

			// Check for symbols on the current line
			for _, si := range symbols[lineNum] {
				if si >= pos[0]-1 && si <= pos[1] {
					num, err := strconv.Atoi(numbers[i])
					if err != nil {
						log.Fatal(err)
					}
					sum += num
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			// Check for symbols on the next line
			if lineNum+1 < len(symbols) {
				for _, si := range symbols[lineNum+1] {
					if si >= pos[0]-1 && si <= pos[1] {
						num, err := strconv.Atoi(numbers[i])
						if err != nil {
							log.Fatal(err)
						}
						sum += num
						skip = true
						break
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func part2(scanner *bufio.Scanner) {
	sum := 0

	// Locate numLocs
	var numLocs [][][]int
	var lines []string
	numRe := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		numLocs = append(numLocs, numRe.FindAllStringIndex(line, -1))
	}

	// Find gears
	starRe := regexp.MustCompile(`[*]`)
	for lineNum := 1; lineNum < len(lines)-1; lineNum++ {
		line := lines[lineNum]
		starLocs := starRe.FindAllStringIndex(line, -1)
		fmt.Println(starLocs)
		for _, starLoc := range starLocs {
			sl := starLoc[0]
			var parts []int
			for ln := lineNum - 1; ln <= lineNum+1; ln++ {
				for _, numLoc := range numLocs[ln] {
					if sl >= numLoc[0]-1 && sl <= numLoc[1] {
						partNum, err := strconv.Atoi(lines[ln][numLoc[0]:numLoc[1]])
						if err != nil {
							log.Fatal(err)
						}
						parts = append(parts, partNum)
						fmt.Printf("Found part %d\n", partNum)
					}
				}
			}
			if len(parts) == 2 {
				ratio := parts[0] * parts[1]
				sum += ratio
				fmt.Println("Gear: ", parts, ratio, sum)
			}
		}
	}
	fmt.Println(sum)
}
