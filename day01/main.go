package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func parseArgs(args []string) (string, error) {
	if len(args) < 2 {
		fmt.Printf("No input file provided")
		fmt.Printf("Usage: go run %s `inputfile`", args[1])
		return "", errors.New("No input file provided")
	}
	if len(args) != 2 {
		fmt.Printf("Illegal number of arguments provided")
		fmt.Printf("Usage: go run %s `inputfile`", args[1])
		return "", errors.New("Illegal number of arguments provided")

	}
	return args[1], nil
}

func stripLine(line string) []string {
	var output []string
	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		var word bytes.Buffer
		for !unicode.IsSpace(char) {
			word.WriteString(string(char))
			i++
			if i >= len(line) {
				break
			}
			char = rune(line[i])
		}
		if word.Len() > 0 {
			output = append(output, word.String())
		}
	}
	return output
}

func assert(condition bool, message string) {
	if !condition {
		panic(errors.New(message))
	}
}

func main() {
	args := os.Args
	filename, err := parseArgs(args)
	checkError(err)
	data, err := os.ReadFile(filename)
	checkError(err)
	input := string(data)
	fmt.Println(q1(input))
	fmt.Println(q2(input))
}

func absoluteVal(num1 int, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}

func getLists(input string) ([]int, []int) {
	var lh []int
	var rh []int
	for _, line := range strings.Split(input, "\r\n") {
		sline := stripLine(line)
		assert(len(sline) == 2, fmt.Sprintf("Expect 2 elements in each line, but got %d", len(sline)))
		lhnumber, err := strconv.Atoi(sline[0])
		checkError(err)
		rhnumber, err := strconv.Atoi(sline[1])
		checkError(err)
		lh = append(lh, lhnumber)
		rh = append(rh, rhnumber)
	}

	return lh, rh
}

func q1(input string) int {
	lh, rh := getLists(input)
	assert(len(lh) == len(rh), fmt.Sprintf("Expected the 2 lists to be of equal length, but they are different.\nlh = %d, rh = %d", len(lh), len(rh)))
	slices.Sort(lh[:])
	slices.Sort(rh[:])
	distance := 0
	for i := 0; i < len(lh); i++ {
		distance += absoluteVal(lh[i], rh[i])
	}
	return distance
}

func q2(input string) int {
	lh, rh := getLists(input)
	assert(len(lh) == len(rh), fmt.Sprintf("Expected the 2 lists to be of equal length, but they are different.\nlh = %d, rh = %d", len(lh), len(rh)))
	slices.Sort(lh[:])
	slices.Sort(rh[:])
	var m map[int]int
	m = make(map[int]int)
	similarity := 0
	for i := 0; i < len(rh); i++ {
		_, ok := m[rh[i]]
		if !ok {
			m[rh[i]] = 0
		}
		m[rh[i]] += 1
	}
	for i := 0; i < len(lh); i++ {
		val, ok := m[lh[i]]
		if !ok {
			val = 0
		}
		similarity += lh[i] * val
	}
	return similarity
}
