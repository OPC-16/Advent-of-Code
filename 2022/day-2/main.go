// In opponent's case:
//A for Rock
//B for Paper
//C for Scissors

// FOR PART-1
   // In our case:
   //X for Rock
   //Y for Paper
   //Z for Scissors

// FOR PART-2
   // In our case:
   //X means loose
   //Y means draw
   //Z means win
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
   total_score := 0

   file, err := os.Open("puzzle_input.txt")
   if err != nil {
      fmt.Println(err)
      return
   }
   defer file.Close()

   // Create a scanner to read lines from the file
   scanner := bufio.NewScanner(file)

   for scanner.Scan() {
      line := scanner.Text()

      shapes := []rune (line)
      if len(shapes) >= 3 {
         if shapes[0] == 'A' && shapes[2] == 'X' {       //winning
            total_score = total_score + 3 + 0
         } else if shapes[0] == 'A' && shapes[2] == 'Y' {      //winning
            total_score = total_score + 1 + 3
         } else if shapes[0] == 'A' && shapes[2] == 'Z' {      //winning
            total_score = total_score + 2 + 6
         } else if shapes[0] == 'B' && shapes[2] == 'X' {      //losing
            total_score = total_score + 1 + 0
         } else if shapes[0] == 'B' && shapes[2] == 'Y' {      //losing
            total_score = total_score + 2 + 3
         } else if shapes[0] == 'B' && shapes[2] == 'Z' {      //losing
            total_score = total_score + 3 + 6
         } else if shapes[0] == 'C' && shapes[2] == 'X' {      //tie
            total_score = total_score + 2 + 0
         } else if shapes[0] == 'C' && shapes[2] == 'Y' {      //tie
            total_score = total_score + 3 + 3
         } else if shapes[0] == 'C' && shapes[2] == 'Z' {      //tie
            total_score = total_score + 1 + 6
         }
      }
   }

   fmt.Printf("Total score is: %d\n", total_score)
} //function-main
