$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	var handlers = {};
	
	// load handlers
	handlers.Roll = function (data) {
		$('#Status').text('Rolled a ' + data.Dice1 + ' and a ' + data.Dice2 + '.');		
	};
	handlers.SetPlayerPosition = function (data) {
		var player = board.GetCurrentPlayer();
		player.Position = data.Position;
		
		UpdatePlayerSpace(board.Turn);
	};
	handlers.NextTurn = function (data) {
		board.Turn = data.Turn;
		board.DoublesCount = data.DoublesCount;
		board.HasRolled = data.HasRolled;
		
		UpdateRolled();
		UpdatePlayerTurn();
	};
	handlers.AddCash = function (data) {
		$('#Status').text('Received ' + data.Cash + ' dollars.');	
	};
	handlers.State = function (data) {
		board = new Board(data.Board);
		
		Initialize();
	};
	handlers.setDoublesCount = function (data) {
		board.DoublesCount = data.DoublesCount;
	};
	handlers.SetHasRolled = function (data) {
		board.HasRolled = data.HasRolled;
		
		UpdateRolled();
	};
	handlers.SetCash = function (data) {
		var player = board.GetCurrentPlayer();
		
		player.Cash = data.Cash;
		
		UpdatePlayerCash(board.Turn);
	};
	
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
		for(var i = 0; i < board.Players.length; ++i) {
			AddPlayerInfo(i);
			UpdatePlayerSpace(i);
			UpdatePlayerCash(i);
		}
		UpdatePlayerTurn();
		UpdateRolled()
		
		$('#Roll').on('click', function () {
			socket.send(Packets.GetRollPacket());
		});
		$('#End').on('click', function () {
			socket.send(Packets.GetEndTurnPacket());
		});
	}
	
	function AddPlayerInfo(id) {
		$player = $($('#PlayerInfos > template').html())
		$player.addClass('Player' + id);
		$player.find('.PlayerId').text(id + 1);
		$('#PlayerInfos').append($player);
	}
	
	function UpdatePlayerSpace(id) {
		$('#PlayerInfos').children('.Player' + id).find('.PlayerSpace').text(board.Spaces[board.GetCurrentPlayer().Position].Name);
	}
	
	function UpdatePlayerCash(id) {
		var $player = $('#PlayerInfos').children('.Player' + id).find('.PlayerCash').text(board.GetCurrentPlayer().Cash);
	}
	
	function UpdatePlayerTurn() {
		$('#CurrentPlayer').text(board.Turn + 1);
	}
	
	function UpdateRolled() {
		$('#Roll').prop('disabled', board.HasRolled);
		$('#End').prop('disabled', !board.HasRolled);
	}
});