package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var wordToDigit = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func findFirstAndLastDigit(line string) (int, int) {
	var firstPosition, lastPosition int
	firstPositionValue := -1
	lastPositionValue := -1

	for word, digit := range wordToDigit {
		re := regexp.MustCompile("(?i)" + word)

		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			if firstPositionValue < 0 || match[0] < firstPosition {
				firstPosition = match[0]
				firstPositionValue = digit
			}

			if lastPositionValue < 0 || match[0] > lastPosition {
				lastPosition = match[0]
				lastPositionValue = digit
			}
		}
	}

	return firstPositionValue, lastPositionValue
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name as an argument")
		return
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := findFirstAndLastDigit(line)

		digitStr := fmt.Sprintf("%d%d", firstDigit, lastDigit)
		digitNum, err := strconv.Atoi(digitStr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}

		sum += digitNum
	}

	fmt.Println(sum)
}
