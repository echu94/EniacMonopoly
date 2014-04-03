function GameBoard(element) {
	this.Context = element.getContext('2d');
	this.Colors = [ 'green', 'red', 'blue', 'orange' ];
}

GameBoard.prototype.DrawBoard = function () {
	return;
	var context = this.Context;
	
	var board = new Image(); 
	board.src = 'board.jpg'; 
	board.onload = function () {
		context.drawImage(board, 0, 0);
	};
};

GameBoard.prototype.DrawToken = function (id, position, active) {
	var context = this.Context;
	ctx = context;
	
	var coords = this.GetCoords(id, position);
	
	var radius = active ? 10 : 5;
	
	
	context.fillStyle = this.Colors[id];
	context.strokeStyle = 'black';
	context.beginPath();	
	context.arc(coords.X, coords.Y, radius, 0, 2 * Math.PI, true);
	context.fill();
    context.lineWidth = 1;
	context.stroke();
};

GameBoard.prototype.ClearToken = function (id, position) {
	var context = this.Context;
	
	var coords = this.GetCoords(id, position);
	
	context.clearRect(coords.X - 11, coords.Y - 11, 22, 22);
};

GameBoard.prototype.GetCoords = function (id, position) {
	var width = 72;
	var height = (4 - id) * 20;
	var offset = 90;
	
	var coords;
	
	if(position == 0) {
		coords = { X: 900 - height, Y: 900 - height };
	}
	else if(position < 10) {
		coords = { X: (10 - position) * width + offset, Y: 900 - height };
	}
	else if(position == 10) {
		coords = { X: height, Y: 900 - height };
	}
	else if(position < 20) {
		coords = { X: height, Y: (20 - position) * width + offset }
	}
	else if(position == 20) {
		coords = { X: height, Y: height };
	}
	else if(position < 30) {
		coords = { X: (position - 20) * width + offset, Y: height }
	}
	else if(position == 30) {
		coords = { X: 900 - height, Y: height };
	}
	else {
		coords = { X: 900 - height, Y: ( position - 30) * width + offset };
	}
	
	return coords;
};

GameBoard.prototype.DrawOwner = function (id, position) {
	var context = this.Context;
	var coords = this.GetCoords(id, position);
	
	var start = {};
	var end = {};
	
	var h = 33;
	
	if(position < 10) {
		start.X = coords.X - h;
		start.Y = 900;
		end.X = coords.X + h;
		end.Y = 900;
	}
	else if(position < 20) {
		start.X = 0;
		start.Y = coords.Y - h;
		end.X = 0;
		end.Y = coords.Y + h;
	}
	else if(position < 30) {
		start.X = coords.X - h;
		start.Y = 0;
		end.X = coords.X + h;
		end.Y = 0;
	}
	else {
		start.X = 900;
		start.Y = coords.Y - h;
		end.X = 900;
		end.Y = coords.Y + h;
	}
	
	context.strokeStyle = this.Colors[id];
	context.beginPath();	
	context.moveTo(start.X, start.Y);
	context.lineTo(end.X, end.Y);
	context.lineWidth = 12;
	context.stroke();
};