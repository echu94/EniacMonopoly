var Packets = {
	GetRollPacket: function() {
		var packet = {};
		packet.Id = "Roll";
		return JSON.stringify(packet);
	},
	GetEndTurnPacket: function() {
		var packet = {};
		packet.Id = "EndTurn";
		return JSON.stringify(packet);
	}
}