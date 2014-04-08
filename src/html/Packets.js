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
	},
	GetSayPacket: function(data) {
		var packet = {};
		packet.Id = "Say";
		packet.Data = data;
		return JSON.stringify(packet);
	},
	GetRoomPacket: function(id) {
		var packet = {};
		packet.Id = "Room";
		packet.Data = id.toString();
		return JSON.stringify(packet);
	}
}