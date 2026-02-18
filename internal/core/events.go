package core

type Event interface{ isEvent() }

type DBFailure struct {Err error}
func (DBFailure) isEvent() {}

type GeneralFailure struct {Err error}
func (GeneralFailure) isEvent() {}

type UsersLoaded struct {Usernames []string}
func (UsersLoaded) isEvent() {}

type SessionLoadFailed struct {Err error}
func (SessionLoadFailed) isEvent() {}

type ActiveUserSet struct{ Username string }
func (ActiveUserSet) isEvent() {}

type ActiveUserCleared struct{}
func (ActiveUserCleared) isEvent() {}

type UserCreated struct{Username string}
func (UserCreated) isEvent() {}

type UserDeleted struct{Username string}
func (UserDeleted) isEvent() {}

type UserSummed struct{Accounts []SummedAccount}
func (UserSummed) isEvent() {}

type SummedCategory struct {
	Name string
	Amount float64
}

type UserTxnsGrouped struct{Groups []SummedCategory}
func (UserTxnsGrouped) isEvent() {}

type EffectUnhandled struct{ Kind string }
func (EffectUnhandled) isEvent() {}

type SummedAccount struct {
	Name    string
	Balance float64
}

// ============================== Account Events ==============================
type AccountCreated struct {
	UserID 		int64
	AccountName string
	Username 	string
	AccountID 	int64
}
func (AccountCreated) isEvent() {}

type AccountDeleted struct {
	UserID 		int64
	AccountName string
	Username 	string
}
func (AccountDeleted) isEvent() {}

type AccountsLoaded struct {AccountNames []string}
func (AccountsLoaded) isEvent() {}

type AccountSummed struct {
	AccountName string
	AccountSum 	float64
}
func (AccountSummed) isEvent() {}

// ============================== Transaction Events ==============================
type Txn struct {
	ID 			int64
	Username 	string
	AccountName string
	Amount		float64
	Description	string
	Income		bool
}

type TxnCreated struct {
	Transaction Txn	
}
func (TxnCreated) isEvent() {}

type TxnDeleted struct {}
func (TxnDeleted) isEvent() {}

type AccountTxnsLoaded struct {
	Transactions []Txn
}
func (AccountTxnsLoaded) isEvent() {}

type TxnsImported struct {
	Transactions []Txn
}
func (TxnsImported) isEvent() {}

type TxnUncategorized struct {ID int64}
func (TxnUncategorized) isEvent() {}

type TxnCategorized struct {
	ID			int64
	Category 	string
}
func (TxnCategorized) isEvent() {}

type TxnSplit struct {
	Old	Txn	
	New [][]string	 // amount, category where category might be ""
}
func (TxnSplit) isEvent() {}

// ============================== Category Events ==============================
type Category struct {
	ID int64
	Name string
	IsIncome bool
}

type CategoryCreated struct{Category Category}
func (CategoryCreated) isEvent() {}

type CategoryDeleted struct{ID int64}
func (CategoryDeleted) isEvent() {}

type CategoriesLoaded struct{Categories []Category}
func (CategoriesLoaded) isEvent() {}
