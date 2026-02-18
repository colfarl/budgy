# Budgy - A Budgeting App In the Terminal

Canonical personal finance project, but it's in the terminal. TUI is WIP, current state is a stable CLI to centralize my personal finances. 
This is my first unguided personal project in Go, mainly meant to explore the language.

Budgy supports multiple users, import from CSV, transaction splitting, and other basic budgeting functionality.

## Design
One of the main goals of this project was to create a flexible backend that could serve to two different UIs. The general flow is as follows:
User Input -> core.Command -> core.Effect -> core.Event -> UI. Commands capture user intent, effects run all necessary side effects, and event captures
the result of the command and effects. 

## Known Limitations
 - money represented as float64, planned move to integer cents
 - error handling and propagration is flimsy

## Usage


```
git clone https://github.com/colfarl/budgy.git
cd budgy
```

budgy expects .env file: 
```
GOOSE_DRIVER=sqlite
GOOSE_DBSTRING= path/to/sqlite/db
GOOSE_MIGRATION_DIR= path/to/budgy/sql/schema
DB_URL=GOOSE_DBSTRING?_foreign_keys=1
```

build cli (tui is not operational): 
```
go install ./cmd/budgy
```

## full list of commands
Usage: budgy <command>

```
Flags:
  -h, --help    Show context-sensitive help.

Commands:
  user add <name>
    Add a user.

  user list
    List users.

  user delete <name>
    Delete a user.

  user balances <name>
    list balances for all accounts.

  account add <username> <account-name>
    Add account to user

  account delete <username> <account-name>
    Remove account from user

  account list <username>
    list all accounts for user

  account balance <username> <account-name>
    list balances for one accounts.

  txn add <username> <account-name> <amount> <description> <date> [flags]
    Add transaction to user account

  txn delete <txn-id>
    Remove transaction from user account

  txn list <username> <account-name> [flags]
    list all transactions for a user account

  txn import-file <username> <account-name> <file-path> <file-origin> <file-type>
    import transactions from a file.

  txn split <txn-id> [flags]
    split one transaction into muliple

  txn categorize <txn-id> <category>
    assign transaction to a category

  txn uncategorize <txn-id>
    uncategorize command

  category add <name> [flags]
    Add Category to budgy

  category delete <category-id>
    Remove Category from Budgy

  category list
    list all categories
    
Run "budgy <command> --help" for more information on a command.
```
