package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"models"
	"net/http"
	"time"
)

type jsonPacket struct {
	Data string
	Id   string
}

type jsonPacketHandler struct {
	Id string
}

type packetWrapper struct {
	Packets []interface{}
}

type jsonHandlePacketler interface {
	handlePacket(string) []interface{}
}

type statePacket struct {
	Id    string
	Board models.Board
}

var jsonPacketHandlers = make(map[string]jsonHandlePacketler)

func loadPacketHandlers() {
	loadRollPacketHandler()
	loadEndTurnPacketHandler()
}

var board models.Board

func initializeBoard() {
	board.Players = make([]models.Player, 2)
	for i := 0; i < len(board.Players); i++ {
		board.Players[i] = models.Player{
			IsHuman:    true,
			Cash:       1500,
			Order:      i,
			Token:      models.Tokens(i),
			JailedTurn: -1,
		}
	}

	initializeSpaces()

	initializeCards()
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	if len(jsonPacketHandlers) == 0 {
		loadPacketHandlers()
	}

	if board.Players == nil {
		initializeBoard()
	}

	// General purpose random
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Incoming web socket request:", r.URL.Path)
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		fmt.Println(err)
		return
	}

	// Send hello & state
	packets := make([]interface{}, 0)
	packets = append(packets, statePacket{Id: "State", Board: board})
	packet := packetWrapper{Packets: packets}
	if err := conn.WriteJSON(&packet); err != nil {
		fmt.Println("Could not write JSON:", err.Error())
		return
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Could not read message:", err.Error())
			return
		}

		var data jsonPacket
		if err := json.Unmarshal(p, &data); err != nil {
			fmt.Println("Could not read json:", err.Error())
			return
		}

		h, d := jsonPacketHandlers[data.Id]
		if !d {
			fmt.Println("Invalid packet id:", data.Id)
			return
		}

		packets := h.handlePacket(data.Data)
		if packets != nil {
			packet := packetWrapper{Packets: packets}
			if err := conn.WriteJSON(packet); err != nil {
				fmt.Println("Could not write JSON:", err.Error())
				return
			}
		}
	}
}

