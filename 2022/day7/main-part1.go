package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Dir struct {
    Name     string
    Children []string
    Size     int
}

func stringToDir(str string) *Dir {
    for _, d := range dirs {
        if str == d.Name {
            return &d
        }
    }
    return &Dir{}
}

func calcDirSize(currDir *Dir) int {
    //base case
    if len(currDir.Children) == 0 {
        return currDir.Size
    }

    for i := 0; i < len(currDir.Children); i++ {
        currDir.Size += calcDirSize(stringToDir(currDir.Children[i]))
    }

    return currDir.Size
}

var dirs []Dir

func main() {
    file, _ := os.Open("test_input.txt")
    defer file.Close()

    sc := bufio.NewScanner(file)

    for sc.Scan() {
        line := sc.Text()

        if line[0] == '$' {     //if that line is a command
            if line == "$ ls" || line == "$ cd .." {
                continue
            } else {
                dir := Dir{
                    Name: line[5:],
                }
                dirs = append(dirs, dir)
            } 
        } else {
            if line[0] == 'd' {     //add child directories
                child := line[4:]
                dirs[len(dirs)-1].Children = append(dirs[len(dirs)-1].Children, child)
            } else {        //add the sizes of files
                var i int
                for i=0; i<len(line); i++ {
                    if line[i] == ' ' {
                        break
                    }
                }

                size, _ := strconv.Atoi(line[:i])
                dirs[len(dirs)-1].Size += size
            }
        }
    }

    for _, dir := range dirs {
        fmt.Printf("DirName: %s\n   Size: %d\n   Children: %v\n", dir.Name, dir.Size, dir.Children)
    }

    //TODO: name, size, children even totalSIze is correct but the changes in size are not reflected in the original dirs slice
    totalSize := calcDirSize(&dirs[0])

    fmt.Printf("\n------------------------------ totalSize = %d\n", totalSize)
    for _, dir := range dirs {
        fmt.Printf("DirName: %s\n   Size: %d\n   Children: %v\n", dir.Name, dir.Size, dir.Children)
    }
}
