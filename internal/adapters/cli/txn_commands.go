package cli

import (
	"fmt"
	"github.com/colfarl/budgy/internal/core"
)

type TxnAddCmd struct {
	Username 	string 	`arg:"" help:"User who owns Account."`
	AccountName string 	`arg:"" help:"Account Name."`
	Amount		float64	`arg:"" help:"transaction amount (positive)."`
	Description	string	`arg:"" help:"transaction description"`
	Date		string	`arg:"" help:"transaction date MM/DD/YYYY"` 
	Income		bool 	`help:"transaction set as income if passed." short:"i"`
}

type TxnImportFileCmd struct {
	Username 	string 	`arg:"" help:"User who owns Account."`
	AccountName string 	`arg:"" help:"Account Name."`
	FilePath	string	`arg:"" help:"File name to read"`
	FileOrigin	string	`arg:"" enum:"boa" help:"file downloaded from... "`
	FileType	string 	`arg:"" enum:"csv" help:"file extension."`
}

type TxnDeleteCmd struct {
	TxnID	int64  	`arg:"" help:"ID of transaction to delete"`
}

type TxnListCmd struct {
	Username string 	`arg:"" help:"Username of transaction owner"`
	AccountName string 	`arg:"" help:"Account of transaction."`
}

func (t *TxnAddCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.CreateTxn{
		Username: t.Username,
		AccountName: t.AccountName,
		Amount: t.Amount,
		Description: t.Description,
		Date: t.Date,
		Income: t.Income,
	}

	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			if v, ok := event.(core.TxnCreated); ok {
				fmt.Printf("Created Transaction ID: %d\n", v.Transaction.ID)
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error occured while creating txn")
			}
		}
	}
}



func (t *TxnDeleteCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.DeleteTxn{ID: t.TxnID}
	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v\n", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			if _, ok := event.(core.TxnDeleted); ok {
				fmt.Println("TRANSACTION DELETED")
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				fmt.Errorf("Unknown error while deleting transaction\n")
			}
		}
	}
}

func (t *TxnListCmd) Run(binds *Context) error {	
	binds.Store.Commands <- core.LoadAccountTxns{
		Username: t.Username,
		AccountName: t.AccountName,
	}

	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v\n", binds.Store.State.Error)
		case event := <- binds.Store.Events: 
			if v, ok := event.(core.AccountTxnsLoaded); ok {
				if len(v.Transactions) == 0 {
					fmt.Printf("%s has no transaction in %s\n", t.Username, t.AccountName)
					return nil
				}
				
				fmt.Printf("Transactions for %s in %s account:\n", v.Transactions[0].Username, v.Transactions[0].AccountName)
				for i, u := range v.Transactions {
					fmt.Printf("	%d. ID: %d, Amount: %.2f, Income: %t\n",
											i + 1, u.ID, u.Amount, u.Income)
				}
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error while deleting transaction\n")
			}

		}
	}
}


