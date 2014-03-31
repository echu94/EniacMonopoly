package models

import (
	"fmt"
)

func (p *Chance) HandleSpace(b *Board) []interface{} {
	card := b.ChanceCards[0]
	fmt.Println("Handling a Chance")

	packets := make([]interface{}, 0)

	switch card.Id {
	// TODO:
	//case 0:
	// Get out of jail
	//case 1:
	// go to jail
	case 3:
		loss := (len(b.Players) - 1) * -50
		b.GetCurrentPlayer().Cash += loss
		packets = append(packets, SetCashPacket{Id: "AddCash", PlayerId: b.Turn, Cash: loss})
		for i := 0; i < len(b.Players); i++ {
			if i != b.Turn {
				b.Players[i].Cash += 50
				packets = append(packets, SetCashPacket{Id: "AddCash", PlayerId: i, Cash: 50})
			}
			packets = append(packets, SetCashPacket{Id: "SetCash", PlayerId: i, Cash: b.Players[i].Cash})
		}
		// TODO: handle possible bankrupt
	}

	if card.Id == 0 {
		// Card is not replaced for get out of jail free
		b.ChanceCards = b.ChanceCards[1:]
	} else {
		b.ChanceCards = append(b.ChanceCards[1:], card)
	}

	return packets
}
