package handlers

import (
	"models"
)

type endTurnPacketHandler struct {
	jsonPacketHandler
}

func loadEndTurnPacketHandler() {
	handler := endTurnPacketHandler{jsonPacketHandler{Id: "EndTurn"}}
	jsonPacketHandlers[handler.Id] = handler
}

func (h endTurnPacketHandler) handlePacket(data string, room *models.Room) []interface{} {
	board := &room.Board

	if !board.HasRolled || board.BuyCost > 0 {
		return nil
	}

	packets := make([]interface{}, 0)

	board.NextTurn()

	packets = append(packets, models.NextTurnPacket{Id: "NextTurn", Turn: board.Turn, DoublesCount: board.DoublesCount, HasRolled: board.HasRolled})
	return packets
}
