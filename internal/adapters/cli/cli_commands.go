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
		Add 		TxnAddCmd			`cmd:"" help:"Add transaction to user account"`
		Delete 		TxnDeleteCmd 		`cmd:"" help:"Remove transaction from user account"`
		List 		TxnListCmd  		`cmd:"" help:"list all transactions for a user account"`
		ImportFile  TxnImportFileCmd	`cmd:"" help:"import transactions from a file."`
	}  `cmd:"" help:"Transaction management"`

	Category struct {
		Add 		CategoryAddCMD		`cmd:"" help:"Add account to user"`
		Delete 		CategoryDeleteCmd 	`cmd:"" help:"Remove account from user"`
		List 		CategoryListCmd  	`cmd:"" help:"list all accounts for user"`
	}  `cmd:"" help:"Category Management"`
}


