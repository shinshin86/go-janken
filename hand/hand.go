package hand

import (
	"errors"
	"fmt"
	"github.com/shinshin86/go-janken/user"
	"math/rand"
	"time"
)

type Hand struct {
	user    *user.User
	handNum int
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

	user, err := user.NewUser(name)

	if err != nil {
		return nil, errors.New("Create user error")
	}

	handInt := selectHand()
	return &Hand{
		user:    user,
		handNum: handInt,
	}, nil
}

// update hand method(public function)
func (hand *Hand) UpdateHand() *Hand {
	hand.handNum = selectHand()
	return hand
}

// get user method (public function)
func (hand *Hand) GetUser() *user.User {
	return hand.user
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
		handA.user.Win()
		handB.user.Lose()
		return fmt.Sprintf("%sの勝ち", handA.user.Name)
	} else if (handB.handNum == 0 && handA.handNum == 1) || (handB.handNum == 1 && handA.handNum == 2) || handB.handNum == 2 && handA.handNum == 0 {
		handA.user.Lose()
		handB.user.Win()
		return fmt.Sprintf("%sの勝ち", handB.user.Name)
	} else {
		handA.user.Draw()
		handB.user.Draw()

		return fmt.Sprintf("引き分け")
	}
}
