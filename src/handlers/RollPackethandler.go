package handlers

import (
	"math/rand"
	"time"
)

type rollResponsePacket struct {
	Id       string
	Response int
}

type rollPacketHandler struct {
	jsonPacketHandler
}

func loadRollPacketHandler() {
	handler := rollPacketHandler{jsonPacketHandler{Id: "Roll"}}
	jsonPacketHandlers[handler.Id] = handler
}

func (h rollPacketHandler) handlePacket(data string) interface{} {
	if board.HasRolled {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	r1 := rand.Intn(6) + 1
	r2 := rand.Intn(6) + 1
	r := r1 + r2
	if r1 == r2 {
		board.DoublesCount++
		if board.DoublesCount == 3 {
			// TODO: Goto jail
			board.NextTurn()
			return changeTurnPacket{Id: "ChangeTurn", Turn: board.Turn}
		}
	} else {
		board.HasRolled = true
	}

	return rollResponsePacket{Id: "Roll", Response: r}
}
