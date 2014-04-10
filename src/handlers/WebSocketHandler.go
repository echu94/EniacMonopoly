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
}

// TODO: This should be a map
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
	var room *models.Room
	var c chan []interface{}

	// Room packet handler
	{
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
		if data.Id != "Room" {
			return
		}

		c = make(chan []interface{}, 5)

		rId, packets := handleRoomPacket(data.Data, c)

		if rId == -1 {
			return
		}

		roomId = rId

		room = &rooms[roomId]
		room.Send(packets)

		go func() {
			for packets := range c {
				packet := packetWrapper{Packets: packets}
				if err := conn.WriteJSON(packet); err != nil {
					fmt.Println("Could not write JSON:", err.Error())
					return
				}
			}
		}()
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Could not read message:", err.Error())
			i := 0
			for ; i < len(room.Clients); i++ {
				if room.Clients[i] == c {
					break
				}
			}
			room.Clients[i], room.Clients = room.Clients[len(room.Clients)-1], room.Clients[:len(room.Clients)-1]
			close(c)
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

		packets := h.handlePacket(data.Data, getRoom(roomId))
		room.Send(packets)
	}
}
