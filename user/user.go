package user

import (
	"errors"
)

type User struct {
	Name      string
	winCount  int
	loseCount int
	drawCount int
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("Required user name")
	}

	return &User{
		Name:      name,
		winCount:  0,
		loseCount: 0,
		drawCount: 0,
	}, nil
}

func (user *User) Win() (*User, error) {
	if user == nil {
		return nil, errors.New("Not found user")
	}
	user.winCount++
	return user, nil
}

func (user *User) Lose() (*User, error) {
	if user == nil {
		return nil, errors.New("Not found user")
	}
	user.loseCount++
	return user, nil
}

func (user *User) Draw() (*User, error) {
	if user == nil {
		return nil, errors.New("Not found user")
	}
	user.drawCount++
	return user, nil
}

func GetBattleCount(user *User) (int, error) {
	if user == nil {
		return 0, errors.New("Not found user")
	}

	return user.winCount + user.loseCount + user.drawCount, nil
}
