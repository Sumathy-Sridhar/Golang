package main

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
	case "ten", "king", "queen", "jack":
		return 10
	default:
		return 0
	}
}
func isBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

func LargeHand(isBlackjack bool, dealerScore int) string {

	if isBlackjack {
		if dealerScore < 10 {
			return "w"
		}
		return "s"
	}
	return "P"
}
func SmallHand(handScore, dealerScore int) string {
	// Switch statement
	switch {
	case (handScore >= 17) || (handScore >= 12 && dealerScore < 7):
		return "S"
	case (handScore <= 11) || (handScore >= 12 && dealerScore >= 7):
		return "H"
	default:
		return ""
	}
}

func main() {
	SmallHand(20, 25)
	LargeHand(true, 10)
}
