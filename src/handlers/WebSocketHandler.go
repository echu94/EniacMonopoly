package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"models"
	"net/http"
	"strconv"
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
	handlePacket(string, *models.Room) []interface{}
}

var jsonPacketHandlers = make(map[string]jsonHandlePacketler)

func loadPacketHandlers() {
	loadRollPacketHandler()
	loadEndTurnPacketHandler()
	loadBuyPacketHandler()
	loadSayPacketHandler()
	loadRoomPacketHandler()
}

var rooms = make([]models.Room, 0)

func getRoom(id int) *models.Room {
	if id == -1 || id >= len(rooms) {
		return nil
	}
	return &rooms[id]
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Move to main
	// Global packet handler initializer
	if len(jsonPacketHandlers) == 0 {
		loadPacketHandlers()
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

	roomId := -1

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

		// Roomless packets (not join room)
		if roomId != -1 && data.Id == "Room" || roomId == -1 && data.Id != "Room" {
			return
		}

		// TODO: better way to chance websocket room id?
		if roomId == -1 {
			roomId, _ = strconv.Atoi(data.Data)
		}

		h, d := jsonPacketHandlers[data.Id]
		if !d {
			fmt.Println("Invalid packet id:", data.Id)
			return
		}

		packets := h.handlePacket(data.Data, getRoom(roomId))
		if packets != nil {
			packet := packetWrapper{Packets: packets}
			if err := conn.WriteJSON(packet); err != nil {
				fmt.Println("Could not write JSON:", err.Error())
				return
			}
		}
	}
}
