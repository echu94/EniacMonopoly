$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	var handlers = {};
	
	// load handlers
	handlers.Hello = function (data) {
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
});