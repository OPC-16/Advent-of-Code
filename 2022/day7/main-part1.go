package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "strings"
)

type Dir struct {
    Name     string
    Size     int
    isFile bool
    sons map[string]*Dir
    father *Dir
}

func main() {
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    sc := bufio.NewScanner(file)

    var currDir *Dir
    dirs := []*Dir{}

    for sc.Scan() {
        line := strings.Fields(sc.Text())
        if len(line) > 2 {
            if line[2] == ".." {
                currDir = currDir.father
            } else if line[2] == "/" {
                currDir = &Dir{"/", 0, false, make(map[string]*Dir), nil}
            } else {
                currDir = currDir.sons[line[2]]
            }
        } else if line[0] == "dir" {
            currDir.sons[line[1]] = &Dir{line[1], 0, false, make(map[string]*Dir), currDir}
            dirs = append(dirs, currDir.sons[line[1]])
        } else if line[0] != "$" {
            size, _ := strconv.Atoi(line[0])
            currDir.sons[line[1]] = &Dir{line[1], size, true, nil, currDir}
        }
    }

    var totalSize int

    for _, dir := range dirs {
        size := calcSize(*dir)
        if size <= 100000 {
            totalSize += size
        }
    }

    fmt.Println(totalSize)
}

func calcSize(root Dir) (size int) {
    if root.isFile {
        return root.Size
    }

    for _, d := range root.sons {
        size += calcSize(*d)
    }

    return size
}
