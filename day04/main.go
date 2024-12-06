package main

import(
	"os"
	"fmt"
	aoc2024helper "github.com/reneichhorn/aoc2024Helper"
)

func main(){
	args := os.Args
	filename, err := aoc2024helper.ParseArgs(args)
	data, err := os.ReadFile(filename)
	aoc2024helper.CheckError(err)
	fmt.Println(q1(string(data)))
	fmt.Println(q2(string(data)))
}

func q1(input string) int {
	sum := 0
	return sum
}

func q2(input string) int {
	sum := 0
	return sum
}