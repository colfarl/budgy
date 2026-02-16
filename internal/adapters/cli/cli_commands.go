package cli

import (
	"context"
	"github.com/colfarl/budgy/internal/store"
)

type Context struct {	
	Ctx	   	context.Context
	Store 	*store.Store
}

type CliCommands struct {
	User struct {
		Add    UserAddCmd    	`cmd:"" help:"Add a user."`
		List   UserListCmd   	`cmd:"" help:"List users."`
		Delete UserDeleteCmd	`cmd:"" help:"Delete a user."`
	}	`cmd:"" help:"User management."`

	Account struct {
		Add AccountAddCmd 			`cmd:"" help:"Add account to user"`
		Delete AccountDeleteCmd 	`cmd:"" help:"Remove account from user"`
		List AccountListCmd  		`cmd:"" help:"list all accounts for user"`
	} `cmd:"" help:"Account management."`

	Txn struct {
		Add 			TxnAddCmd			`cmd:"" help:"Add transaction to user account"`
		Delete 			TxnDeleteCmd 		`cmd:"" help:"Remove transaction from user account"`
		List 			TxnListCmd  		`cmd:"" help:"list all transactions for a user account"`
		ImportFile  	TxnImportFileCmd	`cmd:"" help:"import transactions from a file."`
		Split			TxnSplitCmd			`cmd:"" help:"split one transaction into muliple"`
		Categorize		TxnCategorizeCmd	`cmd:"" help:"assign transaction to a category"`
		Uncategorize 	TxnUncategorizeCmd	`cmd:"" help:"uncategorize command"`
	}  `cmd:"" help:"Transaction management"`

	Category struct {
		Add 		CategoryAddCmd		`cmd:"" help:"Add Category to budgy"`
		Delete 		CategoryDeleteCmd 	`cmd:"" help:"Remove Category from Budgy"`
		List 		CategoryListCmd  	`cmd:"" help:"list all categories"`
	}  `cmd:"" help:"Category Management"`
}


