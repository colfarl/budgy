package core

import "log"

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
		log.Printf("Called Load Users")
		return nil, []Effect{FxLoadAllUsers{}}
	
	case DeleteUser:
		log.Printf("Evaluated DeleteUser")
		return nil, []Effect{FxDeleteUser{Username: &cmd.Username}}
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
	fxs := []Effect{
		FxCreateUser{Username: &username}, // this shouldn't have to be done every time we set active user
		FxSetSessionUser{Username: &username},
	}
	return evs, fxs
}

func evalLoadSession(s State) ([]Event, []Effect) {
	if s.ActiveUser != nil {
		return nil, nil
	}
	return nil, []Effect{FxLoadSession{}}
}

