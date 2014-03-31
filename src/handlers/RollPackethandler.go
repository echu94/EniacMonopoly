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
	Id   string
	Cash int
}

type setCashPacket struct {
	Id   string
	Cash int
}

type setDoublesCount struct {
	Id           string
	DoublesCount int
}

type setPlayerPositionPacket struct {
	Id       string
	Position int
}

type setHasRolledPacket struct {
	Id        string
	HasRolled bool
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
		if board.DoublesCount == 2 {
			// TODO: Goto jail
			board.NextTurn()
			packets = append(packets, nextTurnPacket{Id: "NextTurn", Turn: board.Turn, DoublesCount: board.DoublesCount, HasRolled: board.HasRolled})
			return packets
		}

		board.DoublesCount++
		packets = append(packets, setDoublesCount{Id: "SetDoublesCount", DoublesCount: board.DoublesCount})
	} else {
		board.HasRolled = true
		packets = append(packets, setHasRolledPacket{Id: "SetHasRolled", HasRolled: board.HasRolled})
	}

	if player.Position >= spaces {
		player.Cash += 200
		player.Position -= spaces
		packets = append(packets, addCashPacket{Id: "AddCash", Cash: 200})
		packets = append(packets, setCashPacket{Id: "SetCash", Cash: player.Cash})
	}

	packets = append(packets, rollResponsePacket{Id: "Roll", Dice1: r1, Dice2: r2})
	packets = append(packets, setPlayerPositionPacket{Id: "SetPlayerPosition", Position: player.Position})
	(*board.GetCurrentSpace()).HandleSpace()
	return packets
}