func initializeCards() {
	var chanceCards [16]models.ChanceCard
	chanceCards[0] = models.ChanceCard{Id: 0, Description: "BANK PAYS YOU DIVIDEND OF $50"}
	chanceCards[1] = models.ChanceCard{Id: 1, Description: "GO DIRECTLY TO JAIL - DO NOT PASS GO, DO NOT COLLECT $200"}
	chanceCards[2] = models.ChanceCard{Id: 2, Description: "YOU HAVE BEEN ELECTED CHAIRMAN OF THE BOARD - PAY EACH PLAYER $50"}
	chanceCards[3] = models.ChanceCard{Id: 3, Description: "ADVANCE TOKEN TO THE NEAREST RAILROAD AND PAY OWNER TWICE THE RENTAL TO WHICH HE IS OTHERWISE ENTITLED. - IF RAILROAD IS UNOWNED, YOU MAY BUY IT FROM THE BANK"}
	chanceCards[4] = models.ChanceCard{Id: 4, Description: "ADVANCE TO ILLINOIS AVE."}
	chanceCards[5] = models.ChanceCard{Id: 5, Description: "ADVANCE TO GO - (COLLECT $200)"}
	chanceCards[6] = models.ChanceCard{Id: 6, Description: "GO BACK 3 SPACES"}
	chanceCards[7] = models.ChanceCard{Id: 7, Description: "PAY POOR TAX OF $15"}
	chanceCards[8] = models.ChanceCard{Id: 8, Description: "ADVANCE TOKEN TO THE NEAREST RAILROAD AND PAY OWNER TWICE THE RENTAL TO WHICH HE IS OTHERWISE ENTITLED. - IF RAILROAD IS UNOWNED, YOU MAY BUY IT FROM THE BANK"}
	chanceCards[9] = models.ChanceCard{Id: 9, Description: "TAKE A WALK ON THE BOARD WALK - ADVANCE TOKEN TO BOARD WALK"}
	chanceCards[10] = models.ChanceCard{Id: 10, Description: "THIS CARD MAY BE KEPT UNTIL NEEDED OR SOLD - GET OUT OF JAIL FREE"}
	chanceCards[11] = models.ChanceCard{Id: 11, Description: "ADVANCE TOKEN TO NEAREST UTILITY. - IF UNOWNED YOU MAY BUY IT FROM THE BACK. - IF OWNED, THROW DICE AND PAY OWNER A TOTAL TEN TIMES THE AMOUNT THROWN."}
	chanceCards[12] = models.ChanceCard{Id: 12, Description: "ADVANCE TO ST. CHARLES PLACE - IF YOU PASS GO, COLLECT $200"}
	chanceCards[13] = models.ChanceCard{Id: 13, Description: "TAKE A RIDE ON THE READING - IF YOU PASS GO COLLECT $200"}
	chanceCards[14] = models.ChanceCard{Id: 14, Description: "YOUR BUILDING AND LOAD MATURES - COLLECT $150"}
	chanceCards[15] = models.ChanceCard{Id: 15, Description: "MAKE GENERAL REPAIRS ON ALL YOUR PROPERTY - FOR EACH HOUSE PAY $25 - FOR EACH HOTEL $100"}

	var communityChestCards [16]models.CommunityChestCard
	communityChestCards[0] = models.CommunityChestCard{Id: 0, Description: "XMAS FUND MATURES - COLLECT $100"}
	communityChestCards[1] = models.CommunityChestCard{Id: 1, Description: "YOU INHERIT $100"}
	communityChestCards[2] = models.CommunityChestCard{Id: 2, Description: "FROM SALE OF STOCK - YOU GET $45"}
	communityChestCards[3] = models.CommunityChestCard{Id: 3, Description: "BANK ERROR IN YOUR FAVOR - COLLECT $200"}
	communityChestCards[4] = models.CommunityChestCard{Id: 4, Description: "PAY HOSPITAL $100"}
	communityChestCards[5] = models.CommunityChestCard{Id: 5, Description: "DOCTOR'S FEE - PAY $50"}
	communityChestCards[6] = models.CommunityChestCard{Id: 6, Description: "GET OUT OF JAIL, FREE - THIS CARD MAY BE KEPT UNTIL NEEDED OR SOLD"}
	communityChestCards[7] = models.CommunityChestCard{Id: 7, Description: "RECEIVE FOR SERVICES $25"}
	communityChestCards[8] = models.CommunityChestCard{Id: 8, Description: "PAY SCHOOL TAX OF $150"}
	communityChestCards[9] = models.CommunityChestCard{Id: 9, Description: "ADVANCE TO GO (COLLECT $200)"}
	communityChestCards[10] = models.CommunityChestCard{Id: 10, Description: "YOU HAVE WON SECOND PRIZE IN A BEAUTY CONTEST - COLLECT $10"}
	communityChestCards[11] = models.CommunityChestCard{Id: 11, Description: "GRAND OPERA OPENING - COLLECT $50 FROM EVERY PLAYER - FOR OPENING NIGHT SEATS"}
	communityChestCards[12] = models.CommunityChestCard{Id: 12, Description: "INCOME TAX REFUND - COLLECT $20"}
	communityChestCards[13] = models.CommunityChestCard{Id: 13, Description: "YOU ARE ASSESSED FOR STREET REPAIRS - $40 PER HOUSE - $115 PER HOTEL"}
	communityChestCards[14] = models.CommunityChestCard{Id: 14, Description: "LIFE INSURANCE MATURES - COLLECT $100"}
	communityChestCards[15] = models.CommunityChestCard{Id: 15, Description: "GO TO JAIL - GO DIRECTLY TO JAIL - DO NOT PASS GO - DO NOT COLLECT $200"}

	// Randomize
	perm := rand.Perm(len(chanceCards))
	board.ChanceCards = make([]models.ChanceCard, len(chanceCards))
	for i := 0; i < len(perm); i++ {
		board.ChanceCards[i] = chanceCards[perm[i]]
	}
	perm = rand.Perm(len(communityChestCards))
	board.CommunityChestCards = make([]models.CommunityChestCard, len(communityChestCards))
	for i := 0; i < len(perm); i++ {
		board.CommunityChestCards[i] = communityChestCards[perm[i]]
	}
}

