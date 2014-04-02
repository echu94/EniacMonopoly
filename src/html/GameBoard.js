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
	context.beginPath();	
	context.arc(coords.X, coords.Y, radius, 0, 2 * Math.PI, true);
	context.fill();
	context.stroke();
};

GameBoard.prototype.ClearToken = function (id, position) {
	var context = this.Context;
	ctx = context;
	
	var coords = this.GetCoords(id, position);
	
	context.clearRect(coords.X - 11, coords.Y - 11, 22, 22);
};

GameBoard.prototype.GetCoords = function (id, position) {
	var offset = (4 - id) * 20;
	var width = 73;
	
	var coords;
	
	if(position == 0) {
		coords = { X: 900 - offset, Y: 900 - offset };
	}
	else if(position < 10) {
		coords = { X: (10 - position) * width + 85, Y: 900 - offset };
	}
	else if(position == 10) {
		coords = { X: offset, Y: 900 - offset };
	}
	else if(position < 20) {
		coords = { X: offset, Y: (20 - position) * width + 85 }
	}
	else if(position == 20) {
		coords = { X: offset, Y: offset };
	}
	else if(position < 30) {
		coords = { X: (position - 20) * width + 85, Y: offset }
	}
	else if(position == 30) {
		coords = { X: 900 - offset, Y: offset };
	}
	else {
		coords = { X: 900 - offset, Y: ( position - 30) * width + 85 };
	}
	
	return coords;
};