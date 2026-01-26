package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {	
	Name string

	Text struct {
		Normal lipgloss.Style
		Muted  lipgloss.Style
		Strong lipgloss.Style
		Error  lipgloss.Style
	}	

	Input struct {
		Prompt      lipgloss.Style
		Text        lipgloss.Style
		Placeholder lipgloss.Style
		Cursor      lipgloss.Style
	}	
	
	List struct {
		Normal   lipgloss.Style
		Selected lipgloss.Style
		Dim      lipgloss.Style
		Cursor   lipgloss.Style
	}
}
