package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    const (
        RED = 12
        GREEN = 13
        BLUE = 14
    )
    sum_of_IDs := 0

    file, _ := os.Open("test_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)
    for scn.Scan() {
        var isValidGame bool
        line := scn.Text()

        indexOfColon := strings.Index(line, ":")
        id, _ := strconv.Atoi(line[5:indexOfColon])
        line = line[indexOfColon+2:]
        // fmt.Printf("id=%d\n", id)

        slices := strings.Split(line, ";")
        for _, slice := range slices {
            num := 0
            for i:=0; i<len(slice); i++ {
                if slice[i] >= '0' && slice[i] <= '9' {
                    num = num*10 + int(slice[i]) - 48
                    continue
                }
                if slice[i] == 'r' && slice[i+1] == 'e' && slice[i+2] == 'd' {
                    isValidGame = num < RED
                    continue
                } else if slice[i] == 'g' && slice[i+1] == 'r' && slice[i+2] == 'e' {
                    isValidGame = num < GREEN
                    continue
                } else if slice[i] == 'b' {
                    isValidGame = num < BLUE
                    continue
                }

                if slice[i] == ',' {
                    slice = slice[i+1:]
                    i = 0; num = 0
                }
            }
        }

        if isValidGame {
            fmt.Printf("Valid Id is: %d\n", id)
            sum_of_IDs += id
        }

        // fmt.Printf("%q\n", slices)
    }

    fmt.Println("Ans is:", sum_of_IDs)
}
