package handlers

type setTurnPacket struct {
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

func (h endTurnPacketHandler) handlePacket(data string) []interface{} {
	if !board.HasRolled {
		return nil
	}

	packets := make([]interface{}, 0)

	board.NextTurn()

	packets = append(packets, setTurnPacket{Id: "SetTurn", Turn: board.Turn})
	return packets
}
