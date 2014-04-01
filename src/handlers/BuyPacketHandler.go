package handlers

import (
	"models"
)

type buyPacketHandler struct {
	jsonPacketHandler
}

func loadBuyPacketHandler() {
	handler := buyPacketHandler{jsonPacketHandler{Id: "Buy"}}
	jsonPacketHandlers[handler.Id] = handler
}

func (h buyPacketHandler) handlePacket(data string) []interface{} {
	if board.BuyCost == 0 {
		return nil
	}

	packets := make([]interface{}, 0)
	player := board.GetCurrentPlayer()
	if p, ok := (*board.GetCurrentSpace()).(*models.Property); ok {
		p.Owner = player
	} else if r, ok := (*board.GetCurrentSpace()).(*models.RailRoad); ok {
		r.Owner = player
	} else if u, ok := (*board.GetCurrentSpace()).(*models.Utility); ok {
		u.Owner = player
	}

	player.Cash -= board.BuyCost
	packets = append(packets, models.AddCashPacket{Id: "AddCash", PlayerId: board.Turn, Cash: -board.BuyCost})
	packets = append(packets, models.SetCashPacket{Id: "SetCash", PlayerId: board.Turn, Cash: player.Cash})

	board.BuyCost = 0
	packets = append(packets, models.BuyCostPacket{Id: "BuyCost", Cost: board.BuyCost})
	packets = append(packets, models.PropertyOwnerPacket{Id: "PropertyOwner", PropertyId: player.Position, PlayerId: player.Order})

	return packets
}
