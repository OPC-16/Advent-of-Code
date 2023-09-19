package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var fullyContainableElfCount int

    for scanner.Scan() {
        var firstStart, firstEnd, secondStart, secondEnd int
        fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)

        if firstStart <= secondStart && firstEnd >= secondEnd || secondStart <= firstStart && secondEnd >= firstEnd {
            fullyContainableElfCount++
        }
    }

    fmt.Println(fullyContainableElfCount)
} //function-main
