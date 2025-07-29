package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	result := ""
	switch {
	case playerValue == 22:
		result = "P"
	case playerValue == 21 && dealerValue >= 10:
		result = "S"
	case playerValue == 21 && dealerValue < 10:
		result = "W"
	case playerValue >= 17 && playerValue <= 20:
		result = "S"
	case playerValue >= 12 && playerValue <= 16 && dealerValue >= 7:
		result = "H"
	case playerValue >= 12 && playerValue <= 16 && dealerValue < 7:
		result = "S"
	case playerValue <= 11:
		result = "H"
	}
	return result
}
