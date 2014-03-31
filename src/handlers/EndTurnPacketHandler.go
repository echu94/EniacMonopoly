package handlers

type changeTurnPacket struct {
	Id   string
	Turn int
}

type endTurnPacketHandler struct {
	jsonPacketHandler
}

func loadEndTurnPacketHandler() {
	handler := endTurnPacketHandler{jsonPacketHandler{Id: "EndTurn"}}
	jsonPacketHandlers[handler.Id] = handler
}

func (h endTurnPacketHandler) handlePacket(data string) interface{} {
	if !board.HasRolled {
		return nil
	}

	board.NextTurn()

	return changeTurnPacket{Id: "ChangeTurn", Turn: board.Turn}
}
