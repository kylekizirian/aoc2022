package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	calories := readInput()
	part1(calories)
	part2(calories)
}

func part1(caloriesPerElf [][]int) {
	var max int
	for _, elf := range caloriesPerElf {
		calories := sum(elf)
		if calories > max {
			max = calories
		}
	}
	fmt.Println("part 1: ", max)
}

func part2(caloriesPerElf [][]int) {
	sums := make([]int, len(caloriesPerElf))
	for i, elf := range caloriesPerElf {
		sums[i] = sum(elf)
	}
	numElves := len(caloriesPerElf)
	sort.Ints(sums)
	fmt.Println("part 2: ", sums[numElves-3]+sums[numElves-2]+sums[numElves-1])
}

func sum(arr []int) int {
	var s int
	for _, n := range arr {
		s += n
	}
	return s
}

func readInput() [][]int {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var caloriesPerElf [][]int
	var curCalories []int

	for scanner.Scan() {
		if scanner.Text() == "" {
			if len(curCalories) > 0 {
				caloriesPerElf = append(caloriesPerElf, curCalories)
			}
			curCalories = make([]int, 0)
			continue
		}

		cal, err := strconv.Atoi(scanner.Text())
		checkErr(err)
		curCalories = append(curCalories, cal)
	}

	return caloriesPerElf
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
