package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ElfAssignment struct{ Start, End int }

func main() {
	elves := readInput()
	part1(elves)
	part2(elves)
}

func part1(elves [][2]ElfAssignment) {
	contains := func(first, second ElfAssignment) bool {
		return (first.Start <= second.Start && first.End >= second.End) ||
			(first.Start >= second.Start && first.End <= second.End)
	}

	var numContains int
	for _, pair := range elves {
		if contains(pair[0], pair[1]) {
			numContains++
		}
	}

	fmt.Println("part 1: ", numContains)
}

func part2(elves [][2]ElfAssignment) {
	overlaps := func(first, second ElfAssignment) bool {
		var min, max ElfAssignment
		if first.Start < second.Start {
			min, max = first, second
		} else {
			min, max = second, first
		}

		return max.Start <= min.End
	}

	var numOverlaps int
	for _, pair := range elves {
		if overlaps(pair[0], pair[1]) {
			numOverlaps++
		}
	}

	fmt.Println("part 1: ", numOverlaps)
}

func readInput() [][2]ElfAssignment {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var elves [][2]ElfAssignment

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		e1, e2 := strings.Split(line[0], "-"), strings.Split(line[1], "-")
		elves = append(elves, [2]ElfAssignment{
			{Start: mustAtoi(e1[0]), End: mustAtoi(e1[1])},
			{Start: mustAtoi(e2[0]), End: mustAtoi(e2[1])},
		})
	}

	return elves
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	checkErr(err)
	return i
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
