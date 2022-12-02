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

var (
	letterToRPS = map[string]RockPaperScissors{
		"A": ROCK, "X": ROCK,
		"B": PAPER, "Y": PAPER,
		"C": SCISSORS, "Z": SCISSORS,
	}

	letterToOutcome = map[string]WinLoseDraw{"X": LOSE, "Y": DRAW, "Z": WIN}

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

type Line struct {
	First  string
	Second string
}

type Round struct {
	Opponent RockPaperScissors
	You      RockPaperScissors
}

func main() {
	lines := readInput()
	part1(lines)
	part2(lines)
}

func part1(lines []Line) {
	rounds := make([]Round, len(lines))
	for i, line := range lines {
		rounds[i] = line.toRoundPart1()
	}

	var points int
	for _, round := range rounds {
		points += calculatePoints(round)
	}

	fmt.Println("part 1: ", points)
}

func part2(lines []Line) {
	rounds := make([]Round, len(lines))
	for i, line := range lines {
		rounds[i] = line.toRoundPart2()
	}

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

func (l Line) toRoundPart1() Round {
	return Round{
		Opponent: letterToRPS[l.First],
		You:      letterToRPS[l.Second],
	}
}

func (l Line) toRoundPart2() Round {
	var you RockPaperScissors
	opponent := letterToRPS[l.First]
	outcome := letterToOutcome[l.Second]

	switch outcome {
	case DRAW:
		you = opponent
	case LOSE:
		you = toLosing[opponent]
	case WIN:
		you = toWinning[opponent]
	}

	return Round{
		Opponent: opponent,
		You:      you,
	}
}

func readInput() []Line {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []Line
	for scanner.Scan() {
		choices := strings.Split(scanner.Text(), " ")
		lines = append(lines, Line{
			First:  choices[0],
			Second: choices[1],
		})
	}

	return lines
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
