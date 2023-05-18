package game

// CardColor is a type for card color
type CardColor string

// Enum of card colors
const (
	Clubs    CardColor = "Clubs"
	Diamonds CardColor = "Diamonds"
	Hearts   CardColor = "Hearts"
	Spades   CardColor = "Spades"
)

// CardNumber is a type for car number
type CardNumber string

// Enum of card numbers
const (
	Ace   CardNumber = "Ace"
	Two   CardNumber = "Two"
	Three CardNumber = "Three"
	Four  CardNumber = "Four"
	Five  CardNumber = "Five"
	Six   CardNumber = "Six"
	Seven CardNumber = "Seven"
	Eight CardNumber = "Eight"
	Nine  CardNumber = "Nine"
	Ten   CardNumber = "Ten"
	Jack  CardNumber = "Jack"
	Queen CardNumber = "Queen"
	King  CardNumber = "King"
)

// Card is a struct for a card
type Card struct {
	ID     int
	Color  CardColor
	Number CardNumber
	Value  int
}
