package ui

import (
	"fmt"
	"io"
	
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Login Components will always be in default theme
type loginUserItem string
func (i loginUserItem) FilterValue() string { return "" }

type loginUserDelegate struct{
	theme Theme
}

func (d loginUserDelegate) Height() int { return 1 }
func (d loginUserDelegate) Spacing() int { return 0 }
func (d loginUserDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d loginUserDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	it, ok := listItem.(loginUserItem)
	if !ok {
		// maybe should just allow the panic
		return 
	}

	selected := index == m.Index()

	cursor := "  "
	style := d.theme.List.Normal

	if selected {
		cursor = d.theme.List.Cursor.Render("› ")
		style = d.theme.List.Selected
	}

	fmt.Fprint(w, cursor + style.Render(string(it)))
}

func NewUsersList(t Theme, names []string, width, height int) list.Model {
	items := make([]list.Item, len(names))
	for i := range names {
		items[i] = loginUserItem(names[i])
	}

	lst := list.New(items, loginUserDelegate{}, width, height)
	lst.Title = "Registered Users"

	// make minimal
	lst.SetShowHelp(false) 
	lst.SetShowStatusBar(false)
	lst.SetShowFilter(false)
	lst.SetFilteringEnabled(false) 
	
	//Apply theme
	lst.Styles.Title = t.Text.Strong
	return lst
}

func NewLoginInput(t Theme) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "New User"
	ti.CharLimit = 64
	ti.Width = 20 // might be bad to hardcode

	ti.PromptStyle = t.Input.Prompt
	ti.Cursor.Style = t.Input.Cursor
	ti.PlaceholderStyle = t.Input.Placeholder
	ti.TextStyle = t.Input.Text
	return ti
}
