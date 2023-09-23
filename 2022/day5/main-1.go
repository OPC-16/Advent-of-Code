package main

import (
    "fmt"
	"bufio"
	"os"
)

type Stack struct {
    elements []rune
}

func (s *Stack) push(r rune) {
    s.elements = append(s.elements, r)
}

func (s *Stack) pop() rune {
    r := s.elements[len(s.elements) - 1]
    s.elements = s.elements[:len(s.elements) - 1]
    return r
}

func (s *Stack) addToBottom(r rune) {
    s.elements = append([]rune {r}, s.elements...)
}

func (s Stack) String() string {
    var str string
    for _, r := range s.elements {
        str += string(r) + " "
    }
    return str
}
func main() {
    file, _ := os.Open("test_input.txt")

    defer file.Close()
    
    sc := bufio.NewScanner(file)

    //create slice of Stacks
    stacks := make([]Stack, 9)

    sc.Scan()
    for sc.Text() != " 1   2   3   4   5   6   7   8   9 " {
        for i, r := range sc.Text() {
            if r != ' ' && r != '[' && r != ']' {
                stacks[i/4].addToBottom(r)
            }
        }
        sc.Scan()
    }

    //read the empty line
    sc.Scan()

    for sc.Scan() {
        var toMove, from, to int
        fmt.Sscanf(sc.Text(), "move %d from %d to %d", &toMove, &from, &to)

        //move elements one by one
        for i := 0; i < toMove; i++ {
            r := stacks[from-1].pop()
            stacks[to-1].push(r)
        }
    }

    // for _, s := range stacks {
    //     fmt.Println(string(s.pop()))
    // }
    fmt.Println("sldkfj")
}
