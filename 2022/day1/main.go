package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
   var sumOfCalories = 0
   var currMaxCalorie = 0
   maxCalories := make([]int, 1)

   file, err := os.Open("puzzle_input.txt")
   if err != nil {
      return
   }
   defer file.Close()

   // Create a scanner to read lines from the file
   scanner := bufio.NewScanner(file)

   // Iterate through the lines and process them one by one
   for scanner.Scan() {
      line := scanner.Text()

      if line != "" {
         num, _ := strconv.Atoi(line)
         sumOfCalories += num
      } else {
         currMaxCalorie = max(sumOfCalories, currMaxCalorie)
         maxCalories = append(maxCalories, sumOfCalories)
         sumOfCalories = 0
      }
   }

   if err := scanner.Err(); err != nil {
      fmt.Println(err)
   }

   fmt.Printf("Max calories (for part-one of puzzle) are: %d\n", currMaxCalorie)

   //----- for part-two of puzzle ----
   sumOfCalories = currMaxCalorie
   var index = 0
   var val = 0

   // remove the currMaxCalorie element from the maxCalories slice
   for index, val = range maxCalories {
      if val == currMaxCalorie {
         maxCalories = append(maxCalories[:index], maxCalories[index+1:]...)
      } else {
         continue
      }
   }

   currMaxCalorie = maxCalories[0]
   for index, val = range maxCalories {
      if val > currMaxCalorie {
         currMaxCalorie = val
      } else {
         continue
      }
   }
   sumOfCalories += currMaxCalorie
   for index, val = range maxCalories {
      if val == currMaxCalorie {
         maxCalories = append(maxCalories[:index], maxCalories[index+1:]...)
      } else {
         continue
      }
   }
   currMaxCalorie = maxCalories[0]
   for index, val = range maxCalories {
      if val > currMaxCalorie {
         currMaxCalorie = val
      } else {
         continue
      }
   }
   sumOfCalories += currMaxCalorie
   fmt.Printf("ans for part-two: %d\n", sumOfCalories)
} //function-main
