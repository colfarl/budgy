package app

func (m Model) View() string {
	switch {
	case m.Logged_in:
		return "logged in"
	default:
		return "not logged in"
	}
}
