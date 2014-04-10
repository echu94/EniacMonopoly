package handlers

import (
	"math/rand"
	"models"
)

type rollPacketHandler struct {
	jsonPacketHandler
}

func loadRollPacketHandler() {
	handler := rollPacketHandler{jsonPacketHandler{Id: "Roll"}}
	jsonPacketHandlers[handler.Id] = handler
}

func (h rollPacketHandler) handlePacket(data string, room *models.Room) []interface{} {
	board := &room.Board

	if board.HasRolled || board.BuyCost > 0 {
		return nil
	}

	// TODO: Port to trade complete
	board.Started = true

	packets := make([]interface{}, 0)

	r1 := rand.Intn(6) + 1
	r2 := rand.Intn(6) + 1

	player := board.GetCurrentPlayer()
	player.Position += r1 + r2
	spaces := len(board.Spaces)

	packets = append(packets, models.RollResponsePacket{Id: "Roll", Dice1: r1, Dice2: r2})

	if r1 == r2 {
		if board.DoublesCount == 2 {
			// TODO: Goto jail
			board.NextTurn()
			packets = append(packets, models.NextTurnPacket{Id: "NextTurn", Turn: board.Turn, DoublesCount: board.DoublesCount, HasRolled: board.HasRolled})
			return packets
		}

		board.DoublesCount++
		packets = append(packets, models.SetDoublesCount{Id: "SetDoublesCount", DoublesCount: board.DoublesCount})
	} else {
		board.HasRolled = true
		packets = append(packets, models.SetHasRolledPacket{Id: "SetHasRolled", HasRolled: board.HasRolled})
	}

	if player.Position >= spaces {
		player.Cash += 200
		player.Position -= spaces
		packets = append(packets, models.AddCashPacket{Id: "AddCash", PlayerId: board.Turn, Cash: 200})
		packets = append(packets, models.SetCashPacket{Id: "SetCash", PlayerId: board.Turn, Cash: player.Cash})
	}

	packets = append(packets, models.SetPlayerPositionPacket{Id: "SetPlayerPosition", Position: player.Position})
	if p := (*board.GetCurrentSpace()).HandleSpace(board); p != nil {
		packets = append(packets, p...)
	}
	return packets
}
