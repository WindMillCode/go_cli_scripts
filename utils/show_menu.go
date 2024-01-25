package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ShowMenuModel struct {
	Default               string
	Other                 bool
	OtherString           string
	Prompt                string
	Choices               []string       // items on the to-do list
	Selected              map[int]string // which to-do items are selected
	cursor                int            // which to-do list item our cursor is pointing at
	typing                bool           // if the user is typing
	textInput             textinput.Model
	SelectionLimit        int
	SelectedValues        []string
	SelectedDelimiter     string
}

func ShowMenu(cliInfo ShowMenuModel, enableOtherOption interface{}) string {
	cliInfo.Selected = make(map[int]string)
	if cliInfo.SelectedDelimiter ==  ""{
		cliInfo.SelectedDelimiter  = " "
	}
	if cliInfo.OtherString == "" {
		cliInfo.OtherString = "Other: "
	}
	if cliInfo.Other  {
		cliInfo.Choices = append(cliInfo.Choices, cliInfo.OtherString)
	}
	if cliInfo.Default != "" {
		for i, v := range cliInfo.Choices {
			if v == cliInfo.Default {
				cliInfo.Selected[i] = cliInfo.Default
			}
		}
	}
	cliInfo.textInput = textinput.New()
	p := tea.NewProgram(cliInfo)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	value := cliInfo.SelectedDelimiter
	for _, v := range cliInfo.Selected {
		value += v + cliInfo.SelectedDelimiter
		cliInfo.SelectedValues = append(cliInfo.SelectedValues,value)
	}
	strings.Trim(value,cliInfo.SelectedDelimiter)
	return value
}


func (m ShowMenuModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m ShowMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.typing {
		var newCmd tea.Cmd
		if stringMsg, ok := msg.(tea.KeyMsg); ok {
			if stringMsg.String() != "enter" {
				m.textInput, newCmd = m.textInput.Update(msg)
				return m, newCmd
			} else {
				m.textInput.Blur()
				m.Selected[m.cursor] = m.textInput.Value()
				m.typing = false
				return m, nil
			}
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			os.Exit(1)

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.Choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			if _, ok := m.Selected[m.cursor]; ok {
				delete(m.Selected, m.cursor)
			} else if len(m.Selected) < m.SelectionLimit {
				m.Selected[m.cursor] = m.Choices[m.cursor]
			}
			if m.Choices[m.cursor] == m.OtherString && m.Other {
				m.typing = true
				m.textInput.Focus()
				return m, nil
			}
		}
	}

	return m, cmd
}

func (m ShowMenuModel) View() string {
	s := m.Prompt + "\n"
	if m.typing {
		return fmt.Sprintf("Provide a value for OTHER: %s", m.textInput.View())
	}

	for i, choice := range m.Choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	return s
}
