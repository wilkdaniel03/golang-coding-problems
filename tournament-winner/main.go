package main

import "fmt"

func solution(competitions [][]string, results []int) string {
	winner := ""
	scores := map[string]int{winner: 0}
	for i, comp := range competitions {
		currentResult := results[i]
		homeTeam, awayTeam := comp[0], comp[1]
		currentWinner := homeTeam
		if currentResult == 0 {
			currentWinner = awayTeam
		}
		increaseTeamPoints(currentWinner, 3, scores)
		if compareTeams(currentWinner, winner, scores) {
			winner = currentWinner
		}
	}
	return winner
}

func increaseTeamPoints(team string, points int, scores map[string]int) {
	if !checkIfIsInMap(team, scores) {
		scores[team] = 0
	}
	scores[team] += 3
}

func compareTeams(team string, comparedTeam string, scores map[string]int) bool {
	if scores[team] < scores[comparedTeam] {
		return false
	}
	return true
}

func checkIfIsInMap(team string, scores map[string]int) bool {
	if _, found := scores[team]; !found {
		return false
	}
	return true
}

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	fmt.Printf("Solution => %s\n", solution(competitions, results))
}
