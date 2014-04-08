package handlers

import (
	"models"
	"strconv"
)

type roomPacketHandler struct {
	jsonPacketHandler
}

func loadRoomPacketHandler() {
	handler := roomPacketHandler{jsonPacketHandler{Id: "Room"}}
	jsonPacketHandlers[handler.Id] = handler
}

func (h roomPacketHandler) handlePacket(data string, room *models.Room) []interface{} {
	packets := make([]interface{}, 0)

	// TODO: Room timeout? delete game state after 30 minutes?

	// TODO: return room id for new rooms
	id, err := strconv.Atoi(data)
	if err != nil {
		return nil
	}

	// TODO: -1 for new room
	if len(rooms) <= id {
		rooms = append(rooms, models.Room{})
		room = &rooms[id]
		room.Board.Initialize()
	}

	board := &room.Board

	// Send state
	packets = append(packets, models.StatePacket{Id: "State", Board: *board})

	if board.ElapsedTurns == 0 && len(board.Players) < 4 {
		// add player
		//TODO: Send to all clients in room
		p := board.AddPlayer()
		packets = append(packets, models.PlayerPacket{Id: "Player", Player: *p})
	}

	return packets
}
