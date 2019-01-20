package hand

import (
	"testing"

	"github.com/shinshin86/go-janken/hand"
)

func TestNewHandSuccess(t *testing.T) {
	name := "test name"
	hand, err := hand.NewHand(name)
	if err != nil {
		t.Errorf("error")
	}

	if name != hand.GetUser().Name {
		t.Errorf("got %v\nwant %v", hand.GetUser().Name, name)
	}
	t.Log(hand)
}

func TestNewHandFailed(t *testing.T) {
	_, err := hand.NewHand("")
	if err != nil {
		t.Log(err)
		return
	}
	t.Error("got not error\nwant required name error")
}

func TestGetHandSuccess(t *testing.T) {
	hands := []string{"グー", "チョキ", "パー"}
	testHand, err := hand.NewHand("test name")

	if err != nil {
		t.Errorf("New Hand error : %v", err)
	}

	if arrayContains(hands, testHand.GetHand()) == false {
		t.Error("got not has hands\nwant hands")
	}
	t.Log(testHand)

}

// Check function of has str at arr
func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
