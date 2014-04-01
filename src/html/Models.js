function Space(src) {
	this.Name = src.Name;
	this.Position = src.Position;
}

function Player(src) {
	//IsHuman           bool
	this.Cash = src.Cash;
	this.Order = src.Order;
	//Properties        []Property
	//Token             Tokens
	this.Position = src.Position;
	this.JailedTurn = src.JailedTurn;
	this.HasJailFreeChance = src.HasJailFreeChance;
	this.HasJailFreeChest = src.HasJailFreeChest;
}

function Board(src) {
	this.Turn = src.Turn;
	//ChanceCards         []Chance
	//CommunityChestCards []CommunityChest
	this.Players = [];
	for(var i = 0; i < src.Players.length; ++i) {
		this.Players[i] = new Player(src.Players[i])
	}
	this.DoublesCount = src.DoublesCount
	this.HasRolled = src.HasRolled
	this.Spaces = [];
	for(var i = 0; i < src.Spaces.length; ++i) {
		this.Spaces[i] = new Space(src.Spaces[i])
	}
	this.BuyCost = src.BuyCost;
}

Board.prototype.GetCurrentPlayer = function () {
	return this.Players[this.Turn];
}