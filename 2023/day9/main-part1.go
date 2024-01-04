package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
    file, _ := os.Open("test_input.txt")
    defer file.Close()

    finalAns := 0
    scn := bufio.NewScanner(file)

    for scn.Scan() {
        line := scn.Text()

        //splitting in strings and then converting them into ints
        sliceOfStrings := strings.Split(line, " ")
        sliceOfInts := convertToInts(&sliceOfStrings)

        // fmt.Printf("%v\n", sliceOfInts)

        //creating slice of slices of ints for calculating each input line's differences
        var sliceOfSlicesOfInts [][]int
        sliceOfSlicesOfInts = append(sliceOfSlicesOfInts, sliceOfInts)
        calculateDiff(&sliceOfSlicesOfInts)

        //adding zero to last slice at last index
        sizeOfLastSlice := len(sliceOfSlicesOfInts[len(sliceOfSlicesOfInts)-1])
        sliceWithExtraZero := make([]int, sizeOfLastSlice+1)
        sliceOfSlicesOfInts[len(sliceOfSlicesOfInts)-1] = sliceWithExtraZero //replacing last slice with extra zero wali slice

        fmt.Println(sliceOfSlicesOfInts)

        //calculating the desired element
        for i:=len(sliceOfSlicesOfInts)-1; i>=1; i-- {
            currSlice := sliceOfSlicesOfInts[i]
            prevSlice := sliceOfSlicesOfInts[i-1]
            elementToBeAdded := currSlice[len(currSlice)-1] + prevSlice[len(prevSlice)-1]
            prevSlice = append(prevSlice, elementToBeAdded)

            if i == 1 {
                finalAns += elementToBeAdded
            }
        }
    }

    fmt.Println("finalAns:", finalAns)
}//main

func convertToInts(s* []string) []int {
    var sliceOfInts []int

    for _, s := range *s {
        num, _ := strconv.Atoi(s)
        sliceOfInts = append(sliceOfInts, num)
    }
    return sliceOfInts
}

func calculateDiff(ssInts *[][]int) {
    currSlice := (*ssInts)[0]
    for !isSliceZero(&currSlice) {
        var diffSlice []int
        for i:=0; i<len(currSlice)-1; i++ {
            diffSlice = append(diffSlice, currSlice[i+1] - currSlice[i])
        }
        *ssInts = append(*ssInts, diffSlice)
        currSlice = diffSlice
    }
}

func isSliceZero(s *[]int) bool {
    temp := make([]int, len(*s))    //this will create a slice of size same as arg slice and initialize with zeros
    result := slices.Compare(temp, *s)

    if result == 0 {
        return true
    } else {
        return false
    }
}
