function Space(src) {
	this.Name = src.Name;
	this.Position = src.Position;
	this.Owner = src.Owner || null;
	if(src.Owner) {
		src.Owner.OwnableSpaces.push(this);
	}
	this.IsMonopoly = src.IsMonopoly || false;
	this.MortgageRate = src.MortgageRate || 0;
	this.UnmortgageRate = src.UnmortgageRate || 0;
	this.Cost = src.Cost || 0;
	this.Color = src.Color || -1;
	this.Upgrades = src.Upgrades || 0;
	this.Rent = src.Rent || 0;
	this.UpgradeCost = src.UpgradeCost || 0;
}

function Player(src) {
	//IsHuman           bool
	this.Cash = src.Cash;
	this.Order = src.Order;
	this.OwnableSpaces = []
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
		if(src.Spaces[i].Owner) {
			src.Spaces[i].Owner = this.Players[src.Spaces[i].Owner.Order];
		}
		this.Spaces[i] = new Space(src.Spaces[i])
	}
	this.BuyCost = src.BuyCost;
}

Board.prototype.GetCurrentPlayer = function () {
	return this.Players[this.Turn];
}

Board.prototype.GetCurrentSpace = function () {
	return this.Spaces[this.GetCurrentPlayer().Position];
}