package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
    total := 0
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    for scn.Scan() {
        line := scn.Text()

        //converting string into slice of runes
        runes := []rune{}
        for len(line) > 0 {
            r, size := utf8.DecodeRuneInString(line)
            runes = append(runes, r)
            line = line[size:]
        }

        var num int
        for i := 0; i < len(runes); i++ {
            //check for 'one'
            if runes[i] == 'o' {
                if runes[i+1] == 'n' && runes[i+2] == 'e' {
                    num = 1
                    break
                }
            }
            //check for 'two'
            if runes[i] == 't' {
                if runes[i+1] == 'w' && runes[i+2] == 'o' {
                    num = 2
                    break
                }
            }
            //check for 'three'
            if runes[i] == 't' {
                if runes[i+1] == 'h' && runes[i+2] == 'r' && runes[i+3] == 'e' && runes[i+4] == 'e' {
                    num = 3
                    break
                }
            }
            //check for 'four'
            if runes[i] == 'f' {
                if runes[i+1] == 'o' && runes[i+2] == 'u' && runes[i+3] == 'r' {
                    num = 4
                    break
                }
            }
            //check for 'five'
            if runes[i] == 'f' {
                if runes[i+1] == 'i' && runes[i+2] == 'v' && runes[i+3] == 'e' {
                    num = 5
                    break
                }
            }
            //check for 'six'
            if runes[i] == 's' {
                if runes[i+1] == 'i' && runes[i+2] == 'x' {
                    num = 6
                    break
                }
            }
            //check for 'seven'
            if runes[i] == 's' {
                if runes[i+1] == 'e' && runes[i+2] == 'v' && runes[i+3] == 'e' && runes[i+4] == 'n' {
                    num = 7
                    break
                }
            }
            //check for 'eight'
            if runes[i] == 'e' {
                if runes[i+1] == 'i' && runes[i+2] == 'g' && runes[i+3] == 'h' && runes[i+4] == 't' {
                    num = 8
                    break
                }
            }
            //check for nine
            if runes[i] == 'n' {
                if runes[i+1] == 'i' && runes[i+2] == 'n' && runes[i+3] == 'e' {
                    num = 9
                    break
                }
            }

            //check for integers
            if runes[i] >= '1' && runes[i] <= '9' {
                num = int(runes[i]) - 48
                break
            }
        }
        for i := len(runes)-1; i >= 0; i-- {
            //check for one
            if runes[i] == 'e' {
                if runes[i-1] == 'n' && runes[i-2] == 'o' {
                    num = num*10 + 1
                    break
                }
            }

            //check for two
            if runes[i] == 'o' {
                if runes[i-1] == 'w' && runes[i-2] == 't' {
                    num = num*10 + 2
                    break
                }
            }

            //check for three
            if runes[i] == 'e' {
                if runes[i-1] == 'e' && runes[i-2] == 'r' && runes[i-3] == 'h' && runes[i-4] == 't' {
                    num = num*10 + 3
                    break
                }
            }

            //check for four
            if runes[i] == 'r' {
                if runes[i-1] == 'u' && runes[i-2] == 'o' && runes[i-3] == 'f' {
                    num = num*10 + 4
                    break
                }
            }

            //check for five
            if runes[i] == 'e' {
                if runes[i-1] == 'v' && runes[i-2] == 'i' && runes[i-3] == 'f' {
                    num = num*10 + 5
                    break
                }
            }

            //check for six
            if runes[i] == 'x' {
                if runes[i-1] == 'i' && runes[i-2] == 's' {
                    num = num*10 + 6
                    break
                }
            }

            //check for seven
            if runes[i] == 'n' {
                if runes[i-1] == 'e' && runes[i-2] == 'v' && runes[i-3] == 'e' && runes[i-4] == 's' {
                    num = num*10 + 7
                    break
                }
            }

            //check for eight
            if runes[i] == 't' {
                if runes[i-1] == 'h' && runes[i-2] == 'g' && runes[i-3] == 'i' && runes[i-4] == 'e' {
                    num = num*10 + 8
                    break
                }
            }

            //check for nine
            if runes[i] == 'e' {
                if runes[i-1] == 'n' && runes[i-2] == 'i' && runes[i-3] == 'n' {
                    num = num*10 + 9
                    break
                }
            }


            //check for integers
            if runes[i] >= '1' && runes[i] <= '9' {
                num = num*10 + int(runes[i]) - 48
                break
            }
        }

        total += num

    }//scanner loop

    fmt.Printf("Total is: %d\n", total)
}
