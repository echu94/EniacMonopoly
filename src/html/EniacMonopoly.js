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
		return JSON.stringify(packet);
	}
	
	function Initialize() {
		LoadData();
		
		$('#Roll').on('click', function () {
			console.log(getRollPacket());
			socket.send(getRollPacket());
		});
	}
	
	var spaces;
	
	function LoadData() {
		spaces = [];
		var i = 0;
		spaces.push(new Space("Go", i++));
		spaces.push(new Space("Mediterranean Avenue", i++));
		spaces.push(new Space("Community Chest", i++));
		spaces.push(new Space("Baltic Avenue", i++));
		spaces.push(new Space("Income Tax", i++));
		spaces.push(new Space("Reading Railroad", i++));
		spaces.push(new Space("Oriental Avenue", i++));
		spaces.push(new Space("Chance", i++));
		spaces.push(new Space("Vermont Avenue", i++));
		spaces.push(new Space("Connecticut Avenue", i++));
		spaces.push(new Space("Jail", i++));
		spaces.push(new Space("St. Charles Place", i++));
		spaces.push(new Space("Electric Company", i++));
		spaces.push(new Space("States Avenue", i++));
		spaces.push(new Space("Virginia Avenue", i++));
		spaces.push(new Space("Pennsylvania Railroad", i++));
		spaces.push(new Space("St. James Place", i++));
		spaces.push(new Space("Community Chest", i++));
		spaces.push(new Space("Tennessee Avenue", i++));
		spaces.push(new Space("New York Avenue", i++));
		spaces.push(new Space("Free Parking", i++));
		spaces.push(new Space("Kentucky Avenue", i++));
		spaces.push(new Space("Chance", i++));
		spaces.push(new Space("Indiana Avenue", i++));
		spaces.push(new Space("Illinois Avenue", i++));
		spaces.push(new Space("B. & O. Railroad", i++));
		spaces.push(new Space("Atlantic Avenue", i++));
		spaces.push(new Space("Ventnor Avenue", i++));
		spaces.push(new Space("Water Works", i++));
		spaces.push(new Space("Marvin Gardens", i++));
		spaces.push(new Space("Go To Jail", i++));
		spaces.push(new Space("Pacific Avenue", i++));
		spaces.push(new Space("North Carolina Avenue", i++));
		spaces.push(new Space("Community Chest", i++));
		spaces.push(new Space("Pennsylvania Avenue", i++));
		spaces.push(new Space("Short Line", i++));
		spaces.push(new Space("Chance", i++));
		spaces.push(new Space("Park Place", i++));
		spaces.push(new Space("Luxury Tax", i++));
		spaces.push(new Space("Boardwalk", i++));
		console.log(spaces);
	}
	
});