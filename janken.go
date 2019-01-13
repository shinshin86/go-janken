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

	fmt.Println(handA.GetHand())
	fmt.Println(handB.GetHand())
	fmt.Println(hand.Battle(handA, handB))
	fmt.Println(handA)
	fmt.Println(handB)
}
