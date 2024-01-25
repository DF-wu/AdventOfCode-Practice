/*
Day 7: The One-Handed Tyrant
*/

package main

import (
	"fmt"
	"github.com/DF-wu/AdventOfCode-Practice/dfmisc"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := dfmisc.Filereader("input.txt")
	//fmt.Println(p1(input))
	fmt.Println(p2(input))
	//fmt.Println(input)
}

/*
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
*/

// 1. Five of a kind, where all five cards have the same label: AAAAA
// 2. Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// 3. Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// 4. Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// 5. Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// 6. One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// 7. High card, where all cards' labels are distinct: 23456

var FIVEOFAKIND, FOUROFAKIND, FULLHOUSE, THREEOFAKIND, TWOPAIR, ONEPAIR, HIGHCARD = 7, 6, 5, 4, 3, 2, 1

type HandCard struct {
	hand     string
	handType int
	bid      int
	rank     int
}

/*
implementation of sort interface
https://easonwang.gitbook.io/golang/he-xin-mo-zu/sort
*/
type sortByHand []HandCard

func (a sortByHand) Len() int {
	return len(a)
}
func (a sortByHand) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a sortByHand) Less(i, j int) bool {
	// determine the strength of each char
	// p1 without joker type
	//strengthDict := map[string]int{"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}
	//with  joker tpye
	strengthDict := map[string]int{"A": 14, "K": 13, "Q": 12, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 1}

	//  compare the hand type first !!
	if a[i].handType < a[j].handType {
		return true
	} else if a[i].handType > a[j].handType {
		return false
	} else {
		// if they are the same hand type, compare the strength
		// use each char(string type) as element to find corresponding strength in strengthDict. and compare
		// if a[i] compare to a[j] by each char
		for pos, _ := range a[i].hand {

			if strengthDict[string(a[i].hand[pos])] == strengthDict[string(a[j].hand[pos])] {
				// same char -> continue
				continue
			} else if strengthDict[string(a[i].hand[pos])] < strengthDict[string(a[j].hand[pos])] {
				return true
			} else {
				return false
			}
		}
	}

	// only happen in the case of same hand
	// to maintain the original order -> return false
	return false
}

func p2(input string) int {
	//modeling the input
	HandCards := []HandCard{}
	for _, line := range strings.Split(input, "\n") {
		hand, bid := "", 0
		fmt.Sscanf(line, "%s %d", &hand, &bid)
		HandCards = append(HandCards, HandCard{
			hand:     hand,
			handType: typeIndicatorWithWildCard(hand),
			bid:      bid,
			rank:     0,
		})

	}

	/*
		!!!!!!!!!!!!! MY SUPER MAGIC SORT !!!!!!!!!!!!!

		order from low to high
	*/
	sort.Sort(sortByHand(HandCards))
	// update  each handcard rank
	rankCtr := 1
	for pos, _ := range HandCards {
		HandCards[pos].rank = rankCtr
		rankCtr++
	}

	fmt.Println(HandCards)
	// to get the answer, multiply the bid and rank
	ans := 0
	for _, handCard := range HandCards {
		ans = ans + handCard.bid*handCard.rank
	}
	return ans

}

func p1(input string) int {
	//modeling the input
	HandCards := []HandCard{}
	for _, line := range strings.Split(input, "\n") {
		hand, bid := "", 0
		fmt.Sscanf(line, "%s %d", &hand, &bid)
		HandCards = append(HandCards, HandCard{
			hand:     hand,
			handType: typeIndicator(hand),
			bid:      bid,
			rank:     0,
		})

	}
	//fmt.Println(HandCards)

	/*
		MY SUPER MAGIC SORT !!!!!!!!!!!!!

		order from low to high
	*/
	sort.Sort(sortByHand(HandCards))
	// update  each handcard rank
	rankCtr := 1
	for pos, _ := range HandCards {
		HandCards[pos].rank = rankCtr
		rankCtr++
	}

	// because i use my magic sort. below is not needed
	//
	//// count from  FIVEOFAKIND to HIGHCARD
	//for currHandTpye := 6; currHandTpye < 0; currHandTpye-- {
	//	selectTypeHandCards := []HandCard{}
	//	for _, handCard := range HandCards {
	//		if handCard.handType == currHandTpye {
	//			selectTypeHandCards = append(selectTypeHandCards, handCard)
	//		}
	//	}
	//
	//	if len(selectTypeHandCards) == 1 {
	//		// only one -> assign the rank and append to sortedHandCards
	//		selectTypeHandCards[0].rank = rankCtr
	//		rankCtr--
	//		sortedHandCards = append(sortedHandCards, selectTypeHandCards[0])
	//	} else if len(selectTypeHandCards) == 0 {
	//		// no hand card of this type
	//	} else {
	//		// more than one hand card of this type
	//		// compare the strength
	//		for _, hand := range selectTypeHandCards {
	//
	//		}
	//	}
	//}

	fmt.Println(HandCards)
	// to get the answer, multiply the bid and rank
	ans := 0
	for _, handCard := range HandCards {
		ans = ans + handCard.bid*handCard.rank
	}
	return ans

}

