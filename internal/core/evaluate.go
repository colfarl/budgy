package core

func Evaluate (s State, c Command) ([]Event, []Effect) {
	switch cmd := c.(type) {

	case LoadSession:
		return evalLoadSession(s)

	case SetActiveUser:
		u := cmd.Username
		return evalSetActiveUser(s, u)

	case ClearActiveUser:
		return evalClearActiveUser(s)		

	case LoadAllUsers:
		return nil, []Effect{FxLoadAllUsers{}}
	}
	return nil, nil
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
	fxs := []Effect{FxSetSessionUser{Username: &username}}
	return evs, fxs
}

func evalLoadSession(s State) ([]Event, []Effect) {
	if s.ActiveUser != nil {
		return nil, nil
	}
	return nil, []Effect{FxLoadSession{}}
}

