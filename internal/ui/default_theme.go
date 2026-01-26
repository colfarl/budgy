package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func DefaultTheme() *Theme {
	var t Theme
	
	// AI generated colorscheme, play around with this
	var (
		fg       = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
		muted    = lipgloss.AdaptiveColor{Light: "245", Dark: "243"}
		//border   = lipgloss.AdaptiveColor{Light: "250", Dark: "238"}
		accent   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
		errorCol = lipgloss.AdaptiveColor{Light: "#D70000", Dark: "#FF5F5F"}
		//bgSoft   = lipgloss.AdaptiveColor{Light: "254", Dark: "236"}
	)
	
	// text
	t.Text.Normal = lipgloss.NewStyle().Foreground(fg)
	t.Text.Muted = lipgloss.NewStyle().Foreground(muted)
	t.Text.Strong = lipgloss.NewStyle().Foreground(fg).Bold(true)
	t.Text.Error = lipgloss.NewStyle().Foreground(errorCol).Bold(true)
	
	// inputs
	t.Input.Prompt = lipgloss.NewStyle().
		Foreground(accent).
		Bold(true)

	t.Input.Text = lipgloss.NewStyle().
		Foreground(fg)

	t.Input.Placeholder = lipgloss.NewStyle().
		Foreground(muted)

	t.Input.Cursor = lipgloss.NewStyle().
		Reverse(true)
	
	//lists
	t.List.Normal = lipgloss.NewStyle().Foreground(fg)
	t.List.Dim = lipgloss.NewStyle().Foreground(muted)

	t.List.Selected = lipgloss.NewStyle().
		Foreground(accent).
		Bold(true)

	t.List.Cursor = lipgloss.NewStyle().
		Foreground(accent).
		SetString("› ")

	return &t
}
