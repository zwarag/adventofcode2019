package main

import (
	"fmt"
	"strconv"
)

const input = "245318-765747"

func main() {
	part1()
}

func part1() {
	solutions := 0
	min := input[:6]
	max := input[7:]
	fmt.Println(min, max)
	//solutions := []int{}
	num, _ := strconv.Atoi(min)
	numMax, _ := strconv.Atoi(max)

	for num < numMax {
		candidate := strconv.Itoa(num)
		b := true
		double := false
		occ := 0
		lastChar := rune(' ')

		for pos, char := range candidate {
			if pos == 0 {
				occ++
				lastChar = char
			} else {
				if lastChar == char {
					occ++
					double = true
				} else {
					occ = 0
				}
				if lastChar > char {
					b = false
					break
				}
				lastChar = char
			}
		}
		if b && double {
			solutions++
		}

		num++
	}
	fmt.Println(solutions)
}