func typeIndicatorWithWildCard(hand string) int {
	// p2 with joker
	strengthDict := map[string]int{"A": 14, "K": 13, "Q": 12, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 1}
	dict := map[string]int{}
	for _, str := range hand {
		//cal freq
		dict[string(str)]++
	}

	// check if contains Joker
	_, isExist := dict["J"]
	//fmt.Println(indicator)
	if !isExist {
		// no joker
		// follow the same logic as before
		return TypeMapper(dict)

	} else {
		// contains joker
		jokerCtr := dict["J"]
		delete(dict, "J")

		//Joker always pretent to be the largest card
		// find largest value slice. ex: 3 3 3 2 2 -> 3 3 3
		//if there is over 1 max value(same value as max), find the largest card

		// ↓ cause error. Because extra space is added
		//sortedValues := make([]int, len(dict))
		sortedValues := []int{}
		for _, freq := range dict {
			sortedValues = append(sortedValues, freq)
		}
		// from low to high
		sort.Ints(sortedValues)
		//init. largest number at least 1
		if len(sortedValues) == 0 {
			// JJJJJ
			return FIVEOFAKIND
		} else {
			pivot := len(sortedValues) - 1
			lastNum := sortedValues[pivot]

			for i := len(sortedValues) - 1; i >= 0; i-- {
				curr := sortedValues[i]
				if curr == lastNum {
					// not changed, keep this element
				} else {
					break
				}
				pivot--
			}
			// compasate the pivot
			pivot++
			sliceList := sortedValues[pivot:]
			// may be 1 or more
			// if it is 1, then it is the largest value
			if len(sliceList) == 1 {
				for k, v := range dict {
					if sliceList[0] == v {
						dict[k] = v + jokerCtr
					}
				}
			} else {
				// more than 1
				// find the largest card
				target := sliceList[0] // they should be the same

				//pick a key in dict
				keys := make([]string, len(dict))
				for k := range dict {
					keys = append(keys, k)
				}

				// init
				card := keys[0]
				for k, v := range dict {
					if target == v {
						if strengthDict[k] > strengthDict[card] {
							card = k
						}

					}
				}
				dict[card] = target + jokerCtr

			}
		}

		return TypeMapper(dict)

	}
}

func TypeMapper(dict map[string]int) int {
	sortedValues := []int{}
	for _, freq := range dict {
		sortedValues = append(sortedValues, freq)
	}
	sort.Ints(sortedValues)
	indicator := ""
	// sortd as low to high.  But I like high to low
	for _, value := range sortedValues {
		indicator = strconv.Itoa(value) + indicator
	}
	// check type
	switch indicator {
	case "5":
		return FIVEOFAKIND
	case "41":
		return FOUROFAKIND
	case "32":
		return FULLHOUSE
	case "311":
		return THREEOFAKIND
	case "221":
		return TWOPAIR
	case "2111":
		return ONEPAIR
	case "11111":
		return HIGHCARD
	default:
		//should not be here
	}
	return 0
}

// to determine the hands card type
func typeIndicator(hand string) int {
	dict := map[string]int{}
	for _, str := range hand {
		//cal freq
		dict[string(str)]++
	}
	// each type have a unique indicator to determine the type
	// The sorted dict value list is the key to determine the type
	// 1. Five of a kind, where all five cards have the same label: AAAAA -> 5
	// 2. Four of a kind, where four cards have the same label and one card has a different label: AA8AA -> 4 1
	// 3. Full house, where three cards have the same label, and the remaining two cards share a different label: 23332 -> 3 2
	// 4. Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98 -> 3 1 1
	// 5. Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432 -> 2 2 1
	// 6. One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4 -> 2 1 1 1
	// 7. High card, where all cards' labels are distinct: 23456 -> 1 1 1 1 1
	sortedValues := []int{}
	for _, freq := range dict {
		sortedValues = append(sortedValues, freq)
	}
	sort.Ints(sortedValues)
	indicator := ""
	// sortd as low to high.  But I like high to low
	for _, value := range sortedValues {
		indicator = strconv.Itoa(value) + indicator
	}
	//fmt.Println(indicator)

	switch indicator {
	case "5":
		return FIVEOFAKIND
	case "41":
		return FOUROFAKIND
	case "32":
		return FULLHOUSE
	case "311":
		return THREEOFAKIND
	case "221":
		return TWOPAIR
	case "2111":
		return ONEPAIR
	case "11111":
		return HIGHCARD
	default:
		//should not be here
	}
	return 0
	/* Belowd code didn't consider Fullhouse. So stupid*/
	/*
		maxCardFreq := 0
		//two pair -> 4
		// one pair -> 2
		// high card -> 0
		pairCheckSum := 0
		for _, freq := range dict {
			maxCardFreq = max(maxCardFreq, freq)
			if freq == 2 {
				pairCheckSum += freq
			}
		}
		// first check if it is a straight
		switch maxCardFreq {
		case 5:
			return FIVEOFAKIND
		case 4:
			return FOUROFAKIND
		case 3:
			return THREEOFAKIND
		default:
			//go to next check
		}

		// if not straight, check if it is a kind of pair
		switch pairCheckSum {
		case 4:
			return TWOPAIR
		case 2:
			return ONEPAIR
		case 0:
			return HIGHCARD
		default:
			//should not be here
			fmt.Println("error")
			return 0
		}
	*/
}
