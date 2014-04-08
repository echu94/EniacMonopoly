package handlers

import (
	"models"
)

type sayPacketHandler struct {
	jsonPacketHandler
}

func loadSayPacketHandler() {
	handler := sayPacketHandler{jsonPacketHandler{Id: "Say"}}
	jsonPacketHandlers[handler.Id] = handler
}

func (h sayPacketHandler) handlePacket(data string, room *models.Room) []interface{} {
	board := &room.Board
	packets := make([]interface{}, 0)

	// TODO: Implement source verification
	// TODO: Change PlayerId to source
	packets = append(packets, models.SayPacket{Id: "Say", PlayerId: board.Turn, Data: data})

	return packets
}
