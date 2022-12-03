package main

import (
	"bufio"
	"fmt"
	"os"
)

type Rucksack struct {
	Compartment1Slice []rune
	Compartment1Map   map[rune]bool
	Compartment2Slice []rune
	Compartment2Map   map[rune]bool
}

func main() {
	rucksacks := readInput()
	part1(rucksacks)
}

func part1(rucksacks []Rucksack) {
	var total int
	for _, rucksack := range rucksacks {
		for r, _ := range rucksack.Compartment1Map {
			if rucksack.Compartment2Map[r] {
				total += priority(r)
			}
		}
	}
	fmt.Println("part 1: ", total)
}

func readInput() []Rucksack {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rucksacks []Rucksack

	for scanner.Scan() {
		lineRunes := []rune(scanner.Text())
		rucksack := Rucksack{
			Compartment1Slice: lineRunes[:len(lineRunes)/2],
			Compartment1Map:   make(map[rune]bool),
			Compartment2Slice: lineRunes[len(lineRunes)/2:],
			Compartment2Map:   make(map[rune]bool),
		}
		for _, r := range rucksack.Compartment1Slice {
			rucksack.Compartment1Map[r] = true
		}
		for _, r := range rucksack.Compartment2Slice {
			rucksack.Compartment2Map[r] = true
		}
		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return 1 + int(r-'a')
	}
	return 27 + int(r-'A')
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
