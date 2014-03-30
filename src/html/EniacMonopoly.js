$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	var handlers = {};
	
	// load handlers
	handlers.Hello = function (data) {
		socket.send('{"Id":"Hello","Data":"{\\"Id\\":\\"Hello\\",\\"Data\\":\\"\\"}"}');
	};
	
	
	socket.onmessage = function (msg) {
		var data = JSON.parse(msg.data);
		
		var handler = handlers[data.Id];
		if(handler){
			handler(data.Data);
		}
	};
});