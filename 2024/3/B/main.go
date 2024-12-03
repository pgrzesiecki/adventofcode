package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
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

	is_in_use := true
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		use_pos := make([]bool, len(line))

		re_do := regexp.MustCompile(`do\(\)`)
		re_dont := regexp.MustCompile(`don\'t\(\)`)

		matches_do := re_do.FindAllStringSubmatchIndex(line, -1)
		matches_dont := re_dont.FindAllStringSubmatchIndex(line, -1)

		do_start_pos := []int{}
		dont_start_pos := []int{}

		for _, match := range matches_do {
			do_start_pos = append(do_start_pos, match[0])
		}

		for _, match := range matches_dont {
			dont_start_pos = append(dont_start_pos, match[0])
		}

		for i := range line {
			if slices.Contains(do_start_pos, i) {
				is_in_use = true
			}

			if slices.Contains(dont_start_pos, i) {
				is_in_use = false
			}

			use_pos[i] = is_in_use
		}

		re_mul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := re_mul.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			if use_pos[match[0]] {

				a, err := strconv.Atoi(string(line[match[2]:match[3]]))
				if err != nil {
					panic(err)
				}

				b, err := strconv.Atoi(string(line[match[4]:match[5]]))
				if err != nil {
					panic(err)
				}

				result += a * b
			}
		}
	}

	fmt.Println(result)
}
