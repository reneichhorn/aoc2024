package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func getNextName() string {
	entries, err := os.ReadDir(".")
	checkError(err)
	var dirs []int
	for _, e := range entries {
		if e.IsDir() && strings.Contains(e.Name(), "day") {
			daynumber, err := strconv.Atoi(e.Name()[len(e.Name())-2:])
			checkError(err)
			dirs = append(dirs, daynumber)
		}
	}
	slices.Sort(dirs[:])
	highest := dirs[len(dirs)-1]
	var newName strings.Builder
	newName.WriteString("day")
	if highest < 9 {
		newName.WriteString("0")
	}
	newName.WriteString(strconv.Itoa(highest + 1))
	return newName.String()
}

func createEmptyFile(filename string) error {
	d := []byte("")
	if strings.Contains(filename, "main.go") {
		var contents strings.Builder
		contents.WriteString("package main\n\nimport(\n\t\"os\"\n\t\"fmt\"\n\taoc2024helper \"github.com/reneichhorn/aoc2024Helper\"\n)\n\nfunc main(){\n\targs := os.Args\n\tfilename, err := aoc2024helper.ParseArgs(args)\n\tdata, err := os.ReadFile(filename)\n\taoc2024helper.CheckError(err)\n\tfmt.Println(q1(string(data)))\n\tfmt.Println(q2(string(data)))\n}\n\nfunc q1(input string) int {\n\tsum := 0\n\treturn sum\n}\n\nfunc q2(input string) int {\n\tsum := 0\n\treturn sum\n}")
		d = []byte(contents.String())
	}
	return os.WriteFile(filename, d, 0644)
}

func main() {
	// TODO: Create Following structure
	//	- dayXX
	//		|
	//		L main.go
	//      L input.txt
	//      L test.txt
	// XX should be the number of the day
	// name is taken from the last folder found +1
	next := getNextName()
	err := os.Mkdir(next, 0755)
	checkError(err)
	filesToCreate := []string{"main.go", "input.txt", "test.txt"}
	for _, f := range filesToCreate {
		var filepath strings.Builder
		filepath.WriteString(next)
		filepath.WriteString("/")
		filepath.WriteString(f)
		checkError(createEmptyFile(filepath.String()))
	}
}
