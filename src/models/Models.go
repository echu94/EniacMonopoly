package models

import (
	"math/rand"
)

type Room struct {
	Board   Board
	Clients []chan []interface{}
}

func (r *Room) Send(packets []interface{}) {
	for i := 0; i < len(r.Clients); i++ {
		r.Clients[i] <- packets
	}
}

type Player struct {
	IsHuman           bool
	Cash              int
	Order             int
	OwnableSpaces     []*HandleSpacer `json:"-"`
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
	Rent        [6]int
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
	BuyCost             int
	Started             bool
	//ElapsedTurns        int
}

func (b *Board) NextTurn() {
	b.Turn++
	b.Turn %= len(b.Players)
	b.HasRolled = false
	b.DoublesCount = 0
	//b.ElapsedTurns++
}

func (b *Board) GetCurrentPlayer() *Player {
	return &b.Players[b.Turn]
}

func (b *Board) GetCurrentSpace() *HandleSpacer {
	return &b.Spaces[b.GetCurrentPlayer().Position]
}

func (b *Board) AddPlayer() *Player {
	i := len(b.Players)
	p := Player{
		IsHuman:    true,
		Cash:       1500,
		Order:      i,
		Token:      Tokens(i),
		JailedTurn: -1,
	}
	b.Players = append(b.Players, p)
	return &p
}

type HandleSpacer interface {
	HandleSpace(*Board) []interface{}
}

func (b *Board) Initialize() {
	b.Players = make([]Player, 0)

	b.initializeSpaces()

	b.initializeCards()
}
func (b *Board) initializeCards() {
	var chanceCards [16]ChanceCard
	chanceCards[0] = ChanceCard{Id: 0, Description: "THIS CARD MAY BE KEPT UNTIL NEEDED OR SOLD - GET OUT OF JAIL FREE"}
	chanceCards[1] = ChanceCard{Id: 1, Description: "GO DIRECTLY TO JAIL - DO NOT PASS GO, DO NOT COLLECT $200"}
	chanceCards[2] = ChanceCard{Id: 2, Description: "YOU HAVE BEEN ELECTED CHAIRMAN OF THE BOARD - PAY EACH PLAYER $50"}
	chanceCards[3] = ChanceCard{Id: 3, Description: "ADVANCE TOKEN TO THE NEAREST RAILROAD AND PAY OWNER TWICE THE RENTAL TO WHICH HE IS OTHERWISE ENTITLED. - IF RAILROAD IS UNOWNED, YOU MAY BUY IT FROM THE BANK"}
	chanceCards[4] = ChanceCard{Id: 4, Description: "ADVANCE TO ILLINOIS AVE."}
	chanceCards[5] = ChanceCard{Id: 5, Description: "ADVANCE TO GO - (COLLECT $200)"}
	chanceCards[6] = ChanceCard{Id: 6, Description: "GO BACK 3 SPACES"}
	chanceCards[7] = ChanceCard{Id: 7, Description: "PAY POOR TAX OF $15"}
	chanceCards[8] = ChanceCard{Id: 8, Description: "ADVANCE TOKEN TO THE NEAREST RAILROAD AND PAY OWNER TWICE THE RENTAL TO WHICH HE IS OTHERWISE ENTITLED. - IF RAILROAD IS UNOWNED, YOU MAY BUY IT FROM THE BANK"}
	chanceCards[9] = ChanceCard{Id: 9, Description: "TAKE A WALK ON THE BOARD WALK - ADVANCE TOKEN TO BOARD WALK"}
	chanceCards[10] = ChanceCard{Id: 10, Description: "BANK PAYS YOU DIVIDEND OF $50"}
	chanceCards[11] = ChanceCard{Id: 11, Description: "ADVANCE TOKEN TO NEAREST UTILITY. - IF UNOWNED YOU MAY BUY IT FROM THE BACK. - IF OWNED, THROW DICE AND PAY OWNER A TOTAL TEN TIMES THE AMOUNT THROWN."}
	chanceCards[12] = ChanceCard{Id: 12, Description: "ADVANCE TO ST. CHARLES PLACE - IF YOU PASS GO, COLLECT $200"}
	chanceCards[13] = ChanceCard{Id: 13, Description: "TAKE A RIDE ON THE READING - IF YOU PASS GO COLLECT $200"}
	chanceCards[14] = ChanceCard{Id: 14, Description: "YOUR BUILDING AND LOAD MATURES - COLLECT $150"}
	chanceCards[15] = ChanceCard{Id: 15, Description: "MAKE GENERAL REPAIRS ON ALL YOUR PROPERTY - FOR EACH HOUSE PAY $25 - FOR EACH HOTEL $100"}

	var communityChestCards [16]CommunityChestCard
	communityChestCards[0] = CommunityChestCard{Id: 0, Description: "GET OUT OF JAIL, FREE - THIS CARD MAY BE KEPT UNTIL NEEDED OR SOLD"}
	communityChestCards[1] = CommunityChestCard{Id: 1, Description: "GO TO JAIL - GO DIRECTLY TO JAIL - DO NOT PASS GO - DO NOT COLLECT $200"}
	communityChestCards[2] = CommunityChestCard{Id: 2, Description: "FROM SALE OF STOCK - YOU GET $45"}
	communityChestCards[3] = CommunityChestCard{Id: 3, Description: "BANK ERROR IN YOUR FAVOR - COLLECT $200"}
	communityChestCards[4] = CommunityChestCard{Id: 4, Description: "PAY HOSPITAL $100"}
	communityChestCards[5] = CommunityChestCard{Id: 5, Description: "DOCTOR'S FEE - PAY $50"}
	communityChestCards[0] = CommunityChestCard{Id: 6, Description: "XMAS FUND MATURES - COLLECT $100"}
	communityChestCards[7] = CommunityChestCard{Id: 7, Description: "RECEIVE FOR SERVICES $25"}
	communityChestCards[8] = CommunityChestCard{Id: 8, Description: "PAY SCHOOL TAX OF $150"}
	communityChestCards[9] = CommunityChestCard{Id: 9, Description: "ADVANCE TO GO (COLLECT $200)"}
	communityChestCards[10] = CommunityChestCard{Id: 10, Description: "YOU HAVE WON SECOND PRIZE IN A BEAUTY CONTEST - COLLECT $10"}
	communityChestCards[11] = CommunityChestCard{Id: 11, Description: "GRAND OPERA OPENING - COLLECT $50 FROM EVERY PLAYER - FOR OPENING NIGHT SEATS"}
	communityChestCards[12] = CommunityChestCard{Id: 12, Description: "INCOME TAX REFUND - COLLECT $20"}
	communityChestCards[13] = CommunityChestCard{Id: 13, Description: "YOU ARE ASSESSED FOR STREET REPAIRS - $40 PER HOUSE - $115 PER HOTEL"}
	communityChestCards[14] = CommunityChestCard{Id: 14, Description: "LIFE INSURANCE MATURES - COLLECT $100"}
	communityChestCards[15] = CommunityChestCard{Id: 15, Description: "YOU INHERIT $100"}

	// Randomize
	perm := rand.Perm(len(chanceCards))
	b.ChanceCards = make([]ChanceCard, len(chanceCards))
	for i := 0; i < len(perm); i++ {
		b.ChanceCards[i] = chanceCards[perm[i]]
	}
	perm = rand.Perm(len(communityChestCards))
	b.CommunityChestCards = make([]CommunityChestCard, len(communityChestCards))
	for i := 0; i < len(perm); i++ {
		b.CommunityChestCards[i] = communityChestCards[perm[i]]
	}
}

