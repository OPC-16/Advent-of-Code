package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    Map := make(map[int]int)
    criteriaMap := make(map[int]int)

    //scanning first line
    scn.Scan()
    firstLine := scn.Text()

    indexOfColon := strings.Index(firstLine, ":")
    firstLine = firstLine[indexOfColon+2:]

    seeds := strings.Split(firstLine, " ")
    for i:=0; i<len(seeds); i++ {
        num, _ := strconv.Atoi(seeds[i])
        Map[num] = -1
    }

    fmt.Println("Scanned first line")
    for scn.Scan() {
        fmt.Println("Loop started")
        line := scn.Text()

        if line == "" {
            continue
        }
        if line[0] >= '9' {
            for key, _ := range Map {
                v, present := criteriaMap[key]
                if present {
                    Map[key] = v
                } else {
                    Map[key] = key
                }
            }

            //here, we put values of Map as keys
            Map_copy := Map
            for key, value := range Map_copy {
                delete(Map, key)
                Map[value] = -1
            }

            //here, we clear all the key-value pair from the criteria map
            criteriaMap = make(map[int]int)

            continue
        }

        var destRangeStart, sourceRangeStart, rangeLength int
        fmt.Sscanf(scn.Text(), "%d %d %d", &destRangeStart, &sourceRangeStart, &rangeLength)

        for i:=0; i<rangeLength; i++ {
            criteriaMap[sourceRangeStart + i] = destRangeStart + i
        }
        fmt.Println("Loop ended")
    }

    fmt.Printf("\n")
    fmt.Println(Map)
}

func convertToInts(s []string) []int {
    var nums []int
    for i:=0; i<len(s); i++ {
        num, _ := strconv.Atoi(s[i])
        nums = append(nums, num)
    }

    return nums
}
