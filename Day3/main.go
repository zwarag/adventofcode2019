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

	var w1, w2 string
	var wire1, wire2 []string
	var err error

	file, err := os.Open("D:/Project/Projects/AdventOfCode2019/Day3/input")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		w1 = scanner.Text()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
	}
	if scanner.Scan() {
		w2 = scanner.Text()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wire1 = strings.Split(w1, ",")
	wire2 = strings.Split(w2, ",")
	fmt.Println(wire1)
	fmt.Println(wire2)
	fmt.Println()

	U1, R1, D1, L1, err := maxWidthLength(wire1)
	if err != nil {
		fmt.Println(err)
		return
	}
	U2, R2, D2, L2, err := maxWidthLength(wire2)
	if err != nil {
		fmt.Println(err)
		return
	}

	U, R, D, L := overAllMaxWidthAndHeight(U1, U2, R1, R2, D1, D2, L1, L2)

	var length, height, width int
	height = U + D
	width = R + L
	if math.Mod(float64(height), 2) == 0 {
		height++
	}
	if math.Mod(float64(width), 2) == 0 {
		width++
	}
	if height > width {
		length = height
	} else {
		length = width
	}

	fmt.Printf("length: %d\n", length)

	pf := initPF(length)

	for y, i := range pf {
		for x, _ := range i {
			pf[y][x] = "."
		}
	}

	pf = draw(pf, wire1)
	printPF(pf)
	pf = draw(pf, wire2)

	for y, i := range pf {
		for x, _ := range i {
			if (length/2) == y && (length/2) == x {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func draw(pf [][]string, wire []string) [][]string {
	var x, y int
	x = len(pf) / 2
	y = len(pf) / 2
	pf[y][x] = "o"
	fmt.Println(wire)
	for _, e := range wire {
		if strings.Contains(e, "U") {
			times := getNum(e)
			for i := 0; i < times; i++ {
				y++
				if pf[y][x] == "X" {
				} else if pf[y][x] == "*" {
					pf[y][x] = "X"
				} else {
					pf[y][x] = "*"
				}
			}
		} else if strings.Contains(e, "R") {
			times := getNum(e)
			for i := 0; i < times; i++ {
				x++
				if pf[y][x] == "X" {
				} else if pf[y][x] == "*" {
					pf[y][x] = "X"
				} else {
					pf[y][x] = "*"
				}
			}
		} else if strings.Contains(e, "D") {
			times := getNum(e)
			for i := 0; i < times; i++ {
				y--
				if pf[y][x] == "X" {
				} else if pf[y][x] == "*" {
					pf[y][x] = "X"
				} else {
					pf[y][x] = "*"
				}
			}
		} else if strings.Contains(e, "L") {
			times := getNum(e)
			for i := 0; i < times; i++ {
				x--
				if pf[y][x] == "X" {
				} else if pf[y][x] == "*" {
					pf[y][x] = "X"
				} else {
					pf[y][x] = "*"
				}
			}
		}
	}

	return pf
}

func initPF(l int) (pf [][]string) {
	pf = make([][]string, l)
	for i := range pf {
		pf[i] = make([]string, l)
	}
	return pf
}

func printPF(pf [][]string) {
	for y, i := range pf {
		for x, _ := range i {
			fmt.Print(pf[y][x])
		}
		fmt.Println()
	}
}

func maxWidthLength(wire []string) (U, R, D, L int, err error) {
	for _, s := range wire {
		fmt.Println(s)
		if strings.Contains(s, "U") {
			U += getNum(s)
		} else if strings.Contains(s, "R") {
			R += getNum(s)
		} else if strings.Contains(s, "D") {
			D += getNum(s)
		} else if strings.Contains(s, "L") {
			L += getNum(s)
		} else {
			return 0, 0, 0, 0, errors.New("Error while figuring out out width and height")
		}
	}
	return U, R, D, L, nil
}

func overAllMaxWidthAndHeight(U1, U2, R1, R2, D1, D2, L1, L2 int) (U, R, D, L int) {
	if U1 > U2 {
		U = U1
	} else {
		U = U2
	}

	if R1 > R2 {
		R = R1
	} else {
		R = R2
	}

	if D1 > D2 {
		D = D1
	} else {
		D = D2
	}

	if L1 > L2 {
		L = L1
	} else {
		L = L2
	}

	return U, R, D, L
}

func getNum(s string) int {
	var retval int
	s2 := s[1:]
	retval, _ = strconv.Atoi(s2) // no error check :O
	return retval
}
