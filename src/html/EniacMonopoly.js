$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	var handlers = {};
	
	// load handlers
	handlers.Hello = function () {
		Initialize();
		return '{"Id":"Hello","Data":"{\\"Id\\":\\"Hello\\",\\"Data\\":\\"\\"}"}';
	};
	handlers.Roll = function (data) {
		positions[turn] = (positions[turn] + data.Response) % spaces.length;
		
		UpdateCurrentSpace(turn);
	}
	handlers.ChangeTurn = function (data) {
		turn = data.Turn;
		
		UpdatePlayerTurn();
	}
	
	var turn;
	var positions;
	var players;
	
	
	socket.onmessage = function (msg) {
		var data = JSON.parse(msg.data);
		
		console.log(data);		
		
		var handler = handlers[data.Id];
		if(handler){
			var packet = handler(data);
			if(packet) {
				socket.send(packet);
			}
		}
	};
	
	function Initialize() {
		LoadData();
		
		LoadState();
		
		$('#Roll').on('click', function () {
			socket.send(Packets.GetRollPacket());
		});
		$('#End').on('click', function () {
			socket.send(Packets.GetEndTurnPacket());
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
	
	function LoadState() {
		players = 2;
		turn = 0;
	
		positions = [];
		for(var i = 0; i < players; ++i) {
			positions.push(i);
			UpdateCurrentSpace(i);
		}
		UpdatePlayerTurn();
	}
	
	function UpdateCurrentSpace(id) {
		var $player = $('#PlayerSpaces').children('.Player' + id);
		
		if($player.size() == 0) {
			$player = $($('#PlayerSpaces > template').html())
			$player.attr('class', 'Player' + id);
			$player.find('.PlayerId').text(id + 1);
			$('#PlayerSpaces').append($player);
		}
		
		$player.find('.PlayerSpace').text(spaces[positions[turn]].Name);
	}
	
	function UpdatePlayerTurn() {
		$('#CurrentPlayer').text(turn + 1);
	}
});