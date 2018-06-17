package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(card, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
