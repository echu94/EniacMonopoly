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

type testResponsePacket struct {
	Data   string
	NewKey string
	Id     string
}

func (h testPacketHandler) handlePacket(data string) (string, interface{}, error) {
	var packet testPacket
	if err := json.Unmarshal([]byte(data), &packet); err != nil {
		fmt.Println("Could not read json:", err.Error())
		return "", nil, errors.New("Unable to handle packet")
	}
	fmt.Println("TPH:", packet)

	return "NewId", testResponsePacket{Id: "NewId", NewKey: "NewData"}, nil
}
