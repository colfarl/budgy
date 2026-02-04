package core

type State struct {
	ActiveUser *string
	Error		error
}

func NewState() State {
	return State{
		ActiveUser: nil,
		Error: nil,
	}
}

