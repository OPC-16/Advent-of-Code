package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	//Read input file
	input, _ := os.Open("puzzle_input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var sumOfPriorities int

	for scanner.Scan() {
		// Create three sets with the elements of each elf
		itemsFirst := createMapOfItems(scanner.Text())

		scanner.Scan()
		itemsSecond := createMapOfItems(scanner.Text())

		scanner.Scan()
		itemsThird := createMapOfItems(scanner.Text())

		// For each item of the first (or second or third, it's the same) elf we check if the other two elves have that item too
		for itemFirstElf := range itemsFirst {
			if itemsSecond[itemFirstElf] && itemsThird[itemFirstElf] {
				sumOfPriorities += int(unicode.ToLower(itemFirstElf)-96)
				if unicode.IsUpper(itemFirstElf){
					sumOfPriorities += 26
				}
				break
			}
		}
	}
	fmt.Println(sumOfPriorities)
}

func createMapOfItems(items string) (map[rune]bool) {
    set := make(map[rune]bool)
	for _, item := range items{
		set[item] = true
	}
	return set
}
