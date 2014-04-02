$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	
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
	
	var gameBoard = new GameBoard(document.getElementById('GameBoard'));
	gameBoard.DrawBoard();
	var board;
	
	function AddPlayerInfo(id) {
		$player = $($('#PlayerInfos > template').html());
		$player.addClass('Player' + id);
		$player.find('.PlayerId').text(id + 1);
		$('#PlayerInfos').append($player);
	}
	
	function UpdatePlayerSpace(id, previous) {
		if(previous) {
			gameBoard.ClearToken(id, previous);			
		}
		gameBoard.DrawToken(id, board.Players[id].Position, board.Turn == id);
	}
	
	function UpdatePlayerCash(id) {
		var $player = $('#PlayerInfos').children('.Player' + id).find('.PlayerCash').text(board.Players[id].Cash);
	}
	
	function UpdatePlayerTurn() {
		$('#CurrentPlayer').text(board.Turn + 1);
		Log("It's Player " + (board.Turn + 1) + "'s turn.");
		$('.PlayerInfo').removeClass('Border0');
		$('.PlayerInfo.Player' + board.Turn).addClass('Border0');
		$('#ActivePlayer').attr('class', 'Player' + board.Turn);
		
		for(var i = 0; i < board.Players.length; ++i) {
			gameBoard.ClearToken(i, board.Players[i].Position);	
			gameBoard.DrawToken(i, board.Players[i].Position, i == board.Turn);
		}
		
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
	
	function Initialize() {
		for(var i = 0; i < board.Players.length; ++i) {
			gameBoard.DrawToken(i, board.Players[i].Position);
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
	
	// load handlers
	var handlers = {};
		
	handlers.Roll = function (data) {
		var text = 'Rolled a ' + data.Dice1 + ' and a ' + data.Dice2 + '.';
		if(data.Dice1 == data.Dice2) {
			text += ' Doubles!';	
		}	
		Log(text, 1);		
	};
	handlers.SetPlayerPosition = function (data) {
		var player = board.GetCurrentPlayer();
		var previous = player.Position;
		player.Position = data.Position;
		
		UpdatePlayerSpace(board.Turn, previous);
		Log('Landed on ' + board.GetCurrentSpace().Name + '.', 1);	
	};
	handlers.NextTurn = function (data) {
		board.Turn = data.Turn;
		board.DoublesCount = data.DoublesCount;
		board.HasRolled = data.HasRolled;
		
		UpdateRolled();
		UpdatePlayerTurn();
	};
	handlers.AddCash = function (data) {
		var verb = data.Cash > 0 ? 'Gained ' : 'Lost ';
		Log(verb + Math.abs(data.Cash) + ' dollars.', 1);
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
		Log(board.Spaces[data.PropertyId].Name + ' was bought.', 1);
		// TODO: UI update
	};
	
	// Logger
	function Log(s, level) {
		var t = new Date();
		var $textarea = $('#Log');
		level = level || 0;
		for(var i = 0; i < level; ++i) {
			s = '&nbsp;&nbsp;' + s;
		}
		$textarea.html($textarea.html() + '<br><span class="LogText">' + '[' + t.toLocaleTimeString() + '] ' +  s + '</span>')
		$textarea.scrollTop($textarea[0].scrollHeight);
	}
});