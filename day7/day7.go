package day7

import (
	"adventofcode-go-2022/util"
	"fmt"
	"strings"
)

func Part1() int {
	lines := util.ReadFile("day7/input-ex.txt")
	buildTree(lines)
	return 0
}

type directory struct {
	name        string
	directories []directory
	files       []file
}

type file struct {
	size int
	name string
}

func buildTree(lines []string) {

	var filesystem []directory

	cwd := "/"
	filesystem = append(filesystem, directory{
		name: cwd,
	})

	path := filesystem[0]

	for i, line := range lines {
		tokens := strings.Fields(line)

		switch tokens[0] {

		case "$":
			fmt.Printf("%d: %s - COMMAND\n", i, line)
			if tokens[1] == "cd" {
				if tokens[2] == "/" {
					path = filesystem[0]
					continue
				}

				if tokens[2] == ".." {
					//go up
					continue
				}

				for _, d := range path.directories {
					if d.name == tokens[2] {
						path = d
					}
				}
			}

		case "dir":
			fmt.Printf("%d: %s - DIRECTORY\n", i, line)
			path.directories = append(path.directories, directory{
				name: tokens[1],
			})

		default:
			d := util.ParseInt(tokens[0])
			fmt.Printf("%d: %s - FILE(%d)\n", i, line, d)
			path.files = append(path.files, file{
				size: util.ParseInt(tokens[0]),
				name: tokens[1],
			})
		}
	}

	fmt.Println(filesystem)
}
