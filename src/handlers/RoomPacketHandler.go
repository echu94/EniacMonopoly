package handlers

import (
	"models"
	"strconv"
)

type roomPacketHandler struct {
	jsonPacketHandler
}

func handleRoomPacket(data string, c chan []interface{}) (int, []interface{}) {

	packets := make([]interface{}, 0)

	// TODO: Room timeout? delete game state after 30 minutes?

	// TODO: return room id for new rooms
	id, err := strconv.Atoi(data)
	if err != nil {
		return -1, nil
	}

	var board *models.Board
	var room *models.Room

	// TODO: -1 for new room
	if len(rooms) <= id {
		r := models.Room{}
		r.Clients = make([]chan []interface{}, 0)
		r.Board.Initialize()
		rooms = append(rooms, r)
	}

	room = &rooms[id]
	board = &room.Board

	room.Clients = append(room.Clients, c)

	// State only to current client
	c <- []interface{}{models.StatePacket{Id: "State", Board: *board}}

	if !board.Started && len(board.Players) < 4 {
		// add player
		p := board.AddPlayer()
		packets = append(packets, models.PlayerPacket{Id: "Player", Player: *p})
	}

	return id, packets
}
