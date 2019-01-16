package main

import (
	"fmt"
	"os"

	"github.com/shinshin86/go-janken/hand"
)

func main() {
	fmt.Println("Let's janken...")

	handA, err := hand.NewHand("User A")
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	handB, err := hand.NewHand("User B")
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

	battleCount := 0
	for battleCount < 10 {
		battleCount++
		fmt.Println(battleCount)
		fmt.Println(handA.UpdateHand())
		fmt.Println(handB.UpdateHand())
		fmt.Println(handA.GetHand())
		fmt.Println(handB.GetHand())
		fmt.Println(hand.Battle(handA, handB))
	}
}
