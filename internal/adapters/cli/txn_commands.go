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

type TxnSplitCmd struct {
	TxnID   int64   `arg:"" help:"Transaction id to split."`
	Splits 	[]string `help:"Repeatable split spec: <amount>[:<category>]"`
}

type TxnImportFileCmd struct {
	Username 	string 	`arg:"" help:"User who owns Account."`
	AccountName string 	`arg:"" help:"Account Name."`
	FilePath	string	`arg:"" help:"File name to read"`
	FileOrigin	string	`arg:"" enum:"boa, citi" help:"file downloaded from... "`
	FileType	string 	`arg:"" enum:"csv" help:"file extension."`
}

type TxnDeleteCmd struct {
	TxnID	int64  	`arg:"" help:"ID of transaction to delete"`
}

type TxnCategorizeCmd struct {
	TxnID		int64  	`arg:"" help:"ID of transaction to categorize"`
	Category	string  `arg:"" help:"Name of catgory"`
}

type TxnUncategorizeCmd struct {
	TxnID		int64  	`arg:"" help:"ID of transaction to uncategorize"`
}

type TxnListCmd struct {
	Username string 	`arg:"" help:"Username of transaction owner"`
	AccountName string 	`arg:"" help:"Account of transaction."`
	Description	bool 	`help:"shows description if passed." short:"d"`
	Uncategorized bool 	`help:"shows only uncategorized transactions if passed." short:"u"`
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
				return fmt.Errorf("Unknown error while deleting transaction\n")
			}
		}
	}
}

func (t *TxnListCmd) Run(binds *Context) error {	
	binds.Store.Commands <- core.LoadAccountTxns{
		Username: t.Username,
		AccountName: t.AccountName,
		Uncategorized: t.Uncategorized,
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
					fmt.Printf("	%d. ID: %d, Amount: %.2f, Income: %t",
											i + 1, u.ID, u.Amount, u.Income)
					if t.Description {
						fmt.Printf(" %v", u.Description)
					}
					fmt.Print("\n")
				}
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error while loading transactions\n")
			}

		}
	}
}

func (t *TxnImportFileCmd) Run(binds *Context) error {	
	binds.Store.Commands <- core.ImportTxnsFromFile{
		Username: t.Username,
		AccountName: t.AccountName,
		FileName: t.FilePath,	
		FileType: t.FileType,
		FileOrigin: t.FileOrigin,
	}

	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v\n", binds.Store.State.Error)
		case event := <- binds.Store.Events: 
			if v, ok := event.(core.TxnsImported); ok {
				if len(v.Transactions) == 0 {
					fmt.Printf("did not import any transactions...")
					return nil
				}
				
				fmt.Printf("Transactions uploaded for %s in %s account:\n", v.Transactions[0].Username, v.Transactions[0].AccountName)
				for i, u := range v.Transactions {
					fmt.Printf("	%d. ID: %d, Amount: %.2f, Income: %t\n",
											i + 1, u.ID, u.Amount, u.Income)
				}
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error while importing transactions\n")
			}

		}
	}
}

func (t *TxnUncategorizeCmd) Run(binds *Context) error {	
	binds.Store.Commands <- core.UncategorizeTxn{
		ID: t.TxnID,
	}

	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v\n", binds.Store.State.Error)
		case event := <- binds.Store.Events: 
			if v, ok := event.(core.TxnUncategorized); ok {
				fmt.Printf("transaction %v uncategorized\n", v.ID)
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error while importing transactions\n")
			}

		}
	}
}

func (t *TxnSplitCmd) Run(binds *Context) error {	
	binds.Store.Commands <- core.SplitTxn{
		ID: t.TxnID,
		Splits: t.Splits,	
	}

	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v\n", binds.Store.State.Error)
		case event := <- binds.Store.Events: 
			if v, ok := event.(core.TxnSplit); ok {
				fmt.Printf("OLD: ID: %v; AMOUNT: %v\n\n", v.Old.ID, v.Old.Amount)
				for _, u := range v.New {
					fmt.Printf("NEW: AMOUNT: %v", u[0])
					if u[1] != "" {
						fmt.Printf("CATEGORY: %v", u[1])
					}
					fmt.Println()
				}
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error while importing transactions\n")
			}

		}
	}
}

func (t *TxnCategorizeCmd) Run(binds *Context) error {	
	binds.Store.Commands <- core.CategorizeTxn{
		ID: t.TxnID,
		Category: t.Category,	
	}

	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v\n", binds.Store.State.Error)
		case event := <- binds.Store.Events: 
			if v, ok := event.(core.TxnCategorized); ok {
				fmt.Printf("transaction %v categorized as %v\n", v.ID, v.Category)	
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error while importing transactions\n")
			}

		}
	}
}


