package storage

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"expense-tracker/expense"
)

// Storage is the interface for managing expenses
type Storage interface {
	AddExpense(exp *expense.Expense) error
	GetExpense(id int) (*expense.Expense, error)
	UpdateExpense(exp *expense.Expense) error
	DeleteExpense(id int) error
	ListExpenses() []expense.Expense
	TotalExpenses() (float64, error)
	SummaryByMonth(month int) (float64, error)
	ExportExpenses(filename string) error
}

// JSONStorage is a JSON-based implementation of the Storage interface
type JSONStorage struct {
	filename string
	expenses []expense.Expense
}

// NewJSONStorage initializes JSON storage
func NewJSONStorage(filename string) (*JSONStorage, error) {
	s := &JSONStorage{filename: filename}
	if err := s.load(); err != nil {
		return nil, err
	}
	return s, nil
}

// AddExpense adds an expense to the storage
func (s *JSONStorage) AddExpense(exp *expense.Expense) error {
	s.expenses = append(s.expenses, *exp)
	return s.save()
}

// GetExpense retrieves an expense by ID
func (s *JSONStorage) GetExpense(id int) (*expense.Expense, error) {
	for _, exp := range s.expenses {
		if exp.ID == id {
			return &exp, nil
		}
	}
	return nil, fmt.Errorf("expense not found")
}

// UpdateExpense updates an existing expense
func (s *JSONStorage) UpdateExpense(exp *expense.Expense) error {
	for i, existing := range s.expenses {
		if existing.ID == exp.ID {
			s.expenses[i] = *exp
			return s.save()
		}
	}
	return fmt.Errorf("expense not found")
}

// DeleteExpense removes an expense by ID
func (s *JSONStorage) DeleteExpense(id int) error {
	for i, exp := range s.expenses {
		if exp.ID == id {
			s.expenses = append(s.expenses[:i], s.expenses[i+1:]...)
			return s.save()
		}
	}
	return fmt.Errorf("expense not found")
}

// ListExpenses lists all expenses
func (s *JSONStorage) ListExpenses() []expense.Expense {
	return s.expenses
}

// TotalExpenses returns the total of all expenses
func (s *JSONStorage) TotalExpenses() (float64, error) {
	var total float64
	for _, exp := range s.expenses {
		total += exp.Amount
	}
	return total, nil
}

// SummaryByMonth returns the total expenses for a given month
func (s *JSONStorage) SummaryByMonth(month int) (float64, error) {
	var total float64
	for _, exp := range s.expenses {
		if exp.Date.Month() == time.Month(month) {
			total += exp.Amount
		}
	}
	return total, nil
}

// ExportExpenses exports the expenses to a CSV file
func (s *JSONStorage) ExportExpenses(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV headers
	writer.Write([]string{"ID", "Date", "Description", "Amount", "Category"})

	// Write data
	for _, exp := range s.expenses {
		writer.Write([]string{
			strconv.Itoa(exp.ID),
			exp.Date.Format("2006-01-02"),
			exp.Description,
			fmt.Sprintf("%.2f", exp.Amount),
			exp.Category,
		})
	}
	return nil
}

// load loads expenses from a JSON file
func (s *JSONStorage) load() error {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // If file doesn't exist, start with an empty list
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, &s.expenses)
}

// save saves expenses to a JSON file
func (s *JSONStorage) save() error {
	data, err := json.MarshalIndent(s.expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filename, data, 0644)
}
