$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	
	socket.onmessage = function (msg) {
		var data = JSON.parse(msg.Data);
		console.log(data);
	};
});