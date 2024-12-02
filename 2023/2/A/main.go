package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var games = map[int][]map[string]int{}

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

	for scanner.Scan() {
		line := scanner.Text()
		parseLine(line)
	}

	maxCubes := map[string]int{"red": 12, "green": 13, "blue": 14}
	possibleGames := findPossibleGamesForConfig(maxCubes)

	fmt.Println(sumGames(possibleGames))
}

func sumGames(games []int) int {
	sum := 0
	for _, game := range games {
		sum += game
	}
	return sum
}

func findPossibleGamesForConfig(config map[string]int) []int {
	var possibleGames []int
	for gameNumber, teams := range games {
		if !isImpossibleGame(teams, config) {
			possibleGames = append(possibleGames, gameNumber)
		}
	}
	return possibleGames
}

func isImpossibleGame(teams []map[string]int, config map[string]int) bool {
	for _, team := range teams {
		for color, count := range team {
			if count > config[color] {
				return true
			}
		}
	}
	return false
}

func parseLine(line string) {
	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		fmt.Println("Invalid line format:", line)
		return
	}

	gameNumber, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(parts[0], "Game")))
	if err != nil {
		fmt.Println("Invalid game number:", parts[0])
		return
	}

	teams := strings.Split(parts[1], ";")
	var teamMaps []map[string]int

	for _, team := range teams {
		teamMap := make(map[string]int)
		players := strings.Split(team, ",")
		for _, player := range players {
			playerParts := strings.Fields(strings.TrimSpace(player))
			if len(playerParts) != 2 {
				fmt.Println("Invalid player format:", player)
				return
			}
			count, err := strconv.Atoi(playerParts[0])
			if err != nil {
				fmt.Println("Invalid player count:", playerParts[0])
				return
			}
			color := playerParts[1]
			teamMap[color] = count
		}
		teamMaps = append(teamMaps, teamMap)
	}

	games[gameNumber] = teamMaps
}
