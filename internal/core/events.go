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

type EffectUnhandled struct{ Kind string }
func (EffectUnhandled) isEvent() {}
