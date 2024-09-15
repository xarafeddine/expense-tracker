package cli

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"expense-tracker/expense"
	"expense-tracker/storage"
)

func Run(args []string, store storage.Storage) error {
	if len(args) < 2 {
		return fmt.Errorf("expected 'add', 'update', 'delete', 'list', 'summary' or 'export' subcommands")
	}

	switch args[1] {
	case "add":
		return runAddCommand(args[2:], store)
	case "update":
		return runUpdateCommand(args[2:], store)
	case "delete":
		return runDeleteCommand(args[2:], store)
	case "list":
		return runListCommand(args[2:], store)
	case "summary":
		return runSummaryCommand(args[2:], store)
	case "export":
		return runExportCommand(args[2:], store)
	default:
		return fmt.Errorf("unknown command: %s", args[1])
	}
}

func runAddCommand(args []string, store storage.Storage) error {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	desc := addCmd.String("description", "", "Expense description")
	amount := addCmd.Float64("amount", 0, "Expense amount")
	category := addCmd.String("category", "", "Expense category")
	addCmd.Parse(args)

	if *desc == "" || *amount <= 0 {
		return fmt.Errorf("description and amount must be provided, and amount must be positive")
	}

	exp := expense.New(*desc, *amount, *category)
	if err := store.AddExpense(exp); err != nil {
		return fmt.Errorf("failed to add expense: %w", err)
	}
	fmt.Println("Expense added successfully")
	return nil
}

func runUpdateCommand(args []string, store storage.Storage) error {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	id := updateCmd.Int("id", 0, "Expense ID")
	desc := updateCmd.String("description", "", "New description")
	amount := updateCmd.Float64("amount", 0, "New amount")
	category := updateCmd.String("category", "", "New category")
	updateCmd.Parse(args)

	if *id == 0 || *amount < 0 {
		return fmt.Errorf("id and valid amount must be provided")
	}

	exp, err := store.GetExpense(*id)
	if err != nil {
		return fmt.Errorf("expense not found: %w", err)
	}

	exp.Description = *desc
	exp.Amount = *amount
	exp.Category = *category

	if err := store.UpdateExpense(exp); err != nil {
		return fmt.Errorf("failed to update expense: %w", err)
	}

	fmt.Println("Expense updated successfully")
	return nil
}

func runDeleteCommand(args []string, store storage.Storage) error {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Int("id", 0, "Expense ID")
	deleteCmd.Parse(args)

	if *id == 0 {
		return fmt.Errorf("id must be provided")
	}

	if err := store.DeleteExpense(*id); err != nil {
		return fmt.Errorf("failed to delete expense: %w", err)
	}

	fmt.Println("Expense deleted successfully")
	return nil
}

func runListCommand(args []string, store storage.Storage) error {

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	category := listCmd.String("category", "", "Expense category")
	listCmd.Parse(args)

	expenses := store.ListExpenses()

	if len(expenses) == 0 {
		fmt.Println("No data is found.")
		return nil
	}

	if *category != "" {
		fmt.Println("Filtered by category:", *category)
	}

	fmt.Printf("%-12s %-18s %-18s %-10s %s\n", "ID", "Date", "Description", "Amount", "Category")

	// Print the rows
	var date time.Time
	for _, expense := range expenses {
		if *category != "" && *category != expense.Category {
			continue
		}
		if date.IsZero() || expense.Date.Day() != date.Day() {
			date = expense.Date
			fmt.Println("#", date.Format("2006 January 02"), "-----------------------------------------------------")

		}
		fmt.Printf("%-12d %-18s %-18s %-10.2f %s\n", expense.ID, expense.Date.Format("2006-01-02 15:04"), expense.Description, expense.Amount, expense.Category)
	}

	return nil
}

func runSummaryCommand(args []string, store storage.Storage) error {
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	month := summaryCmd.Int("month", 0, "Month number (1-12)")
	summaryCmd.Parse(args)

	if *month > 0 {
		total, err := store.SummaryByMonth(*month)
		if err != nil {
			return fmt.Errorf("failed to summarize expenses for the month: %w", err)
		}
		fmt.Printf("Total expenses for %s: $%.2f DH\n", strconv.Itoa(*month), total)
	} else {
		total, err := store.TotalExpenses()
		if err != nil {
			return fmt.Errorf("failed to summarize expenses: %w", err)
		}
		fmt.Printf("Total expenses: %.2f DH\n", total)
	}
	return nil

}

func runExportCommand(args []string, store storage.Storage) error {
	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	file := exportCmd.String("file", "expenses.csv", "CSV file to export to")
	exportCmd.Parse(args)

	if err := store.ExportExpenses(*file); err != nil {
		return fmt.Errorf("failed to export expenses: %w", err)
	}

	fmt.Println("Expenses exported successfully")
	return nil
}
