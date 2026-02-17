# Budgy - A Budgeting App In the Terminal

Canonical personal finance project, but it's in the terminal. TUI is in the works, current state is a usable cli to centralize my personal finances. 
Budgy supports multiple users, import from csv, transaction splitting and other basic budgetting functionality.

## Usage
This is not intended for external use but I am open to feedback / will provide set up from source instructions.

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

