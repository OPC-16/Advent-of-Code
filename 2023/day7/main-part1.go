package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "slices"
)

type rankAndBid struct {
    rank int
    bid  int
}

// this maps strength to card (numbers doesn't matter only comparision matters)
var cardToStrengthMap = make(map[byte]int)

func main() {
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    //map of hand 'string' to rankAndBid 'struct' 
    Map := make(map[string]rankAndBid)

    // this maps strength to card (numbers doesn't matter only comparision matters)
    cardToStrengthMap['A'] = 100
    cardToStrengthMap['K'] = 80
    cardToStrengthMap['Q'] = 50
    cardToStrengthMap['J'] = 30
    cardToStrengthMap['T'] = 10

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

        //put the current hand in appropriate i 'slice'
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

    //Sort the individual types 'slices', each in ascending order
    slices.SortFunc(fiveKind, sortFiveKind)
    slices.SortFunc(fourKind, sortRest)
    slices.SortFunc(fullHouse, sortRest)
    slices.SortFunc(threeKind, sortRest)
    slices.SortFunc(twoPair, sortRest)
    slices.SortFunc(onePair, sortRest)
    slices.SortFunc(highCard, sortRest)

    //Now increment the ranks from weakest to strongest i
    rankCounter := 1
    for i:=0; i<len(highCard); i++ {
        val, _ := Map[highCard[i]]
        val.rank = rankCounter
        Map[highCard[i]] = val
        rankCounter++
    }
    for i:=0; i<len(onePair); i++ {
        val, _ := Map[onePair[i]]
        val.rank = rankCounter
        Map[onePair[i]] = val
        rankCounter++
    }
    for i:=0; i<len(twoPair); i++ {
        val, _ := Map[twoPair[i]]
        val.rank = rankCounter
        Map[twoPair[i]] = val
        rankCounter++
    }
    for i:=0; i<len(threeKind); i++ {
        val, _ := Map[threeKind[i]]
        val.rank = rankCounter
        Map[threeKind[i]] = val
        rankCounter++
    }
    for i:=0; i<len(fullHouse); i++ {
        val, _ := Map[fullHouse[i]]
        val.rank = rankCounter
        Map[fullHouse[i]] = val
        rankCounter++
    }
    for i:=0; i<len(fourKind); i++ {
        val, _ := Map[fourKind[i]]
        val.rank = rankCounter
        Map[fourKind[i]] = val
        rankCounter++
    }
    for i:=0; i<len(fiveKind); i++ {
        val, _ := Map[fiveKind[i]]
        val.rank = rankCounter
        Map[fiveKind[i]] = val
        rankCounter++
    }

    fmt.Println("Map:" , Map)

    ans := 0
    for _, value := range Map {
        ans += value.bid * value.rank
    }
    fmt.Println("Final ans: ", ans)
} //main

func sortFiveKind(a, b string) int {
    firstCharOfa := a[0]
    firstCharOfb := b[0]

    //if only numbers are present
    if firstCharOfa >= '2' && firstCharOfa <= '9' || firstCharOfb >= '2' && firstCharOfb <= '9' {
        if firstCharOfa < firstCharOfb {
            return -1
        } else if firstCharOfa > firstCharOfb {
            return 1
        } else {
            return 0
        }
    } else { //means that it is alphabets
        if cardToStrengthMap[firstCharOfa] < cardToStrengthMap[firstCharOfb] {
            return -1
        } else if cardToStrengthMap[firstCharOfa] > cardToStrengthMap[firstCharOfb] {
            return 1
        } else {
            return 0
        }
    }
}

func sortRest(a, b string) int {
    for i:=0; i<5; i++ {
        if a[i] >= '2' && a[i] <= '9' || b[i] >= '2' && b[i] <= '9' {
            if a[i] < b[i] {
                return -1
            } else if a[i] > b[i] {
                return 1
            } else {
                continue
            }
        } else {
            if cardToStrengthMap[a[i]] < cardToStrengthMap[b[i]] {
                return -1
            } else if cardToStrengthMap[a[i]] > cardToStrengthMap[b[i]] {
                return 1
            } else {
                continue
            }
        }

    }
    return 0
}
