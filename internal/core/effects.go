package core 

type Effect interface{ isEffect() }

// ============================== User Effects ==============================
type FxLoadSession struct{}
func (FxLoadSession) isEffect() {}

type FxLoadAllUsers struct{}
func (FxLoadAllUsers) isEffect() {}

type FxSetSessionUser struct{ Username *string }
func (FxSetSessionUser) isEffect() {}

type FxCreateUser struct{ Username *string }
func (FxCreateUser) isEffect() {}

type FxGetUserBalances struct{ Username string }
func (FxGetUserBalances) isEffect() {}

type FxDeleteUser struct{ Username *string }
func (FxDeleteUser) isEffect() {}

type FxClearSessionUser struct{}
func (FxClearSessionUser) isEffect() {}

// ============================== Account Effects ==============================
type FxCreateAccount struct{
	Username 	string
	AccountName string
}
func (FxCreateAccount) isEffect() {}

type FxGetAccountBalance struct{
	Username 	string
	AccountName string
}
func (FxGetAccountBalance) isEffect() {}


type FxDeleteAccount struct{
	Username 	string
	AccountName string
}
func (FxDeleteAccount) isEffect() {}

type FxLoadAccounts struct {Username string}
func (FxLoadAccounts) isEffect() {}

// ============================== Transaction Effects ==============================
type FxCreateTxn struct{
	Username 	string
	AccountName string
	Amount		float64
	Description	string
	Date		int64
	Income		bool
}
func (FxCreateTxn) isEffect() {}

type FxDeleteTxn struct{TxnID int64}
func (FxDeleteTxn) isEffect() {}

type FxLoadAccountTxns struct{
	Username 		string
	AccountName 	string
	Uncategorized 	bool
}
func (FxLoadAccountTxns) isEffect() {}

type FxCategorizeTxn struct{
	ID 			int64	
	Category 	string
}
func (FxCategorizeTxn) isEffect() {}

type FxUncategorizeTxn struct{ID int64}
func (FxUncategorizeTxn) isEffect() {}

type FxSplitTransaction struct{
	ID 		int64
	Splits 	[]string
}
func (FxSplitTransaction) isEffect() {}

type FxImportTxnsFromFile struct{
	Username 	string
	AccountName string
	FilePath	string	
	FileOrigin	string	
	FileType	string 	
}
func (FxImportTxnsFromFile) isEffect() {}

// ============================== Category Effects ==============================
type FxCreateCategory struct{Category Category}
func (FxCreateCategory) isEffect() {}

type FxDeleteCategory struct{ID int64}
func (FxDeleteCategory) isEffect() {}

type FxLoadCategories struct{}
func (FxLoadCategories) isEffect() {}
