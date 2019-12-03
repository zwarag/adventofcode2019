package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var err error
	file, err := os.Open("D:/Project/Projects/AdventOfCode2019/Day2/input")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	codes := strings.Split(line, ",")

	noun := 0
	verb := 0

	for verb != 100 {
		res, err := compute(noun, verb, codes)
		if err != nil {
			fmt.Println(err)
			return
		}
		if res == 19690720 {
			fmt.Println("found it")
			fmt.Printf("noun: %d, verb: %d\n", noun, verb)
			fmt.Printf("What is 100 * noun + verb = %d", 100*noun+verb)
			return
		}
		noun++
		if math.Mod(float64(noun), 100) == 0 {
			noun = 0
			verb++
		}
	}
	fmt.Println("no solution found")
}

func compute(i1, i2 int, codes []string) (int, error) {
	code := []int{}
	for _, j := range codes {
		i, err := strconv.Atoi(j)
		if err != nil {
			s := fmt.Sprintf("Error converting string to int %v\n", j)
			return 0, errors.New(s)
		}
		code = append(code, i)
	}

	code[1] = i1
	code[2] = i2

	var opCode int
	var op1 int
	var op2 int
	var pntRes int

	opCode = 0
	for {

		op1 = opCode + 1
		op2 = opCode + 2
		pntRes = opCode + 3

		f := operation(code[opCode])
		if code[opCode] == 99 {
			break
		}
		res, err := f(code[code[op1]], code[code[op2]])
		if err != nil {
			s := fmt.Sprintln(err)
			return 0, errors.New(s)
		}
		code[code[pntRes]] = res
		opCode = opCode + 4
	}
	return code[0], nil
}

func operation(opCode int) func(num1, num2 int) (int, error) {
	switch opCode {
	case 1:
		return func(num1, num2 int) (int, error) {
			return num1 + num2, nil
		}
	case 2:
		return func(num1, num2 int) (int, error) {
			return num1 * num2, nil
		}
	case 99:
		return func(num1, num2 int) (int, error) {
			return -1, nil
		}
	default:
		break
	}
	err := fmt.Sprintf("Error, unknown OPCODE %d\n", opCode)
	return func(num1, num2 int) (int, error) { return -2, errors.New(err) }
}
