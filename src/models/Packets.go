package models

type NextTurnPacket struct {
	Id           string
	Turn         int
	DoublesCount int
	HasRolled    bool
}

type RollResponsePacket struct {
	Id    string
	Dice1 int
	Dice2 int
}

type AddCashPacket struct {
	Id       string
	PlayerId int
	Cash     int
}

type SetCashPacket struct {
	Id       string
	PlayerId int
	Cash     int
}

type SetDoublesCount struct {
	Id           string
	DoublesCount int
}

type SetPlayerPositionPacket struct {
	Id       string
	Position int
}

type SetHasRolledPacket struct {
	Id        string
	HasRolled bool
}

type StatePacket struct {
	Id    string
	Board Board
}

type BuyCostPacket struct {
	Id   string
	Cost int
}

type PropertyOwnerPacket struct {
	Id         string
	PropertyId int
	PlayerId   int
}

type BuySpacePacket struct {
	Id         string
	PropertyId int
}
