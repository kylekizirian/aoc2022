package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Move struct {
	Num, From, To int
}

func main() {
	crates, moves := readInput()
	part1(deepCopyCrates(crates), moves)
	part2(deepCopyCrates(crates), moves)
}

func part1(crates map[int][]string, moves []Move) {
	// 0th index is the top, so we remove from the front and pre-prend
	for _, move := range moves {
		for i := 0; i < move.Num; i++ {
			crates[move.To] = append([]string{crates[move.From][0]}, crates[move.To]...)
			crates[move.From] = crates[move.From][1:]
		}
	}

	var top string
	for i := 1; i < 10; i++ {
		if len(crates[i]) > 0 {
			top += crates[i][0]
		}
	}

	fmt.Println("part 1: ", top)
}

func part2(crates map[int][]string, moves []Move) {
	for _, move := range moves {
		temp := deepCopyStrings(crates[move.From][:move.Num])
		crates[move.To] = append(temp[:move.Num], crates[move.To]...)
		crates[move.From] = crates[move.From][move.Num:]
	}

	var top string
	for i := 1; i < 10; i++ {
		if len(crates[i]) > 0 {
			top += crates[i][0]
		}
	}

	fmt.Println("part 2: ", top)
}

func deepCopyStrings(in []string) []string {
	deepCopy := make([]string, len(in))
	copy(deepCopy, in)
	return deepCopy
}

func deepCopyCrates(in map[int][]string) map[int][]string {
	deepCopy := make(map[int][]string)
	for i, s := range in {
		deepCopy[i] = make([]string, len(s))
		for j, c := range s {
			deepCopy[i][j] = c
		}
	}
	return deepCopy
}

func readInput() (map[int][]string, []Move) {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	crateRe := regexp.MustCompile(`\[[A-Z]\]`)
	moveRe := regexp.MustCompile(`move \d+ from \d+ to \d+`)
	numRe := regexp.MustCompile(`\d+`)

	var moves []Move
	locToCrates := make(map[int][]string)
	for i := 1; i < 10; i++ {
		locToCrates[i] = make([]string, 0)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if crateRe.MatchString(line) {
			//figure out crate location and do append
			indices := crateRe.FindAllStringIndex(line, -1)
			for _, index := range indices {
				crateLoc := (index[0] / 4) + 1
				locToCrates[crateLoc] = append(locToCrates[crateLoc], string(line[index[0]+1]))
			}
		}

		if moveRe.MatchString(line) {
			// extract all digits
			digits := numRe.FindAllString(line, -1)
			moves = append(moves, Move{
				Num:  mustAtoi(digits[0]),
				From: mustAtoi(digits[1]),
				To:   mustAtoi(digits[2]),
			})
		}
	}

	return locToCrates, moves
}

func mustAtoi(s string) int {
	num, err := strconv.Atoi(s)
	checkErr(err)
	return num
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
