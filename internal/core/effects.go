package core 

type Effect interface{ isEffect() }

type FxLoadSession struct{}
func (FxLoadSession) isEffect() {}

type FxLoadAllUsers struct{}
func (FxLoadAllUsers) isEffect() {}

type FxSetSessionUser struct{ Username *string }
func (FxSetSessionUser) isEffect() {}

type FxClearSessionUser struct{}
func (FxClearSessionUser) isEffect() {}

