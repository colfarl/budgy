package core

type Event interface{ isEvent() }


type DBFailure struct {Err error}
func (DBFailure) isEvent() {}

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
