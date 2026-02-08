package cli

import (
	"fmt"
	"log"

	"github.com/colfarl/budgy/internal/core"
)

type AccountAddCmd struct {
	Username string `arg:"" help:"User who owns Account."`
	AccountName string `arg:"" help:"Account Name."`
}

type AccountDeleteCmd struct {
	Username string `arg:"" help:"User who owns Account."`
	AccountName string `arg:"" help:"Account Name."`
}

type AccountListCmd struct {Username string `arg:"" help:"Username of accounts."`}

func (a *AccountAddCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.CreateAccount{Username: a.Username, AccountName: a.AccountName}		
	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			if v, ok := event.(core.AccountCreated); ok {
				fmt.Printf(
					"Account \"%s\"(%d) created for %s (%d)\n",
					v.AccountName,
					v.AccountID,
					v.Username,
					v.UserID,
				)
				return nil
			}
			if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			}
			return fmt.Errorf("Unknown error while creating account")
		}
	}
}

func (a *AccountDeleteCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.DeleteAccount{Username: a.Username, AccountName: a.AccountName}		
	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			log.Println("account was deleted?")
			if v, ok := event.(core.AccountDeleted); ok {
				fmt.Printf(
					"Account \"%s\" deleted for %s (%d)\n",
					v.AccountName,
					v.Username,
					v.UserID,
				)
				return nil
			}
			if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			}
			return fmt.Errorf("Unknown error while creating account")
		}
	}
}




