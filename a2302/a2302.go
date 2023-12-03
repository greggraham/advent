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
	gameRe := regexp.MustCompile(`Game (\d+):`)
	redRe := regexp.MustCompile(`(\d+) red`)
	greenRe := regexp.MustCompile(`(\d+) green`)
	blueRe := regexp.MustCompile(`(\d+) blue`)
	sum := 0
	for scanner.Scan() {
		skip := false
		line := scanner.Text()

		// Get game number
		res := gameRe.FindStringSubmatch(line)
		gameNum, err := strconv.Atoi(res[1])
		if err != nil {
			log.Fatal(err)
		}

		// Check reds
		mres := redRe.FindAllStringSubmatch(line, -1)
		for _, r := range mres {
			redCount, err := strconv.Atoi(r[1])
			if err != nil {
				log.Fatal(err)
			}
			if redCount > 12 {
				skip = true
				fmt.Println("Red Skipped", line)
				break
			}
		}
		if skip {
			continue
		}

		// Check greens
		mres = greenRe.FindAllStringSubmatch(line, -1)
		for _, r := range mres {
			greenCount, err := strconv.Atoi(r[1])
			if err != nil {
				log.Fatal(err)
			}
			if greenCount > 13 {
				skip = true
				fmt.Println("Green Skipped", line)
				break
			}
		}
		if skip {
			continue
		}

		// Check blues
		mres = blueRe.FindAllStringSubmatch(line, -1)
		for _, r := range mres {
			blueCount, err := strconv.Atoi(r[1])
			if err != nil {
				log.Fatal(err)
			}
			if blueCount > 14 {
				skip = true
				fmt.Println("Blue Skipped", line)
				break
			}
		}
		if skip {
			continue
		}

		sum += gameNum
	}
	fmt.Println(sum)
}

func part2(scanner *bufio.Scanner) {
	redRe := regexp.MustCompile(`(\d+) red`)
	greenRe := regexp.MustCompile(`(\d+) green`)
	blueRe := regexp.MustCompile(`(\d+) blue`)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		minRed := 0
		minGreen := 0
		minBlue := 0

		// Check reds
		mres := redRe.FindAllStringSubmatch(line, -1)
		for _, r := range mres {
			redCount, err := strconv.Atoi(r[1])
			if err != nil {
				log.Fatal(err)
			}
			if redCount > minRed {
				minRed = redCount
			}
		}

		// Check greens
		mres = greenRe.FindAllStringSubmatch(line, -1)
		for _, r := range mres {
			greenCount, err := strconv.Atoi(r[1])
			if err != nil {
				log.Fatal(err)
			}
			if greenCount > minGreen {
				minGreen = greenCount
			}
		}

		// Check blues
		mres = blueRe.FindAllStringSubmatch(line, -1)
		for _, r := range mres {
			blueCount, err := strconv.Atoi(r[1])
			if err != nil {
				log.Fatal(err)
			}
			if blueCount > minBlue {
				minBlue = blueCount
			}
		}
		power := minRed * minGreen * minBlue
		fmt.Println(line, power, minRed, minGreen, minBlue)
		sum += power
	}
	fmt.Println(sum)
}
