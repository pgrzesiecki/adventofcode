package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var word = []string{"M", "A", "S"}

func count_file_lines(file *os.File) (int, error) {
	defer file.Seek(0, 0)

	line_count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line_count++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error reading file: %s", err)
	}
	return line_count, nil
}

func prepare_matrix(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	file_lines, err := count_file_lines(file)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	matrix := make([][]string, file_lines)
	line_count := 0

	for scanner.Scan() {
		line := scanner.Text()
		matrix[line_count] = strings.Split(line, "")

		line_count++
	}

	return matrix, nil
}

func count_words_from(matrix [][]string, x, y int) int {

	found := 0

	y_len := len(matrix)
	x_len := len(matrix[y])

	// Diagonal (right, bottom)
	if x+1 < x_len && 0 <= x-1 && y+1 < y_len && 0 <= y-1 {
		if (matrix[y+1][x+1] == word[2] && matrix[y-1][x-1] == word[0]) || (matrix[y+1][x+1] == word[0] && matrix[y-1][x-1] == word[2]) {
			if (matrix[y+1][x-1] == word[2] && matrix[y-1][x+1] == word[0]) || (matrix[y+1][x-1] == word[0] && matrix[y-1][x+1] == word[2]) {
				found++
			}
		}
	}

	return found
}

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("Please provide a file name as an argument"))
	}

	matrix, err := prepare_matrix(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("Can not prepare matrix: %s", err))
	}

	found_words := 0
	for y, yv := range matrix {
		for x, xv := range yv {
			if xv == word[1] {
				found_words += count_words_from(matrix, x, y)
			}
		}
	}

	fmt.Println(found_words)
}
