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
	Data   string
	NewKey string
	Id     string
}

func (h testPacketHandler) handlePacket(data string) interface{} {
	var packet testPacket
	if err := json.Unmarshal([]byte(data), &packet); err != nil {
		fmt.Println("Could not read json:", err.Error())
		return nil
	}
	fmt.Println("TPH:", packet)

	return testResponsePacket{Id: "NewId", NewKey: "NewData"}
}
