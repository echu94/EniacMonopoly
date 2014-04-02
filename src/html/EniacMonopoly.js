$(function () {
	var socket = new WebSocket('ws://localhost:7765/websocket');
	var handlers = {};
	
	// load handlers
	handlers.Roll = function (data) {
		var text = 'Rolled a ' + data.Dice1 + ' and a ' + data.Dice2 + '.';
		if(data.Dice1 == data.Dice2) {
			text += ' Doubles!';	
		}	
		Log(text, 1);		
	};
	handlers.SetPlayerPosition = function (data) {
		var player = board.GetCurrentPlayer();
		player.Position = data.Position;
		
		UpdatePlayerSpace(board.Turn);
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
			AddToken(i);
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
		$player = $($('#PlayerInfos > template').html());
		$player.addClass('Player' + id);
		$player.find('.PlayerId').text(id + 1);
		$('#PlayerInfos').append($player);
	}
	
	function AddToken(id) {
		$token = $($('#Tokens > template').html());
		// todo: remove id
		$token.attr('id', 'Token' + id);
		$token.addClass('Player' + id);
		$('#Tokens').append($token);
	}
	
	function UpdatePlayerSpace(id) {
		var position = board.Players[id].Position;
		
		var width = 73;
		var edgeOffset = 73;
		var offset = 5 + 20 * id;
		var css = {
			top: '',
			right: '',
			bottom: '',
			left: ''
		};
		if(position == 0) {
			css.right = offset + 'px';
			css.bottom = offset + 'px';
		}
		else if(position < 10) {
			css.bottom = offset + 'px';
			css.right = position * width + edgeOffset + 'px';
		}
		else if(position == 10) {
			css.bottom = offset + 'px';
			css.left = offset + 'px';
		}
		else if(position < 20) {
			css.bottom = (position % 10) * width + edgeOffset + 'px';
			css.left = offset + 'px';
		}
		else if(position == 20) {
			css.top = offset + 'px';
			css.left = offset + 'px';
		}
		else if(position < 30) {
			css.top = offset + 'px';
			css.left = (position % 10) * width + edgeOffset + 4 + 'px';
		}
		else if(position == 30) {
			css.top = offset + 'px';
			css.right = offset + 'px';
		}
		else {
			css.top = (position % 10) * width + edgeOffset + 5 + 'px';
			css.right = offset + 'px';
		}
		
		$('#Token' + id).css(css);
	}
	
	function UpdatePlayerCash(id) {
		var $player = $('#PlayerInfos').children('.Player' + id).find('.PlayerCash').text(board.Players[id].Cash);
	}
	
	function UpdatePlayerTurn() {
		$('#CurrentPlayer').text(board.Turn + 1);
		Log("It's Player " + (board.Turn + 1) + "'s turn.");
		$('#Tokens > div').removeClass('Active');
		$('#Token' + board.Turn).addClass('Active');
		$('.PlayerInfo').removeClass('Border0');
		$('.PlayerInfo.Player' + board.Turn).addClass('Border0');
		$('#ActivePlayer').attr('class', 'Player' + board.Turn);
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