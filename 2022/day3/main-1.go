// for a-z, minus 96 from their ASCII value to get priority
// for A-Z, minus 38 from their ASCII value to get priority
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
   sumOfPriorities := 0

   file, err := os.Open("puzzle_input.txt")
   if err != nil {
      fmt.Println(err)
      return
   }
   defer file.Close()

   scanner := bufio.NewScanner(file)

   for scanner.Scan() {
      line := scanner.Text()

      midPoint := len(line) / 2

      // splitting the rucksack into 2 compartments
      firstCompartment := []rune(line[:midPoint])
      secondCompartment := []rune(line[midPoint:])

      // making a map of only first compartment
      runeMap := make(map[rune] int)

      // putting unique values in map from first compartment slice
      for i := 0; i < len(firstCompartment); i++ {
         if _, exists := runeMap[firstCompartment[i]]; !exists {
            runeMap[firstCompartment[i]] = 0
         }
      }

      for i := 0; i < midPoint; i++ {
         for j := 0; j < midPoint; j++ {
            if firstCompartment[i] == secondCompartment[j] && runeMap[firstCompartment[i]] == 0 {
               if unicode.IsUpper(firstCompartment[i]) {
                  sumOfPriorities = sumOfPriorities + int(firstCompartment[i]) - 38
               } else {
                  sumOfPriorities = sumOfPriorities + int(firstCompartment[i]) - 96
               }
               runeMap[firstCompartment[i]] = 1
               break
            }
         }
      }
   }

   fmt.Printf("Total sum: %d\n", sumOfPriorities)
} //function-main
