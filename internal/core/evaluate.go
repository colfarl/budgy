package core

import (
	"fmt"
	"log"
	"time"
)

func Evaluate (s State, c Command) ([]Event, []Effect) {
	switch cmd := c.(type) {

	//User / Session Level 
	case LoadSession:
		return evalLoadSession(s)

	case SetActiveUser:
		u := cmd.Username
		return evalSetActiveUser(s, u)

	case ClearActiveUser:
		return evalClearActiveUser(s)		

	case CreateUser:
		u := cmd.Username
		return evalCreateUser(s, u)		

	case LoadAllUsers:
		log.Printf("Called Load Users")
		return nil, []Effect{FxLoadAllUsers{}}
	
	case GetUserBalances:
		return nil, []Effect{FxGetUserBalances{
			cmd.Username,
			cmd.StartDate,
			cmd.EndDate,
		}}		

	case SumTxnsByCategory:
		return nil, []Effect{FxSumTxnByCategory{
			cmd.Username,
			cmd.StartDate,
			cmd.EndDate,
			cmd.Income,
		}}		

	case DeleteUser:
		log.Printf("Evaluated DeleteUser")
		return nil, []Effect{FxDeleteUser{Username: &cmd.Username}}
	
	//Account Level 
	case CreateAccount:
		return evalCreateAccount(s, cmd.Username, cmd.AccountName)		

	case GetAccountBalance:
		return nil, []Effect{FxGetAccountBalance{cmd.Username, cmd.AccountName, cmd.StartDate, cmd.EndDate}}		

	case DeleteAccount:
		return evalDeleteAccount(s, cmd.Username, cmd.AccountName)

	case ListAccounts:
		return nil, []Effect{FxLoadAccounts{cmd.Username}}

	//Transaction Level 
	case CreateTxn:
		return evalCreateTxn(s, cmd)

	case DeleteTxn:
		return nil, []Effect{FxDeleteTxn{TxnID: cmd.ID}}

	case LoadAccountTxns:
		return nil, []Effect{FxLoadAccountTxns{Username: cmd.Username, AccountName: cmd.AccountName, Uncategorized: cmd.Uncategorized}}
	
	case CategorizeTxn:
		return nil, []Effect{FxCategorizeTxn{ID: cmd.ID, Category: cmd.Category}}

	case UncategorizeTxn:
		return nil, []Effect{FxUncategorizeTxn{ID: cmd.ID}}

	case SplitTxn:
		return nil, []Effect{FxSplitTransaction{ID: cmd.ID, Splits: cmd.Splits}}

	case ImportTxnsFromFile:
		return evalImportTxnsFromFile(s, cmd)

	//Category Level 
	case CreateCategory:
		params := Category{
			Name: cmd.Name,
			IsIncome: cmd.IsIncome,
		}
		return nil, []Effect{FxCreateCategory{Category: params}}

	case DeleteCategory:
		return nil, []Effect{FxDeleteCategory{ID: cmd.ID}}

	case LoadAllCategories:
		return nil, []Effect{FxLoadCategories{}}

	}
	
	return nil, nil
}

func evalDeleteAccount(_ State, username string, accountName string) ([]Event, []Effect) {		
	evs := []Event{}
	fxs := []Effect{FxDeleteAccount{Username: username, AccountName: accountName}} 
	return evs, fxs
}

func evalCreateAccount(_ State, username string, accountName string) ([]Event, []Effect) {		
	evs := []Event{}
	fxs := []Effect{FxCreateAccount{Username: username, AccountName: accountName}} 
	return evs, fxs
}

func evalCreateUser(_ State, username string) ([]Event, []Effect) {	
	// maybe do some handling
	evs := []Event{}
	fxs := []Effect{FxCreateUser{Username: &username}} 
	return evs, fxs
}

func evalClearActiveUser(s State) ([]Event, []Effect) {
	if s.ActiveUser == nil {
		return nil, nil
	}
	fxs := []Effect{FxClearSessionUser{}}
	return nil, fxs
}

func evalSetActiveUser(s State, username string) ([]Event, []Effect) {
	if s.ActiveUser != nil && *s.ActiveUser == username {
		return nil, nil
	}
	evs := []Event{}
	fxs := []Effect{
		FxCreateUser{Username: &username}, // this shouldn't have to be done every time we set active user
		FxSetSessionUser{Username: &username},
	}
	return evs, fxs
}

func evalLoadSession(s State) ([]Event, []Effect) {
	if s.ActiveUser != nil {
		return nil, nil
	}
	return nil, []Effect{FxLoadSession{}}
}


func evalCreateTxn(_ State, cmd CreateTxn) ([]Event, []Effect) {
	layout := "01/02/2006"
	t, err := time.Parse(layout, cmd.Date)
	if err != nil {
		return []Event{GeneralFailure{Err: err}}, nil
	}

	if cmd.Amount < 0 {
		return []Event{GeneralFailure{Err: fmt.Errorf("AMOUNT must be positive\n")}}, nil
	}
	ts := t.Unix()
	 		
	return nil, []Effect{
		FxCreateTxn{
			Username: cmd.Username,
			AccountName: cmd.AccountName,
			Amount: cmd.Amount,
			Description: cmd.Description,
			Date: ts,
			Income: cmd.Income,
		},
	}
}

func evalImportTxnsFromFile(_ State, cmd ImportTxnsFromFile) ([]Event, []Effect) {
	return nil, []Effect{FxImportTxnsFromFile{
		Username: cmd.Username,
		AccountName: cmd.AccountName,
		FilePath: cmd.FileName,
		FileOrigin: cmd.FileOrigin,
		FileType: cmd.FileType,
	}}
}
