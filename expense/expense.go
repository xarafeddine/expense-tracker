package expense

import (
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category,omitempty"`
}

// New creates a new Expense instance
func New(description string, amount float64, category string) *Expense {

	return &Expense{
		ID:          int(time.Now().Unix()),
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
		Category:    category,
	}
}
