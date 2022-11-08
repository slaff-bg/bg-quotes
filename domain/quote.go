package domain

type Quote struct {
	Id          int
	Quote       string
	SmokingRoom bool
	Author      Author
}
