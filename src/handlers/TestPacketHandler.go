package handlers

import (
	"encoding/json"
	"errors"
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

func (h testPacketHandler) handlePacket(data string) (jsonPacket, error) {
	var packet testPacket
	if err := json.Unmarshal([]byte(data), &packet); err != nil {
		fmt.Println("Could not read json:", err.Error())
		return jsonPacket{}, errors.New("Unable to handle packet")
	}
	fmt.Println("TPH:", packet)

	return jsonPacket{Id: "NewId", Data: "NewData"}, nil
}
