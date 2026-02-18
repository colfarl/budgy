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

type AccountBalanceCmd struct {
	Username string `arg:"" help:"User who owns Account."`
	AccountName string `arg:"" help:"Account Name."`
	StartDate string `help:"sum transactions only after this date if passed." short:"s"`
	EndDate string `help:"sum transactions only before this date if passed." short:"e"`
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
					"Account \"%s\" created for %s\n",
					v.AccountName,
					v.Username,
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
					"Account \"%s\" deleted for %s\n",
					v.AccountName,
					v.Username,
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

func (a *AccountListCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.ListAccounts{Username: a.Username}		
	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			if v, ok := event.(core.AccountsLoaded); ok {
				if len(v.AccountNames) == 0 {
					fmt.Printf("%s has no accounts\n", a.Username)
				}
				for i, u := range v.AccountNames {
					fmt.Printf("%d. %s\n", i + 1, u)
				}
				return nil
			}

			if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			}
			return fmt.Errorf("Unknown error while creating account")
		}
	}
}

func (a *AccountBalanceCmd) Run(binds *Context) error {
	start, end, err := getStartEndDate(a.StartDate, a.EndDate)
	if err != nil {
		return err
	}
	binds.Store.Commands <- core.GetAccountBalance{
		Username: a.Username, 
		AccountName: a.AccountName,
		StartDate: start,
		EndDate: end,
	}		

	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			if v, ok := event.(core.AccountSummed); ok {
				fmt.Printf(
					"Account \"%s\" has balance %.2f\n",
					v.AccountName,
					v.AccountSum,
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
