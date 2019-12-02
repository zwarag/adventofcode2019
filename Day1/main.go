package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var fuel int

	file, err := os.Open("D:/Project/Projects/AdventOfCode2019/Day1/input")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fuel += calculateFuel(i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fuel)
}

func calculateFuel(i int) int {
	if i <= 0 {
		return 0
	} else {
		ii := ((i / 3) - 2)
		if ii <= 0 {
			ii = 0
		}
		return ii + calculateFuel(ii)
	}
}
