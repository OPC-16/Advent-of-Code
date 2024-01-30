package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateHash2(s string) (int, bool) {
    var currValue int = 0
    for i := 0; i < len(s); i++ {
        if s[i] == 61 || s[i] == 45 {
            break
        }
        currValue += int(s[i])
        currValue *= 17
        currValue %= 256
    }

    if s[len(s)-1] == 45 {
        return currValue, false  //return false for '-' at last
    } else {
        return currValue, true  //return true for a number at last
    }
}

type Box struct {
    boxNum      int
    slot        int
    focalLength int
}

func main() {
    file, _ := os.Open("test_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    scn.Scan()
    line := scn.Text()

    var finalAns int = 0
    steps := strings.Split(line, ",")
    stringToIntMap := make(map[string]int)
    stringToBox := make(map[string]Box)

    for i := 0; i < len(steps); i++ {
        var hash int
        var hasFocal bool
        if _, ok := stringToIntMap[steps[i]]; !ok {
            hash, hasFocal = calculateHash2(steps[i])
            stringToIntMap[steps[i]] = hash
        }

        if hasFocal == true {
            box := Box{ boxNum: stringToIntMap[steps[i]]-1, slot: 0, focalLength: int(steps[i][len(steps[i])-1])}
            str := steps[i][:len(steps[i])-2]

            //deciding slot for current string
            slot := 1
            for _, v := range stringToBox {
                tempBox := v
                if tempBox.boxNum == box.boxNum {
                    slot++
                }
            }
            box.slot = slot

            stringToBox[str] = box
        } else {
            str := steps[i][:len(steps[i])-1]
            delete(stringToBox, str)

            boxNum := stringToIntMap[steps[i]]-1
            for _, v := range stringToBox {
                if boxNum == v.boxNum {
                    v.slot--
                }
            }
        }
    }


    fmt.Println(stringToBox)
    fmt.Println("Final Ans:", finalAns)
}
