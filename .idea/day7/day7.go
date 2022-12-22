package day7

import (
	"adventofcode-go-2022/util"
	"math"
	"strings"
)

func Part1() int {
	root := buildTree(util.ReadFile("day7/input.txt"))
	return sumDirsOneHundredKilo(root)
}

func Part2() int {
	root := buildTree(util.ReadFile("day7/input.txt"))
	totalDiskSpace := 70000000
	spaceNeeded := 30000000

	directoryMinSize := spaceNeeded - (totalDiskSpace - root.totalSize)
	return findSmallestDir(root, directoryMinSize)
}

func findSmallestDir(itr *dir, dirMinSize int) int {
	smallest := math.MaxInt64
	if itr.totalSize >= dirMinSize {
		smallest = util.MinInt(smallest, itr.totalSize)
	}

	for _, childDirs := range itr.childDir {
		smallest = util.MinInt(smallest, findSmallestDir(childDirs, dirMinSize))
	}

	return smallest
}

func sumDirsOneHundredKilo(itr *dir) int {
	sum := 0
	if itr.totalSize <= 100000 {
		sum += itr.totalSize
	}
	for _, child := range itr.childDir {
		sum += sumDirsOneHundredKilo(child)
	}
	return sum
}

type dir struct {
	name      string
	parentDir *dir
	childDir  map[string]*dir
	files     map[string]int
	totalSize int
}

func buildTree(cmds []string) *dir {
	root := &dir{
		name:     "root",
		childDir: map[string]*dir{},
		files:    map[string]int{},
	}
	itr := root

	c := 0
	for c < len(cmds) {
		switch cmd := cmds[c]; cmd[0:1] {
		case "$":
			if cmd == "$ ls" {
				c++
			} else {
				changeDir := strings.Split(cmd, "cd ")[1]
				changeDir = strings.TrimSpace(changeDir)

				if changeDir == ".." {
					itr = itr.parentDir
				} else {
					if _, ok := itr.childDir[changeDir]; !ok {
						itr.childDir[changeDir] = &dir{
							name:      changeDir,
							parentDir: itr,
							childDir:  map[string]*dir{},
							files:     map[string]int{},
						}
					}
					itr = itr.childDir[changeDir]
				}
				c++
			}
		default:
			if strings.HasPrefix(cmd, "dir") {
				childDirName := cmd[4:]
				if _, ok := itr.childDir[childDirName]; !ok {
					itr.childDir[childDirName] = &dir{
						name:      childDirName,
						parentDir: itr,
						childDir:  map[string]*dir{},
						files:     map[string]int{},
					}
				}
			} else {
				filename := strings.Split(cmd, " ")
				itr.files[filename[0]] = util.ParseInt(filename[0])
			}
			c++
		}
	}
	calculateDirSize(root)
	return root
}

func calculateDirSize(itr *dir) int {
	totalSize := 0
	for _, childItr := range itr.childDir {
		totalSize += calculateDirSize(childItr)
	}
	for _, size := range itr.files {
		totalSize += size
	}
	itr.totalSize = totalSize
	return totalSize
}
