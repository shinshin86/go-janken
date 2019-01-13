package hand

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Hand struct {
	Name     string
	winCount int
	handNum  int
}

var handMap = map[int]string{
	0: "グー",
	1: "チョキ",
	2: "パー",
}

// Constructor
func NewHand(name string) (*Hand, error) {
	if name == "" {
		return nil, errors.New("Required hand name")
	}

	handInt := selectHand()
	return &Hand{
		Name:     name,
		handNum:  handInt,
		winCount: 0,
	}, nil
}

// update hand method(public function)
func (hand Hand) UpdateHand() Hand {
	hand.handNum = selectHand()
	return hand
}

// get handNum function(private function)
func selectHand() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

// get hand function (public function)
func (hand Hand) GetHand() string {
	return handMap[hand.handNum]
}

// Battle function (public function)
func Battle(handA, handB *Hand) string {
	if (handA.handNum == 0 && handB.handNum == 1) || (handA.handNum == 1 && handB.handNum == 2) || handA.handNum == 2 && handB.handNum == 0 {
		handA.winCount++
		return fmt.Sprintf("%sの勝ち", handA.Name)
	} else if (handB.handNum == 0 && handA.handNum == 1) || (handB.handNum == 1 && handA.handNum == 2) || handB.handNum == 2 && handA.handNum == 0 {
		handB.winCount++
		return fmt.Sprintf("%sの勝ち", handB.Name)
	} else {
		return fmt.Sprintf("引き分け")
	}
}
