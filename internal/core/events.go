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

type EffectUnhandled struct{ Kind string }
func (EffectUnhandled) isEvent() {}

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
