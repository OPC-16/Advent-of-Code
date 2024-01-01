package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rankAndBid struct {
    rank int
    bid  int
}

func main() {
    file, _ := os.Open("test_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    //map of hand 'string' to rankAndBid 'struct' 
    Map := make(map[string]rankAndBid)

    //slices for the types of hands from strongest to weakest
    var fiveKind []string
    var fourKind []string
    var fullHouse []string
    var threeKind []string
    var twoPair []string
    var onePair []string
    var highCard []string

    for scn.Scan() {
        line := scn.Text()

        handBidSlice := strings.Split(line, " ")
        bidInt, _ := strconv.Atoi(handBidSlice[1])
        fmt.Printf("hand: %s  Bid: %d\n", handBidSlice[0], bidInt)

        //putting the hand and their bids in Map with -1 rank
        Map[handBidSlice[0]] = rankAndBid{ rank: -1, bid: bidInt, }

        //put the current hand in appropriate type 'slice'
        sliceOfCards := strings.Split(handBidSlice[0], "")
        handMap := make(map[string]int)
        for i := 0; i<len(sliceOfCards); i++ {
            if _, ok := handMap[sliceOfCards[i]]; ok {
                handMap[sliceOfCards[i]]++
            } else {
                handMap[sliceOfCards[i]] = 1
            }
        }
        if len(handMap) == 1 {      //it is fiveKind
            fiveKind = append(fiveKind, handBidSlice[0])
        } else if len(handMap) == 2 {   //might be fourKind or fullHouse
            for _, value := range handMap {
                if value == 4 {
                    fourKind = append(fourKind, handBidSlice[0])
                    break
                }
                if value == 3 {
                    fullHouse = append(fullHouse, handBidSlice[0])
                    break
                }
            }
        } else if len(handMap) == 3 {   //might be 3Kind or 2pair
            for _, value := range handMap {
                if value == 3 {
                    threeKind = append(threeKind, handBidSlice[0])
                    break
                }
                if value == 2 {
                    twoPair = append(twoPair, handBidSlice[0])
                    break
                }
            }

        } else if len(handMap) == 4 {  //is onepair
            onePair = append(onePair, handBidSlice[0])
        } else {                        //is highCard
            highCard = append(highCard, handBidSlice[0])
        }
    }

    //Sort the individual types 'slices'

    fmt.Println("Map:" , Map)
}
