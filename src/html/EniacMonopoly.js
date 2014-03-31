$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	var handlers = {};
	
	// load handlers
	handlers.Hello = function () {
		Initialize();
		return '{"Id":"Hello","Data":"{\\"Id\\":\\"Hello\\",\\"Data\\":\\"\\"}"}';
	};
	handlers.Roll = function (data) {
		var player = board.GetCurrentPlayer();
		player.Position = (player.Position + data.Dice1 + data.Dice2) % spaces.length;
		
		$('#Status').text('Rolled a ' + data.Dice1 + ' and a ' + data.Dice2 + '.');
		
		UpdatePlayerSpace(board.Turn);
	}
	handlers.SetTurn = function (data) {
		board.Turn = data.Turn;
		
		UpdatePlayerTurn();
	}
	handlers.AddCash = function (data) {
		var player = board.GetCurrentPlayer();
		
		player.Cash += data.Cash;
		
		UpdatePlayerCash(board.Turn);
	}
	
	var board;
	
	
	socket.onmessage = function (msg) {	
		var packets = JSON.parse(msg.data).Packets;
		
		for(var i = 0; i < packets.length; ++i) {
			var data = packets[i]
			console.log(data);		
			var handler = handlers[data.Id];
			if(handler){
				var packet = handler(data);
				if(packet) {
					socket.send(packet);
				}
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
	
	// TODO: put spaces in board
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
	}
	
	function LoadState() {
		board = new Board();
		var players = 2;
	
		for(var i = 0; i < players; ++i) {
			board.Players.push(new Player(i));
			AddPlayerInfo(i);
			UpdatePlayerSpace(i);
			UpdatePlayerCash(i);
		}
		UpdatePlayerTurn();
	}
	
	function AddPlayerInfo(id) {
		$player = $($('#PlayerInfos > template').html())
		$player.addClass('Player' + id);
		$player.find('.PlayerId').text(id + 1);
		$('#PlayerInfos').append($player);
	}
	
	function UpdatePlayerSpace(id) {
		$('#PlayerInfos').children('.Player' + id).find('.PlayerSpace').text(spaces[board.GetCurrentPlayer().Position].Name);
	}
	
	function UpdatePlayerCash(id) {
		var $player = $('#PlayerInfos').children('.Player' + id).find('.PlayerCash').text(board.GetCurrentPlayer().Cash);
	}
	
	function UpdatePlayerTurn() {
		$('#CurrentPlayer').text(board.Turn + 1);
	}
});