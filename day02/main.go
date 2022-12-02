package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RockPaperScissors int
type WinLoseDraw int

const (
	ROCK = RockPaperScissors(iota)
	PAPER
	SCISSORS

	WIN = WinLoseDraw(iota)
	LOSE
	DRAW
)

type Round struct {
	Opponent RockPaperScissors
	You      RockPaperScissors
}

var (
	letterToRPS = map[string]RockPaperScissors{
		"A": ROCK, "X": ROCK,
		"B": PAPER, "Y": PAPER,
		"C": SCISSORS, "Z": SCISSORS,
	}

	rpsToPoints = map[RockPaperScissors]int{ROCK: 1, PAPER: 2, SCISSORS: 3}
	wldToPoints = map[WinLoseDraw]int{LOSE: 0, DRAW: 3, WIN: 6}

	toWinning = map[RockPaperScissors]RockPaperScissors{
		ROCK:     PAPER,
		PAPER:    SCISSORS,
		SCISSORS: ROCK,
	}
	toLosing = map[RockPaperScissors]RockPaperScissors{
		PAPER:    ROCK,
		SCISSORS: PAPER,
		ROCK:     SCISSORS,
	}
)

func main() {
	rounds := readInput()
	part1(rounds)
}

func part1(rounds []Round) {
	var points int
	for _, round := range rounds {
		points += calculatePoints(round)
	}

	fmt.Println("part 1: ", points)
}

func calculatePoints(round Round) int {
	outcome := getOutcome(round)
	return rpsToPoints[round.You] + wldToPoints[outcome]
}

func getOutcome(round Round) WinLoseDraw {
	if round.You == round.Opponent {
		return DRAW
	}
	if winningMove := toWinning[round.Opponent]; winningMove == round.You {
		return WIN
	}
	return LOSE
}

func readInput() []Round {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rounds []Round
	for scanner.Scan() {
		choices := strings.Split(scanner.Text(), " ")
		rounds = append(rounds, Round{
			Opponent: letterToRPS[choices[0]],
			You:      letterToRPS[choices[1]],
		})
	}

	return rounds
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
