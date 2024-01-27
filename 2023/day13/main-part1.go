package main

import (
    "fmt"
    "os"
    "bufio"
)

func calcHorizontal(s* []string) (int, bool) {
    noOfRows := -1
    isHori := false

    for i := 0; i < len(*s)-1; i++ {
        if (*s)[i] == (*s)[i+1] {
            isHori = true
            noOfRows = i + 1
            break
        }
    }

    return noOfRows, isHori
}

func calcVertical(s* []string) int {
    var sliceOfCols []string

    for i := 0; i < len((*s)[0]); i++ {
        str := ""
        for j := 0; j < len(*s); j++ {
            str += string((*s)[j][i])
        }
        sliceOfCols = append(sliceOfCols, str)
    }

    noCol, _ := calcHorizontal(&sliceOfCols)
    return noCol
}

func main() {
    file, _ := os.Open("test_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    var sliceOfStrings []string

    var isHorizontal bool
    var noOfRows int
    var noOfCols int
    var finalAns int = 0
    var counter int = 0

    for scn.Scan() {
        line := scn.Text()

        if line == "" {
            counter++
            var temp int
            temp, isHorizontal = calcHorizontal(&sliceOfStrings)
            if isHorizontal == false {
                noOfCols = calcVertical(&sliceOfStrings)
            } else {
                noOfRows = temp
                noOfRows *= 100
            }

            if counter != 0 && counter % 2 == 0 {
                finalAns = finalAns + noOfCols + (noOfRows * 100)
            }

            //removing all strings from slice
            sliceOfStrings = sliceOfStrings[:0]
        } else {
            sliceOfStrings = append(sliceOfStrings, line)
        }

    }
    fmt.Println(isHorizontal)
    fmt.Println(noOfRows)
    fmt.Println(noOfCols)
    fmt.Println("final ANs: ", finalAns)
}
//col + rows*100
