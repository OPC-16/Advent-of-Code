package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// each characters' position can be defined by row and column
type pointer struct {
    row    int
    column int
}

func main() {
    file, _ := os.Open("test_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)
    var sliceOfStrings []string
    var pointerToStart pointer
    rowNum := 0

    //storing each string in slice and marking the position of 'S'
    for scn.Scan() {
        line := scn.Text()

        sliceOfStrings = append(sliceOfStrings, line)

        indexOfStart := strings.Index(line, "S")
        if indexOfStart != -1 {
            pointerToStart.row = rowNum
            pointerToStart.column = indexOfStart
        }
        rowNum++
    }

    fmt.Printf("%q\n", sliceOfStrings)
    fmt.Println("index of S:", pointerToStart.row, pointerToStart.column)

    //pointer for two directions
    var firstPointer pointer
    var secondPointer = pointer{row: -1, column: -1}

    var top = pointer{row: pointerToStart.row-1, column: pointerToStart.column}
    var bottom = pointer{row: pointerToStart.row+1, column: pointerToStart.column}
    var left = pointer{row: pointerToStart.row, column: pointerToStart.column-1}
    var right = pointer{row: pointerToStart.row, column: pointerToStart.column+1}

    var topChar byte
    if top.row != -1 {
        topChar = sliceOfStrings[top.row][top.column]
    }
    var bottomChar byte
    if bottom.row != -1 {
        bottomChar = sliceOfStrings[bottom.row][bottom.column]
    }
    var leftChar byte
    if left.column != -1 {
        leftChar = sliceOfStrings[left.row][left.column]
    }
    var rightChar byte
    if right.column != -1 {
       rightChar = sliceOfStrings[right.row][right.column]
    }

    if top.row != -1 && topChar == '7' && topChar == 'F' && topChar == '|' {
        firstPointer.row = top.row
        firstPointer.column = top.column
    }
    if bottom.row != -1 && bottomChar == '|' && bottomChar == 'L' && bottomChar == 'J' {
        firstPointer.row = bottom.row
        firstPointer.column = bottom.column
    }
    if left.column != -1 && leftChar == '-' && leftChar == 'L' && leftChar == 'F' {
        secondPointer.row = left.row
        secondPointer.column = left.column
    }
    if right.column != -1 && rightChar == '-' && rightChar == 'J' && rightChar == '7' {
        secondPointer.row = right.row
        secondPointer.column = right.column
    }

    // for firstPointer.row != secondPointer.row && firstPointer.column != secondPointer.column {
    // }

    fmt.Printf("firstPointer, row=%d, column=%d\n", firstPointer.row, firstPointer.column)
    fmt.Printf("secondPointer, row=%d, column=%d\n", secondPointer.row, secondPointer.column)
}
