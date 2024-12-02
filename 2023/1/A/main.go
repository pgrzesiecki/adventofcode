package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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

		var firstDigit, lastDigit rune
		foundFirst := false

		for _, char := range line {
			if unicode.IsDigit(char) {
				if !foundFirst {
					firstDigit = char
					foundFirst = true
				}
				lastDigit = char
			}
		}

		digitStr := fmt.Sprintf("%c%c", firstDigit, lastDigit)
		digitNum, err := strconv.Atoi(digitStr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}

		sum += digitNum
	}

	fmt.Println(sum)
}
