# expense CLI App in Go

This is a simple command-line expense application written in Go. It allows you to manage your expenses throughout the year by adding, deleting, and listing expenses, all stored in a JSON file, not to mention th ability to generate for all the expense or for a certain month on the year. finaly it allows users to export expenses to a CSV file.

## Features

- Users can add an expense with a description and amount.
- Users can update an expense.
- Users can delete an expense.
- Users can view all expenses.
- Users can view a summary of all expenses.
- Users can view a summary of expenses for a specific month (for the current year).
- users can filter expenses by category.
- users can export expenses to a CSV file.

## Usage

### Build and Run

1. Install Go: [Go installation](https://golang.org/dl/)

2. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_folder>
   ```
3. Build the app:
   ```bash
    go build -o expense-tracker
   ```
4. Run the app:

   ```bash
   ./expense-tracker <command> [options]

   ```

### Example

1. Add a new expense:

   ```bash
   ./expense-tracker add --description="Dinner" --amount=40 --category=Food
   ```

2. List all expenses:

   ```bash
   ./expense-tracker list
   ```

3. Delete an expense:

   ```bash
   ./expense-tracker delete --id=123456789
   ```

4. Update an expense:

   ```bash
   ./expense-tracker update -id=234234 --description="Dinner" --amount=40 --category=Food
   ```

5. Generate a summary of the month:

   ```bash
   ./expense-tracker summary --month=3
   ```

6. export to an CSV file:

   ```bash
   ./expense-tracker export --file="expenses.csv"
   ```

### JSON Storage

All expenses are stored in a expenses.json file in the same directory as the app. The file is automatically created if it doesn't exist.
