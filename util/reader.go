package util

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	fixPath()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func fixPath() {
	wd, _ := os.Getwd()
	//fmt.Println("CURRENT WD")
	//fmt.Println(wd)
	if strings.HasSuffix(wd[:len(wd)-1], "day") {
		os.Chdir("..")
	}
}
