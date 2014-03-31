package handlers

type nextTurnPacket struct {
	Id           string
	Turn         int
	DoublesCount int
	HasRolled    bool
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

	packets = append(packets, nextTurnPacket{Id: "NextTurn", Turn: board.Turn, DoublesCount: board.DoublesCount, HasRolled: board.HasRolled})
	return packets
}
