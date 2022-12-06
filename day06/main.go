package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := readInput()
	part1(in)
	part2(in)
}

func part1(s string) {
	fmt.Println("part 1: ", findMarker(s, 4))
}

func part2(s string) {
	fmt.Println("part 2: ", findMarker(s, 14))
}

func findMarker(s string, numUnique int) int {
	for i := numUnique; i < len(s); i++ {
		if charsUnique(s[i-numUnique : i]) {
			return i
		}
	}
	return -1
}

func charsUnique(s string) bool {
	m := make(map[rune]bool)
	for _, c := range s {
		if m[c] {
			return false
		}
		m[c] = true
	}

	return true
}

func readInput() string {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
