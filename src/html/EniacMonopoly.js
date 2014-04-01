$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	var handlers = {};
	
	// load handlers
	handlers.Roll = function (data) {
		var text = 'Rolled a ' + data.Dice1 + ' and a ' + data.Dice2 + '.';
		if(data.Dice1 == data.Dice2) {
			text += ' Doubles!';	
		}	
		Log(text);		
	};
	handlers.SetPlayerPosition = function (data) {
		var player = board.GetCurrentPlayer();
		player.Position = data.Position;
		
		UpdatePlayerSpace(board.Turn);
		Log('Landed on ' + board.GetCurrentSpace().Name + '.');	
	};
	handlers.NextTurn = function (data) {
		board.Turn = data.Turn;
		board.DoublesCount = data.DoublesCount;
		board.HasRolled = data.HasRolled;
		
		UpdateRolled();
		UpdatePlayerTurn();
	};
	handlers.AddCash = function (data) {
		var verb = data.Cash > 0 ? ' gained ' : ' lost ';
		Log('Player ' + (data.PlayerId + 1) + verb + Math.abs(data.Cash) + ' dollars.');
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
		var player = board.Players[data.PlayerId];
		
		player.Cash = data.Cash;
		
		UpdatePlayerCash(data.PlayerId);
	};
	handlers.BuyCost = function (data) {
		board.BuyCost = data.Cost;
		
		UpdateBuy();
	};
	handlers.PropertyOwner = function (data) {
		board.Spaces[data.PropertyId].Owner = board.Players[data.PlayerId];
	};
	handlers.BuySpace = function (data) {
		Log(board.Spaces[data.PropertyId].Name + ' was bought by Player ' + (board.Turn + 1) + '.');
		// TODO: UI update
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
		
		console.log(board);
	};
	
	function Initialize() {
		for(var i = 0; i < board.Players.length; ++i) {
			AddPlayerInfo(i);
			UpdatePlayerSpace(i);
			UpdatePlayerCash(i);
		}
		UpdatePlayerTurn();
		UpdateBuy();
		
		$('#Roll').on('click', function () {
			socket.send(Packets.GetRollPacket());
		});
		$('#End').on('click', function () {
			socket.send(Packets.GetEndTurnPacket());
		});
		$('#Buy').on('click', function () {
			socket.send(Packets.GetBuyPacket());
		});
		$('#Pass').on('click', function () {
			socket.send(Packets.GetPassPacket());
		});
	}
	
	function AddPlayerInfo(id) {
		$player = $($('#PlayerInfos > template').html())
		$player.addClass('Player' + id);
		$player.find('.PlayerId').text(id + 1);
		$('#PlayerInfos').append($player);
	}
	
	function UpdatePlayerSpace(id) {
		$('#PlayerInfos').children('.Player' + id).find('.PlayerSpace').text(board.Spaces[board.Players[id].Position].Name);
	}
	
	function UpdatePlayerCash(id) {
		var $player = $('#PlayerInfos').children('.Player' + id).find('.PlayerCash').text(board.Players[id].Cash);
	}
	
	function UpdatePlayerTurn() {
		$('#CurrentPlayer').text(board.Turn + 1);
		Log("It's Player " + (board.Turn + 1) + "'s turn.");
	}
	
	function UpdateRolled() {
		$('#Roll').prop('disabled', board.HasRolled || board.BuyCost > 0);
		$('#End').prop('disabled', !board.HasRolled || board.BuyCost > 0);
	}
	
	function UpdateBuy() {
		if(board.BuyCost > 0) {
			$('#Status').text('Buy ' + board.GetCurrentSpace().Name + ' for ' + board.BuyCost + '?');	
		}
		else {
			$('#Status').text('');	
		}
		$('#Buy').prop('disabled', board.BuyCost == 0);
		$('#Pass').prop('disabled', board.BuyCost == 0);
		
		UpdateRolled();
	}
	
	function Log(s) {
		var t = new Date();
		var $textarea = $('textarea');
		$textarea.val($textarea.val() + "\n" + '[' + t.toLocaleString() + '] ' +  s)
		$textarea.scrollTop($textarea[0].scrollHeight);
	}
});