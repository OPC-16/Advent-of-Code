// actually this is 2015's day2 challenge ;)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
   totalWrappingPaper := 0

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

      //split the string using 'x' as delimiter
      parts := strings.Split(line, "x")

      if len(parts) >= 3 {
         length, _ := strconv.Atoi(parts[0])
         fmt.Println(length)
         width, _ := strconv.Atoi(parts[1])
         fmt.Println(width)
         height, _ := strconv.Atoi(parts[2])
         fmt.Println(height)

         wrappingPaper := 2*length*width + 2*width*height + 2*length*height
         min1, min2 := find2min(length, width, height)
         slack := min1*min2

         totalWrappingPaper = totalWrappingPaper + wrappingPaper + slack
      }
   }

   fmt.Printf("Total wrapping required is: %d\n", totalWrappingPaper)
} //function-main

func find2min(a, b, c int) (int, int){
   var min1, min2 int

   min1 = min(a, b)
   if min1 == a {
      min2 = min(b, c)
   } else {
      min2 = min(a, c)
   }

   return min1, min2
} //function-find2min
