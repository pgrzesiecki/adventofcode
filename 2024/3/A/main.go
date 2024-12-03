package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	result := 0
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}

			b, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}

			result += a * b
		}
	}

	fmt.Println(result)
}
