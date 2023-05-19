package game

// CardColor is a type for card color
type CardColor int

// Enum of card colors
const (
	Clubs CardColor = iota
	Diamonds
	Hearts
	Spades
)

func (cardColor CardColor) String() string {
	switch cardColor {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamond"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	default:
		return "Unknown"
	}
}

// CardNumber is a type for car number
type CardNumber int

// Enum of card numbers
const (
	Ace CardNumber = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var numberToValue = map[CardNumber]int{
	Ace:   12,
	Two:   0,
	Three: 1,
	Four:  2,
	Five:  3,
	Six:   4,
	Seven: 5,
	Eight: 6,
	Nine:  7,
	Ten:   8,
	Jack:  9,
	Queen: 10,
	King:  11,
}

func (cardNumber CardNumber) String() string {
	switch cardNumber {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return "Unknown"
	}
}

// Card is a struct for a card
type Card struct {
	ID     int
	Color  CardColor
	Number CardNumber
	Value  int
}

// NewCard creates a new card and return the pointer
func NewCard(ID int, color CardColor, number CardNumber) *Card {
	card := new(Card)
	card.ID = ID
	card.Color = color
	card.Number = number
	card.Value = numberToValue[number]
	return card
}

func (card *Card) String() string {
	return card.Number.String() + " of " + card.Color.String()
}
