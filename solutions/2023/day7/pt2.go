package day7

import (
	"advent-of-go/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var ranksWithJokers = []rune{'J','2','3','4','5','6','7','8','9','T','Q','K','A'}

func Pt2() utils.Solution {
	return utils.Solution{
		Year: 2023,
		Day: 7,
		Part: 2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := utils.GetLines(input)
	hands := make([]hand, len(lines))
	for i, line := range lines {
		hand, e := getHandWithJokers(line)
		if e != nil {
			return "", fmt.Errorf("Malformed input: something happened when parsing hand from line %s: %w", line, e)
		}
		hands[i] = hand
	}
	slices.SortFunc(hands, sortHandsWithJokers)
	totalWinnings := 0
	for i, hand_ := range hands {
		totalWinnings += (i + 1) * hand_.bid
	}
	return strconv.Itoa(totalWinnings), nil
}

func sortHandsWithJokers(hand1, hand2 hand) int {
	if hand1.handType != hand2.handType {
		return hand1.handType - hand2.handType
	}
	for i := range hand1.cards {
		hand1Card, hand2Card := rune(hand1.cards[i]), rune(hand2.cards[i])
		h1Idx, h2Idx := slices.Index(ranksWithJokers, hand1Card), slices.Index(ranksWithJokers, hand2Card)
		if h1Idx != h2Idx {
			return h1Idx - h2Idx
		}
	}
	return 0
}

func getHandWithJokers(line string) (hand, error) {
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
	return hand{cards, getHandTypeWithJokers(cards), bid}, nil
}

func getHandTypeWithJokers(cards string) HandType {
	cardCounts := getCardCounts(cards)
	preprocessJokers(cardCounts)
	return getHandTypeFromCardCounts(cardCounts)
}

func preprocessJokers(cardCounts map[rune]int) {
	var maxRune rune
	var maxVal int
	var numJ int
	for k, v := range cardCounts {
		if k == 'J' {
			numJ += v
		} else if v > maxVal {
			maxRune = k
			maxVal = v
		}
	}
	if maxVal == 0 {
		return
	}
	cardCounts[maxRune] += numJ
	delete(cardCounts, 'J')
}