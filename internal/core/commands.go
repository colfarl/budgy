package core 

type Command interface{ isCommand() }

type LoadSession struct{}
func (LoadSession) isCommand() {}

type LoadAllUsers struct{}
func (LoadAllUsers) isCommand() {}

type CreateUser struct{ Username string }
func (CreateUser) isCommand() {}

type DeleteUser struct{ Username string }
func (DeleteUser) isCommand() {}

type SetActiveUser struct{ Username string }
func (SetActiveUser) isCommand() {}

type ClearActiveUser struct{}
func (ClearActiveUser) isCommand() {}
