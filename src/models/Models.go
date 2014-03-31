package models

type Player struct {
	IsHuman           bool
	Cash              int
	Order             int
	Properties        []Property
	Token             Tokens
	Position          int
	JailedTurn        int
	HasJailFreeChance bool
	HasJailFreeChest  bool
}

type Tokens int

const (
	Battleship = iota
	Dog
	Iron
	RaceCar
	Shoe
	Thimble
	TopHat
	Wheelbarrow
)

type Colors int

const (
	Brown = iota
	Sky
	Pink
	Orange
	Red
	Yellow
	Green
	Blue
)

type Space struct {
	Name     string
	Position int
}

type OwnableSpace struct {
	Space
	Owner          *Player
	IsMonopoly     bool
	MortgageRate   int
	UnmortgageRate int
	Cost           int
}

type Utility struct {
	OwnableSpace
}

type IncomeTax struct {
	Space
}

type LuxaryTax struct {
	Space
}

type GoToJail struct {
	Space
}

type RailRoad struct {
	OwnableSpace
}

type Property struct {
	OwnableSpace
	Color       Colors
	Upgrades    int
	Rent        []int
	UpgradeCost int
}

type CommunityChest struct {
	Space
}

type Chance struct {
	Space
}

type CommunityChestCard struct {
	Id          int
	Description string
}

type ChanceCard struct {
	Id          int
	Description string
}

type Board struct {
	Turn                int
	ChanceCards         []ChanceCard         `json:"-"`
	CommunityChestCards []CommunityChestCard `json:"-"`
	Players             []Player
	DoublesCount        int
	HasRolled           bool
	Spaces              []HandleSpacer
}

func (b *Board) NextTurn() {
	b.Turn++
	b.Turn %= len(b.Players)
	b.HasRolled = false
	b.DoublesCount = 0
}

func (b *Board) GetCurrentPlayer() *Player {
	return &b.Players[b.Turn]
}

func (b *Board) GetCurrentSpace() *HandleSpacer {
	return &b.Spaces[b.GetCurrentPlayer().Position]
}

type HandleSpacer interface {
	HandleSpace(*Board) []interface{}
}
