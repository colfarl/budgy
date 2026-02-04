package core

func Transition(s State, e Event) (State) {
	switch v := e.(type) {
	case ActiveUserCleared:
		s.ActiveUser = nil	

	case ActiveUserSet:
		s.ActiveUser = &v.Username
	
	}
	return s
}
