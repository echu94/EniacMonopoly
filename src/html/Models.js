function Space(name, position) {
	this.Name = name;
	this.Position = position;
}

function Player(order) {
	//IsHuman           bool
	this.Cash = 1500;
	this.Order = order;
	//Properties        []Property
	//Token             Tokens
	this.Position = 0;
	this.JailedTurn = -1;
	this.HasJailFreeChance = false;
	this.HasJailFreeChest = false
}

function Board() {
	this.Turn = 0;
	//ChanceCards         []Chance
	//CommunityChestCards []CommunityChest
	this.Players = [];
	//DoublesCount        int
	//HasRolled           bool
}

Board.prototype.GetCurrentPlayer = function () {
	return this.Players[this.Turn];
}