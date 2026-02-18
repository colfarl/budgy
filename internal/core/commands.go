package core 

type Command interface{ isCommand() }

type LoadSession struct{}
func (LoadSession) isCommand() {}

type LoadAllUsers struct{}
func (LoadAllUsers) isCommand() {}

type CreateUser struct{ Username string }
func (CreateUser) isCommand() {}

type GetUserBalances struct{ 
	Username string 
	StartDate int64
	EndDate int64
}
func (GetUserBalances) isCommand() {}

type SumTxnsByCategory struct{ 
	Username string 
	StartDate int64
	EndDate int64
	Income	bool
}
func (SumTxnsByCategory) isCommand() {}

type DeleteUser struct{ Username string }
func (DeleteUser) isCommand() {}

type SetActiveUser struct{ Username string }
func (SetActiveUser) isCommand() {}

type ClearActiveUser struct{}
func (ClearActiveUser) isCommand() {}

// ============================== Account Commands ==============================
type CreateAccount struct{Username string; AccountName string}
func (CreateAccount) isCommand() {}

type GetAccountBalance struct{
	Username string
	AccountName string
	StartDate int64
	EndDate int64
}
func (GetAccountBalance) isCommand() {}

type DeleteAccount struct{Username string; AccountName string}
func (DeleteAccount) isCommand() {}

type ListAccounts struct{Username string}
func (ListAccounts) isCommand() {}

// ============================== Transaction Commands ===============================
type CreateTxn struct{
	Username 	string
	AccountName string
	Amount		float64
	Description string
	Date		string
	Income		bool
}
func (CreateTxn) isCommand() {}

type DeleteTxn struct{ID int64}
func (DeleteTxn) isCommand() {}

//needs to have date filters
type LoadAccountTxns struct{
	Username string
	AccountName string
	Uncategorized bool
}
func (LoadAccountTxns) isCommand() {}

type CategorizeTxn struct{ID int64; Category string}
func (CategorizeTxn) isCommand() {}

type UncategorizeTxn struct{ID int64}
func (UncategorizeTxn) isCommand() {}

type SplitTxn struct{ID int64; Splits []string}
func (SplitTxn) isCommand() {}

type ImportTxnsFromFile struct {
	Username 	string
	AccountName string
	FileName	string
	FileType	string
	FileOrigin	string // Bank or place where it came from 
}
func (ImportTxnsFromFile) isCommand() {}

// ============================== Category Commands ===============================
type CreateCategory struct{
	Name string
	IsIncome bool
}
func (CreateCategory) isCommand() {}

type DeleteCategory struct{ID int64}
func (DeleteCategory) isCommand() {}

type LoadAllCategories struct{}
func (LoadAllCategories) isCommand() {}