func (b *Board) initializeSpaces() {
	board := b
	board.Spaces = make([]HandleSpacer, 0)
	board.Spaces = append(board.Spaces, &Space{Name: "Go", Position: 0})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Mediterranean Avenue", Position: 1}, nil, false, 0, 0, 60}, 0, 0, [6]int{2, 10, 30, 90, 160, 250}, 50})
	board.Spaces = append(board.Spaces, &CommunityChest{Space{Name: "Community Chest", Position: 2}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Baltic Avenue", Position: 3}, nil, false, 0, 0, 60}, 0, 0, [6]int{4, 20, 60, 180, 320, 450}, 50})
	board.Spaces = append(board.Spaces, &IncomeTax{Space{Name: "Income Tax", Position: 4}})
	board.Spaces = append(board.Spaces, &RailRoad{OwnableSpace{Space{Name: "Reading Railroad", Position: 5}, nil, false, 0, 0, 200}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Oriental Avenue", Position: 6}, nil, false, 0, 0, 100}, 1, 0, [6]int{6, 30, 90, 270, 400, 550}, 50})
	board.Spaces = append(board.Spaces, &Chance{Space{Name: "Chance", Position: 7}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Vermont Avenue", Position: 8}, nil, false, 0, 0, 100}, 1, 0, [6]int{6, 30, 90, 270, 400, 550}, 50})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Connecticut Avenue", Position: 9}, nil, false, 0, 0, 120}, 1, 0, [6]int{8, 40, 100, 300, 450, 600}, 50})
	board.Spaces = append(board.Spaces, &Space{Name: "Jail", Position: 10})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "St. Charles Place", Position: 11}, nil, false, 0, 0, 140}, 2, 0, [6]int{10, 50, 150, 450, 625, 750}, 100})
	board.Spaces = append(board.Spaces, &Utility{OwnableSpace{Space{Name: "Electric Company", Position: 12}, nil, false, 0, 0, 150}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "States Avenue", Position: 13}, nil, false, 0, 0, 140}, 2, 0, [6]int{10, 50, 150, 450, 625, 750}, 100})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Virginia Avenue", Position: 14}, nil, false, 0, 0, 160}, 2, 0, [6]int{12, 60, 180, 500, 700, 900}, 100})
	board.Spaces = append(board.Spaces, &RailRoad{OwnableSpace{Space{Name: "Pennsylvania Railroad", Position: 15}, nil, false, 0, 0, 200}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "St. James Place", Position: 16}, nil, false, 0, 0, 180}, 3, 0, [6]int{14, 70, 200, 550, 750, 950}, 100})
	board.Spaces = append(board.Spaces, &CommunityChest{Space{Name: "Community Chest", Position: 17}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Tennessee Avenue", Position: 18}, nil, false, 0, 0, 180}, 3, 0, [6]int{14, 70, 200, 550, 750, 950}, 100})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "New York Avenue", Position: 19}, nil, false, 0, 0, 200}, 3, 0, [6]int{16, 80, 220, 600, 800, 1000}, 100})
	board.Spaces = append(board.Spaces, &Space{Name: "Free Parking", Position: 20})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Kentucky Avenue", Position: 21}, nil, false, 0, 0, 220}, 4, 0, [6]int{18, 90, 250, 700, 875, 1050}, 150})
	board.Spaces = append(board.Spaces, &Chance{Space{Name: "Chance", Position: 22}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Indiana Avenue", Position: 23}, nil, false, 0, 0, 220}, 4, 0, [6]int{18, 90, 250, 700, 875, 1050}, 150})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Illinois Avenue", Position: 24}, nil, false, 0, 0, 240}, 4, 0, [6]int{20, 100, 300, 750, 925, 1100}, 150})
	board.Spaces = append(board.Spaces, &RailRoad{OwnableSpace{Space{Name: "B. & O. Railroad", Position: 25}, nil, false, 0, 0, 200}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Atlantic Avenue", Position: 26}, nil, false, 0, 0, 260}, 5, 0, [6]int{22, 110, 330, 800, 975, 1150}, 150})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Ventnor Avenue", Position: 27}, nil, false, 0, 0, 260}, 5, 0, [6]int{22, 110, 330, 800, 975, 1150}, 150})
	board.Spaces = append(board.Spaces, &Utility{OwnableSpace{Space{Name: "Water Works", Position: 28}, nil, false, 0, 0, 150}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Marvin Gardens", Position: 29}, nil, false, 0, 0, 280}, 5, 0, [6]int{24, 120, 360, 850, 1025, 1200}, 150})
	board.Spaces = append(board.Spaces, &GoToJail{Space{Name: "Go To Jail", Position: 30}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Pacific Avenue", Position: 31}, nil, false, 0, 0, 300}, 6, 0, [6]int{26, 130, 390, 900, 1100, 1275}, 200})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "North Carolina Avenue", Position: 32}, nil, false, 0, 0, 300}, 6, 0, [6]int{26, 130, 390, 900, 1100, 1275}, 200})
	board.Spaces = append(board.Spaces, &CommunityChest{Space{Name: "Community Chest", Position: 33}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Pennsylvania Avenue", Position: 34}, nil, false, 0, 0, 320}, 6, 0, [6]int{28, 150, 450, 1000, 1200, 1400}, 200})
	board.Spaces = append(board.Spaces, &RailRoad{OwnableSpace{Space{Name: "Short Line", Position: 35}, nil, false, 0, 0, 200}})
	board.Spaces = append(board.Spaces, &Chance{Space{Name: "Chance", Position: 36}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Park Place", Position: 37}, nil, false, 0, 0, 350}, 7, 0, [6]int{35, 175, 500, 1100, 1300, 1500}, 200})
	board.Spaces = append(board.Spaces, &LuxaryTax{Space{Name: "Luxury Tax", Position: 38}})
	board.Spaces = append(board.Spaces, &Property{OwnableSpace{Space{Name: "Boardwalk", Position: 39}, nil, false, 0, 0, 400}, 7, 0, [6]int{50, 200, 600, 1400, 1700, 2000}, 200})
}
