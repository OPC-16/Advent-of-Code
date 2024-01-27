package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    file, _ := os.Open("test_input.txt")
    defer file.Close()

    totalPermutations := 0

    scn := bufio.NewScanner(file)

    for scn.Scan() {
        line := scn.Text()
        // permutation := 0

        fmt.Printf("%q\n", strings.Split(line, " "))
    }

    fmt.Println("Final ans: ", totalPermutations)
}
