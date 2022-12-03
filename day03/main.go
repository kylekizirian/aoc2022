package main

import (
	"bufio"
	"fmt"
	"os"
)

type Rucksack struct {
	Compartment1 map[rune]bool
	Compartment2 map[rune]bool
	Sack         map[rune]bool
}

func main() {
	rucksacks := readInput()
	part1(rucksacks)
	part2(rucksacks)
}

func part1(rucksacks []Rucksack) {
	var total int
	for _, rucksack := range rucksacks {
		for r, _ := range rucksack.Compartment1 {
			if rucksack.Compartment2[r] {
				total += priority(r)
			}
		}
	}
	fmt.Println("part 1: ", total)
}

func part2(rucksacks []Rucksack) {
	var total int
	for i := 0; i < len(rucksacks); i += 3 {
		r1, r2, r3 := rucksacks[i], rucksacks[i+1], rucksacks[i+2]
		for r, _ := range r1.Sack {
			if r2.Sack[r] && r3.Sack[r] {
				total += priority(r)
				break
			}
		}
	}
	fmt.Println("part 2: ", total)
}

func readInput() []Rucksack {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rucksacks []Rucksack

	for scanner.Scan() {
		line := []rune(scanner.Text())
		rucksack := Rucksack{
			Compartment1: make(map[rune]bool),
			Compartment2: make(map[rune]bool),
			Sack:         make(map[rune]bool),
		}
		for _, r := range line[len(line)/2:] {
			rucksack.Compartment1[r] = true
			rucksack.Sack[r] = true
		}
		for _, r := range line[:len(line)/2] {
			rucksack.Compartment2[r] = true
			rucksack.Sack[r] = true
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
