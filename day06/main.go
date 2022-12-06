package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := readInput()
	part1(in)
}

func part1(s string) {
	for i := 4; i < len(s); i++ {
		if charsUnique(s[i-4 : i]) {
			fmt.Println("part 1: ", i)
			break
		}
	}
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
