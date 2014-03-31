package handlers

import (
	"math/rand"
	"time"
)

type rollResponsePacket struct {
	Id    string
	Dice1 int
	Dice2 int
}

type rollPacketHandler struct {
	jsonPacketHandler
}

func loadRollPacketHandler() {
	handler := rollPacketHandler{jsonPacketHandler{Id: "Roll"}}
	jsonPacketHandlers[handler.Id] = handler
}

type addCashPacket struct {
	Id       string
	PlayerId int
	Cash     int
}

func (h rollPacketHandler) handlePacket(data string) []interface{} {
	if board.HasRolled {
		return nil
	}

	packets := make([]interface{}, 0)

	rand.Seed(time.Now().UnixNano())
	r1 := rand.Intn(6) + 1
	r2 := rand.Intn(6) + 1

	player := board.GetCurrentPlayer()
	player.Position += r1 + r2
	spaces := len(board.Spaces)

	if r1 == r2 {
		board.DoublesCount++
		if board.DoublesCount == 3 {
			// TODO: Goto jail
			board.NextTurn()
			packets = append(packets, setTurnPacket{Id: "SetTurn", Turn: board.Turn})
			return packets
		}
	} else {
		board.HasRolled = true
	}

	if player.Position >= spaces {
		player.Cash += 200
		player.Position -= spaces
		packets = append(packets, addCashPacket{Id: "AddCash", PlayerId: board.Turn, Cash: 200})
	}

	packets = append(packets, rollResponsePacket{Id: "Roll", Dice1: r1, Dice2: r2})
	return packets
}
