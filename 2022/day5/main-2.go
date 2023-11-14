package main

import (
    "fmt"
	"bufio"
	"os"
)

type stack struct {
    elements []rune
}

func (s *stack) push(r rune) {
    s.elements = append(s.elements, r)
}

func (s *stack) pop() rune {
    r := s.elements[len(s.elements) - 1]
    s.elements = s.elements[:len(s.elements) - 1]
    return r
}

func (s *stack) addToBottom(r rune) {
    s.elements = append([]rune {r}, s.elements...)
}

func (s stack) convertToString() string {
    var str string
    for _, r := range s.elements {
        str += string(r) + " "
    }
    return str
}

func main() {
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    sc := bufio.NewScanner(file)

    //create slice of stacks
    stacks := make([]stack, 9)

    //parsing the input
    sc.Scan()
    for sc.Text() != " 1   2   3   4   5   6   7   8   9 " {
        for i, r := range sc.Text() {
            if r != ' ' && r != '[' && r != ']' {
                stacks[i/4].addToBottom(r)
            }
        }
        sc.Scan()
    }

    //reading the empty line
    sc.Scan()

    for sc.Scan() {
        var howManyToMove, fromMove, toMove int
        fmt.Sscanf(sc.Text(), "move %d from %d to %d", &howManyToMove, &fromMove, &toMove)

        var tempStack stack
        //moving crates to temp stack in reverse order
        for i := 0; i < howManyToMove; i++ {
            tempStack.push(stacks[fromMove-1].pop())
        }

        //moving crates from temp stack to destination stack in again reverse order so that original order is achieved (coz we reversed it 2 times!)
        for i := 0; i < howManyToMove; i++ {
            stacks[toMove-1].push(tempStack.pop())
        }
    }

    for _, s := range stacks {
        fmt.Print(string(s.pop()))
    }
    
}
