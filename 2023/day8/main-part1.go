package main

import (
    "fmt"
    "os"
    "bufio"
)

type leftRight struct {
    left  string
    right string
}

func main() {
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    //scanning the instruction part of input
    scn.Scan()
    instructions := scn.Text()
    fmt.Println("instructions:", instructions)
    scn.Scan()

    //a Map, key will be src string and value will be struct with leftDest & rightDest
    Map := make(map[string]leftRight)

    isFirst := true
    var startingNode string
    for scn.Scan() {
        line := scn.Text()

        src := string(line[0]) + string(line[1]) + string(line[2])
        if isFirst {
            startingNode = src
            isFirst = false
            fmt.Printf("starting node is: %s\n", startingNode)
        }
        leftDest := string(line[7]) + string(line[8]) + string(line[9])
        rightDest := string(line[12]) + string(line[13]) + string(line[14])

        Map[src] = leftRight{leftDest, rightDest}
        fmt.Printf("src: %s, left: %s, right: %s\n", src, leftDest, rightDest)
    }

    currNode := startingNode
    numOfSteps := 0
    i := 0
    for {
        if instructions[i] == 'R' {
            currNode = Map[currNode].right
        } else {
            currNode = Map[currNode].left
        }
        numOfSteps++
        i++

        if currNode == "ZZZ" {
            break
        }
        if i >= len(instructions) {
            i = 0
        }
    }

    fmt.Println("number of steps:", numOfSteps)
    // fmt.Println(Map)
}
