package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"image/color"
	"models"
	"net/http"
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
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	if len(jsonPacketHandlers) == 0 {
		loadPacketHandlers()
	}

	if board.Players == nil {
		initializeBoard()
	}

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

func initializeSpaces() {
	{
		board.Spaces = make([]models.HandleSpacer, 0)
		board.Spaces = append(board.Spaces, &models.Space{Name: "Go", Position: 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Mediterranean Avenue", Position: 1}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.CommunityChest{models.Space{Name: "Community Chest", Position: 2}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Baltic Avenue", Position: 3}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.IncomeTax{models.Space{Name: "Income Tax", Position: 4}})
		board.Spaces = append(board.Spaces, &models.OwnableSpace{models.Space{Name: "Reading Railroad", Position: 5}, nil, false, 0, 0, 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Oriental Avenue", Position: 6}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Chance{models.Space{Name: "Chance", Position: 7}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Vermont Avenue", Position: 8}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Connecticut Avenue", Position: 9}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Space{Name: "Jail", Position: 10})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "St. Charles Place", Position: 11}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Utility{models.OwnableSpace{models.Space{Name: "Electric Company", Position: 12}, nil, false, 0, 0, 0}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "States Avenue", Position: 13}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Virginia Avenue", Position: 14}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.OwnableSpace{models.Space{Name: "Pennsylvania Railroad", Position: 15}, nil, false, 0, 0, 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "St. James Place", Position: 16}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.CommunityChest{models.Space{Name: "Community Chest", Position: 17}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Tennessee Avenue", Position: 18}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "New York Avenue", Position: 19}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Space{Name: "Free Parking", Position: 20})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Kentucky Avenue", Position: 21}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Chance{models.Space{Name: "Chance", Position: 22}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Indiana Avenue", Position: 23}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Illinois Avenue", Position: 24}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.OwnableSpace{models.Space{Name: "B. & O. Railroad", Position: 25}, nil, false, 0, 0, 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Atlantic Avenue", Position: 26}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Ventnor Avenue", Position: 27}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Utility{models.OwnableSpace{models.Space{Name: "Water Works", Position: 28}, nil, false, 0, 0, 0}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Marvin Gardens", Position: 29}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.GoToJail{models.Space{Name: "Go To Jail", Position: 30}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Pacific Avenue", Position: 31}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "North Carolina Avenue", Position: 32}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.CommunityChest{models.Space{Name: "Community Chest", Position: 33}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Pennsylvania Avenue", Position: 34}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.OwnableSpace{models.Space{Name: "Short Line", Position: 35}, nil, false, 0, 0, 0})
		board.Spaces = append(board.Spaces, &models.Chance{models.Space{Name: "Chance", Position: 36}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Park Place", Position: 37}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
		board.Spaces = append(board.Spaces, &models.LuxaryTax{models.Space{Name: "Luxury Tax", Position: 38}})
		board.Spaces = append(board.Spaces, &models.Property{models.OwnableSpace{models.Space{Name: "Boardwalk", Position: 39}, nil, false, 0, 0, 0}, color.RGBA{}, 0, nil, 0})
	}
}
