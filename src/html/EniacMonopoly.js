$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	var handlers = {};
	
	// load handlers
	handlers.Hello = function (data) {
		Initialize();
		return '{"Id":"Hello","Data":"{\\"Id\\":\\"Hello\\",\\"Data\\":\\"\\"}"}';
	};
	
	
	socket.onmessage = function (msg) {
		var data = JSON.parse(msg.data);
		
		console.log(data);		
		
		var handler = handlers[data.Id];
		if(handler){
			var packet = handler(data.Data);
			if(packet) {
				socket.send(packet);
			}
		}
	};
	
	function getRollPacket() {
		var packet = {};
		packet.Id = "Roll";
		packet.Data = "Roll";
		return JSON.stringify(packet);
	}
	
	function Initialize() {
		$('#Roll').on('click', function () {
			console.log(getRollPacket());
			socket.send(getRollPacket());
		});
	}
});