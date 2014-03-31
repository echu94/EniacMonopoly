package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"models"
	"net/http"
)

type jsonPacket struct {
	Data string
	Id   string
}

type HelloPacket struct {
	Id string
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
	loadTestPacketHandler()
	loadRollPacketHandler()
	loadEndTurnPacketHandler()
}

var board models.Board

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	if len(jsonPacketHandlers) == 0 {
		loadPacketHandlers()
	}

	if board.Players == nil {
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
		// TODO: Initialize spaces
		board.Spaces = make([]models.Space, 40)
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
	packets = append(packets, HelloPacket{Id: "Hello"})
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
