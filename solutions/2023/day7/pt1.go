package day7

import (
	"advent-of-go/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type HandType = int

const (
	HighCard HandType = iota
	Pair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var ranks = []rune{'2','3','4','5','6','7','8','9','T','J','Q','K','A'}

func Pt1() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 7,
		Part: 1,
		Calculator: pt1,
	}
}

type hand struct {
	cards string
	handType HandType
	bid int
}

func pt1(input string) (string, error) {
	lines := utils.GetLines(input)
	hands := make([]hand, len(lines))
	for i, line := range lines {
		hand, e := getHand(line)
		if e != nil {
			return "", fmt.Errorf("Malformed input: something happened when parsing hand from line %s: %w", line, e)
		}
		hands[i] = hand
	}
	slices.SortFunc(hands, sortHands)
	totalWinnings := 0
	for i, hand_ := range hands {
		totalWinnings += (i + 1) * hand_.bid
	}
	return strconv.Itoa(totalWinnings), nil
}

func sortHands(hand1, hand2 hand) int {
	if hand1.handType != hand2.handType {
		return hand1.handType - hand2.handType
	}
	for i := range hand1.cards {
		hand1Card, hand2Card := rune(hand1.cards[i]), rune(hand2.cards[i])
		h1Idx, h2Idx := slices.Index(ranks, hand1Card), slices.Index(ranks, hand2Card)
		if h1Idx != h2Idx {
			return h1Idx - h2Idx
		}
	}
	return 0
}

func getHand(line string) (hand, error) {
	toks := strings.Split(line, " ")
	if len(toks) != 2 {
		return hand{}, fmt.Errorf("Malformed data: line %s doesn't have the correct number of tokens", line)
	}
	cards, bidStr := toks[0], toks[1]
	if len(cards) != 5 {
		return hand{}, fmt.Errorf("Malformed data: hand %s should have 5 cards!", cards)
	}
	for _, card := range cards {
		if slices.Index(ranks, card) == -1 {
			return hand{}, fmt.Errorf("Malformed data: Cards %s have illegal characters", cards)
		}
	}
	bid, e := strconv.Atoi(bidStr)
	if e != nil {
		return hand{}, fmt.Errorf("Malformed data: bid string wasn't numeric: %w", e)
	}
	return hand{cards, getHandType(cards), bid}, nil
}

func getHandType(cards string) HandType {
	cardCounts := getCardCounts(cards)
	return getHandTypeFromCardCounts(cardCounts)
}

func getCardCounts(cards string) map[rune]int {
	cardCounts := map[rune]int{}
	for _, r := range cards {
		if _, ok := cardCounts[r]; ok {
			cardCounts[r] += 1
		} else {
			cardCounts[r] = 1
		}
	}
	return cardCounts
}

func getHandTypeFromCardCounts(cardCounts map[rune]int) HandType {
	if getNumEntriesWithValue(cardCounts, 5) == 1 {
		return FiveOfAKind
	}
	if getNumEntriesWithValue(cardCounts, 4) == 1 {
		return FourOfAKind
	}
	numPairs := getNumEntriesWithValue(cardCounts, 2)
	if getNumEntriesWithValue(cardCounts, 3) == 1 {
		if numPairs == 1 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	} 
	if numPairs == 2 {
		return TwoPair
	}
	if numPairs == 1 {
		return Pair
	} 
	return HighCard
}

func getNumEntriesWithValue[K, V comparable](toCheck map[K]V, value V) int {
	count := 0
	for _, v := range toCheck {
		if v == value {
			count ++
		}
	}
	return count
}