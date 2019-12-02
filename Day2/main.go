package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
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
	code := []int{}
	for _, j := range codes {
		i, err := strconv.Atoi(j)
		if err != nil {
			fmt.Printf("Error converting string to int %v\n", j)
		}
		code = append(code, i)
	}

	var opCode int
	var op1 int
	var op2 int
	var pntRes int

	fmt.Println(code)

	opCode = 0
	for {

		op1 = opCode + 1
		op2 = opCode + 2
		pntRes = opCode + 3
		if code[opCode] != 99 {
			fmt.Printf("%d, %d, %d, %d\n", code[opCode], code[op1], code[op2], code[pntRes])
		}

		f := operation(code[opCode])
		if code[opCode] == 99 {
			break
		}
		res, err := f(code[code[op1]], code[code[op2]])
		fmt.Printf("res: %d\n", res)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("save %d in %d\n", res, pntRes)
		code[code[pntRes]] = res
		opCode = opCode + 4
		fmt.Println(code)
	}

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
