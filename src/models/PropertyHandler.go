package models

import (
	"fmt"
)

func (p *Property) HandleSpace(b *Board) []interface{} {
	fmt.Println("Handling a property")

	packets := make([]interface{}, 0)

	// TODO crate packet helper
	if p.Owner == nil {
		b.BuyCost = p.Cost
		packets = append(packets, BuyCostPacket{Id: "BuyCost", Cost: p.Cost})
	} else if player := b.GetCurrentPlayer(); p.Owner != player {
		rent := p.Rent[p.Upgrades]
		if p.IsMonopoly && p.Upgrades == 0 {
			rent *= 2
		}
		player.Cash -= rent
		packets = append(packets, AddCashPacket{Id: "AddCash", PlayerId: player.Order, Cash: -rent})
		packets = append(packets, SetCashPacket{Id: "SetCash", PlayerId: player.Order, Cash: player.Cash})

		p.Owner.Cash += rent
		packets = append(packets, AddCashPacket{Id: "AddCash", PlayerId: p.Owner.Order, Cash: rent})
		packets = append(packets, SetCashPacket{Id: "SetCash", PlayerId: p.Owner.Order, Cash: p.Owner.Cash})
	}

	return packets
}
