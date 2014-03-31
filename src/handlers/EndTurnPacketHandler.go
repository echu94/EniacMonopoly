package handlers

type endTurnResponsePacket struct {
	Id       string
	Response int
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
	board.Turn %= 1
	board.HasRolled = false

	return endTurnResponsePacket{Id: "EndTurn", Response: board.Turn}
}