func initializeSpaces() {
	board.Spaces = make([]models.HandleSpacer, 0)
	board.Spaces = append(board.Spaces, &models.Space{Name: "Go", Position: 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Mediterranean Avenue", Position: 1}, nil, false, 0, 0, 0}, 0, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.CommunityChest{models.Space{Name: "Community Chest", Position: 2}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Baltic Avenue", Position: 3}, nil, false, 0, 0, 0}, 0, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.IncomeTax{models.Space{Name: "Income Tax", Position: 4}})
	board.Spaces = append(board.Spaces, &models.OwnableSpace{models.Space{Name: "Reading Railroad", Position: 5}, nil, false, 0, 0, 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Oriental Avenue", Position: 6}, nil, false, 0, 0, 0}, 1, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Chance{models.Space{Name: "Chance", Position: 7}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Vermont Avenue", Position: 8}, nil, false, 0, 0, 0}, 1, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Connecticut Avenue", Position: 9}, nil, false, 0, 0, 0}, 1, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Space{Name: "Jail", Position: 10})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "St. Charles Place", Position: 11}, nil, false, 0, 0, 0}, 2, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Utility{models.OwnableSpace{models.Space{Name: "Electric Company", Position: 12}, nil, false, 0, 0, 0}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "States Avenue", Position: 13}, nil, false, 0, 0, 0}, 2, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Virginia Avenue", Position: 14}, nil, false, 0, 0, 0}, 2, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.OwnableSpace{models.Space{Name: "Pennsylvania Railroad", Position: 15}, nil, false, 0, 0, 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "St. James Place", Position: 16}, nil, false, 0, 0, 0}, 3, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.CommunityChest{models.Space{Name: "Community Chest", Position: 17}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Tennessee Avenue", Position: 18}, nil, false, 0, 0, 0}, 3, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "New York Avenue", Position: 19}, nil, false, 0, 0, 0}, 3, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Space{Name: "Free Parking", Position: 20})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Kentucky Avenue", Position: 21}, nil, false, 0, 0, 0}, 4, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Chance{models.Space{Name: "Chance", Position: 22}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Indiana Avenue", Position: 23}, nil, false, 0, 0, 0}, 4, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Illinois Avenue", Position: 24}, nil, false, 0, 0, 0}, 4, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.OwnableSpace{models.Space{Name: "B. & O. Railroad", Position: 25}, nil, false, 0, 0, 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Atlantic Avenue", Position: 26}, nil, false, 0, 0, 0}, 5, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Ventnor Avenue", Position: 27}, nil, false, 0, 0, 0}, 5, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Utility{models.OwnableSpace{models.Space{Name: "Water Works", Position: 28}, nil, false, 0, 0, 0}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Marvin Gardens", Position: 29}, nil, false, 0, 0, 0}, 5, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.GoToJail{models.Space{Name: "Go To Jail", Position: 30}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Pacific Avenue", Position: 31}, nil, false, 0, 0, 0}, 6, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "North Carolina Avenue", Position: 32}, nil, false, 0, 0, 0}, 6, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.CommunityChest{models.Space{Name: "Community Chest", Position: 33}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Pennsylvania Avenue", Position: 34}, nil, false, 0, 0, 0}, 6, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.OwnableSpace{models.Space{Name: "Short Line", Position: 35}, nil, false, 0, 0, 0})
	board.Spaces = append(board.Spaces, &models.Chance{models.Space{Name: "Chance", Position: 36}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Park Place", Position: 37}, nil, false, 0, 0, 0}, 7, 0, nil, 0})
	board.Spaces = append(board.Spaces, &models.LuxaryTax{models.Space{Name: "Luxury Tax", Position: 38}})
	board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Boardwalk", Position: 39}, nil, false, 0, 0, 0}, 7, 0, nil, 0})
}
