package models

import (
	"fmt"
)

func (p *RailRoad) HandleSpace(b *Board) []interface{} {
	fmt.Println("Handling a property")

	packets := make([]interface{}, 0)

	if p.Owner == nil {
		b.BuyCost = p.Cost
		packets = append(packets, BuyCostPacket{Id: "BuyCost", Cost: p.Cost})
	} else if player := b.GetCurrentPlayer(); p.Owner != player {
		// Count railroads owned
		owned := 0
		for i := 0; i < 4; i++ {
			if r := (b.Spaces[i*10+5]).(*RailRoad); r.Owner == p.Owner {
				owned++
			}
		}
		var rent int
		switch owned {
		case 1:
			rent = 25
		case 2:
			rent = 50
		case 3:
			rent = 100
		case 4:
			rent = 200
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
