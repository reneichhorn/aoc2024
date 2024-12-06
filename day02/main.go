package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func parseArgs(args []string) (string, error) {
	if len(args) < 2 {
		fmt.Println("No input file provided")
		fmt.Printf("Usage: go run %s `inputfile`\n", args[0])
		return "", errors.New("no input file provided")
	}
	if len(args) != 2 {
		fmt.Println("Illegal number of arguments provided")
		fmt.Printf("Usage: go run %s `inputfile`\n", args[1])
		return "", errors.New("illegal number of arguments provided")

	}
	return args[1], nil
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func stripLine(line string) []int {
	var output []int
	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		var word strings.Builder
		for !unicode.IsSpace(char) {
			word.WriteString(string(char))
			i++
			if i >= len(line) {
				break
			}
			char = rune(line[i])
		}
		if word.Len() > 0 {
			val, err := strconv.Atoi(word.String())
			checkError(err)
			output = append(output, val)
		}
	}
	return output
}

func absoluteVal(num1 int, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}

type LevelState int

const (
	StateNew LevelState = iota
	StateDecrease
	StateIncrease
)

func main() {
	args := os.Args
	filename, err := parseArgs(args)
	checkError(err)
	data, err := os.ReadFile(filename)
	checkError(err)
	input := string(data)
	fmt.Println(input)
	fmt.Println(q1(input))
	fmt.Println(q2(input))
}

func q1(input string) int {
	safeCount := 0
	for _, entry := range strings.Split(input, "\r\n") {
		sline := stripLine(entry)
		isSafe := true
		state := StateNew
		for i := 0; i < len(sline)-1; i++ {
			var newState LevelState
			if sline[i] > sline[i+1] {
				newState = StateDecrease
			} else if sline[i] < sline[i+1] {
				newState = StateIncrease
			} else {
				isSafe = false
				break
			}
			if absoluteVal(sline[i], sline[i+1]) > 3 {
				isSafe = false
			}
			if newState != state && state != StateNew {
				isSafe = false
				break
			}
			state = newState
		}
		if isSafe {
			safeCount++
		}
	}
	return safeCount
}

func isLineSafe(sline []int) bool {
	isSafe := true
	state := StateNew
	for i := 0; i < len(sline)-1; i++ {
		var newState LevelState
		if sline[i] > sline[i+1] {
			newState = StateDecrease
		} else if sline[i] < sline[i+1] {
			newState = StateIncrease
		} else {
			isSafe = false
			break
		}
		if absoluteVal(sline[i], sline[i+1]) > 3 {
			isSafe = false
		}
		if newState != state && state != StateNew {
			isSafe = false
			break
		}
		state = newState
	}
	return isSafe
}

func q2(input string) int {
	safeCount := 0
	for _, entry := range strings.Split(input, "\r\n") {
		sline := stripLine(entry)
		isSafe := false
		isSafe = isLineSafe(sline)
		if !isSafe {
			for i := range sline {
				newIndex := 0
				newLine := make([]int, len(sline)-1)
				for j := range sline {
					if i != j {
						newLine[newIndex] = sline[j]
						newIndex++
					}
				}
				isSafe = isLineSafe(newLine)
				if isSafe {
					break
				}
			}
		}
		if isSafe {
			safeCount++
		}
	}
	return safeCount
}
