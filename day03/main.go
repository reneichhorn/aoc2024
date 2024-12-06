package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	aoc2024helper "github.com/reneichhorn/aoc2024Helper"
)

type Token int

const (
	OP Token = iota
	OpenP
	Num1
	Comma
	Num2
	CloseP
	Done
)

func CurInBound(cur int, inputLength int) bool {
	return cur < inputLength
}

type calc struct {
	num1    int
	num2    int
	enabled bool
	op      string
}

func getValidInput(input string) []calc {
	var output []calc
	validOPs := []string{"mul"}
	tokens := map[Token]string{
		OpenP:  "(",
		Comma:  ",",
		CloseP: ")",
	}
	lCur := 0
	lCurW := 7
	rCur := 0
	inputLength := len(input)
	enabled := true
	for {
		if !CurInBound(lCur, inputLength) {
			break
		}
		if !CurInBound(lCur+lCurW, inputLength) {
			break
		}
		sel := input[lCur : lCur+lCurW]
		fmt.Println("Selection:", sel, validOPs[:], lCur)
		if strings.Contains(sel, "do()") {
			enabled = true
			lCur += 5
		}
		if strings.Contains(sel, "don't()") {
			enabled = false
			lCur += 8
		}
		lCurW = 3
		sel = input[lCur : lCur+lCurW]
		if slices.Contains(validOPs[:], sel) {
			state := OP
			rCur = lCur + lCurW
			var num1, num2 strings.Builder
			op := sel
			running := true
			for running {
				if !CurInBound(rCur, len(input)) {
					break
				}
				curSel := string(input[rCur])
				fmt.Println(state, curSel)
				switch state {
				case OP:
					if curSel == tokens[OpenP] {
						state = OpenP
					} else {
						running = false
					}
					break
				case OpenP:
					if strings.ContainsAny(curSel, "0123456789") {
						state = Num1
						num1.WriteString(curSel)
					} else {
						running = false
					}
					break
				case Num1:
					if curSel == "," {
						state = Comma
					} else if strings.ContainsAny(curSel, "0123456789") {
						state = Num1
						num1.WriteString(curSel)
					} else {
						running = false
					}
				case Comma:
					if strings.ContainsAny(curSel, "0123456789") {
						state = Num2
						num2.WriteString(curSel)
					} else {
						running = false
					}
					break
				case Num2:
					if curSel == ")" {
						state = CloseP
					} else if strings.ContainsAny(curSel, "0123456789") {
						state = Num2
						num2.WriteString(curSel)
					} else {
						running = false
					}
				case CloseP:
					intNum1, err := strconv.Atoi(num1.String())
					aoc2024helper.CheckError(err)
					intNum2, err := strconv.Atoi(num2.String())
					aoc2024helper.CheckError(err)
					c := calc{num1: intNum1, num2: intNum2, op: op, enabled: enabled}
					output = append(output, c)
					running = false
				}
				rCur++
				if !running {
					op = ""
					num1.Reset()
					num2.Reset()
				}
			}
		}
		lCur++
		lCurW = 7
	}
	return output
}

func main() {
	args := os.Args
	filename, err := aoc2024helper.ParseArgs(args)
	aoc2024helper.CheckError(err)
	data, err := os.ReadFile(filename)
	fmt.Println(q1(string(data)))
	fmt.Println(q2(string(data)))
}

func q1(input string) int {
	calcs := getValidInput(input)
	fmt.Println(calcs)
	sum := 0
	for _, calc := range calcs {
		switch calc.op {
		case "mul":
			sum += calc.num1 * calc.num2
		}
	}
	return sum
}

func q2(input string) int {
	calcs := getValidInput(input)
	sum := 0
	for _, calc := range calcs {
		if calc.enabled {
			switch calc.op {
			case "mul":
				sum += calc.num1 * calc.num2
			}
		}
	}
	return sum
}
