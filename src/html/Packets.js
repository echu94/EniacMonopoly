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
	},
	GetBuyPacket: function() {
		var packet = {};
		packet.Id = "Buy";
		return JSON.stringify(packet);
	},
	GetPassPacket: function() {
		var packet = {};
		packet.Id = "Pass";
		return JSON.stringify(packet);
	}
}