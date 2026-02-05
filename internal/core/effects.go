package core 

type Effect interface{ isEffect() }

type FxLoadSession struct{}
func (FxLoadSession) isEffect() {}

type FxLoadAllUsers struct{}
func (FxLoadAllUsers) isEffect() {}

type FxSetSessionUser struct{ Username *string }
func (FxSetSessionUser) isEffect() {}

type FxCreateUser struct{ Username *string }
func (FxCreateUser) isEffect() {}

type FxDeleteUser struct{ Username *string }
func (FxDeleteUser) isEffect() {}

type FxClearSessionUser struct{}
func (FxClearSessionUser) isEffect() {}
