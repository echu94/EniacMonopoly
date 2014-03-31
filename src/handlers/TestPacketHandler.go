package handlers

import (
	"encoding/json"
	"fmt"
)

func loadTestPacketHandler() {
	handler := testPacketHandler{jsonPacketHandler{Id: "Hello"}}
	jsonPacketHandlers[handler.Id] = handler
}

type testPacketHandler struct {
	jsonPacketHandler
}

type testPacket struct {
	Data string
	Id   string
}

type testResponsePacket struct {
	NewKey string
	Id     string
}

func (h testPacketHandler) handlePacket(data string) []interface{} {
	var packet testPacket
	if err := json.Unmarshal([]byte(data), &packet); err != nil {
		fmt.Println("Could not read json:", err.Error())
		return nil
	}

	packets := make([]interface{}, 0)

	packets = append(packets, testResponsePacket{Id: "NewId", NewKey: "NewData"})
	return packets
}
