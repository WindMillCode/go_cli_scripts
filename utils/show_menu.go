package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ShowMenuModel struct {
	Default            string
	Other              bool
	OtherString        string
	Prompt             string
	Choices            []string       // items on the to-do list
	Selected           map[int]string // which to-do items are selected
	cursor             int            // which to-do list item our cursor is pointing at
	typing             bool           // if the user is typing
	textInput          textinput.Model
	// TODO implement per coommand nonintreactive
	nonInteractive      bool
}

func ShowMenu(cliInfo ShowMenuModel, enableOtherOption interface{}) string {
	cliInfo.Selected = make(map[int]string)

	if cliInfo.Default != "" && GLOBAL_VARS.NonInteractive.Global {
		return cliInfo.Default
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
	var value string
	for _, v := range cliInfo.Selected {
		value = v
		break
	}
	return value
}

func (m ShowMenuModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m ShowMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd = nil
	if m.typing {
		var newCmd tea.Cmd
		if  stringMsg,ok := msg.(tea.KeyMsg); ok{
			if stringMsg.String() != "enter"{
				m.textInput,newCmd = m.textInput.Update(msg)
				cmd = newCmd
				return m,cmd
			} else{
				m.textInput.Blur()
				m.Selected[m.cursor] = m.textInput.Value()
				return m, tea.Quit
			}

		}
	}
	switch msg := msg.(type) {

		// Is it a key press?
		case tea.KeyMsg:

			// Cool, what was the actual key pressed?
			switch msg.String() {

			// These keys should exit the program.
			case "ctrl+c", "q":
				os.Exit(1)


			// The "up" and "k" keys move the cursor up
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			// The "down" and "j" keys move the cursor down
			case "down", "j":
				if m.cursor < len(m.Choices)-1 {
					m.cursor++
				}

			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
			case "enter", " ":

				for key := range m.Selected {
					delete(m.Selected, key)
				}
				choice := m.Choices[m.cursor]
				m.Selected[m.cursor] = choice
				if choice == m.OtherString && m.Other  {
					m.typing = true
					m.textInput.Focus()
					return m, nil
				}
				return m, tea.Quit
			}


	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}

func (m ShowMenuModel) View() string {
	// The header
	s := m.Prompt + "\n"
	if m.typing{
		return  fmt.Sprintf("Provide a value for OTHER: %s ",m.textInput.View())
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

	// The footer
	// s += "\nPress q to confirm selection.\n"

	// Send the UI for rendering
	return s
}


type ShowMenuMultipleModel struct {
	Defaults            []string          // Default selections
	Prompt              string
	Choices             []string          // Items on the list
	Selected            map[int]string    // Selected items
	cursor              int               // Cursor position
	typing              bool              // If the user is typing
	textInput           textinput.Model   // Text input model
	Other               bool              // Enable 'Other' option
	OtherString         string            // String for 'Other' option
	SelectionLimit      int               // Maximum number of selections allowed
	SelectedDelimiter   string            // Delimiter for selected items in the returned string
}

func ShowMenuMultipleOptions(cliInfo ShowMenuMultipleModel, enableOtherOption interface{}) []string {

	if len(cliInfo.Defaults) != 0 && GLOBAL_VARS.NonInteractive.Global {
		return cliInfo.Defaults
	}
	cliInfo.Selected = make(map[int]string)
	cliInfo.textInput = textinput.New()

	if cliInfo.OtherString == "" {
		cliInfo.OtherString = "Other: "
	}
	if cliInfo.Other {
		cliInfo.Choices = append(cliInfo.Choices, cliInfo.OtherString)
	}

	// Mark defaults as selected
	for _, defaultChoice := range cliInfo.Defaults {
		for i, choice := range cliInfo.Choices {
			if choice == defaultChoice {
				cliInfo.Selected[i] = choice
			}
		}
	}

	p := tea.NewProgram(cliInfo)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// Formatting the selected items
	var selectedItems []string
	for _, item := range cliInfo.Selected {
		selectedItems = append(selectedItems, item)
	}

	return selectedItems
}

func (m ShowMenuMultipleModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m ShowMenuMultipleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.typing {
		var newCmd tea.Cmd
		m.textInput, newCmd = m.textInput.Update(msg)
		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "enter" {
			m.textInput.Blur()
			m.Selected[m.cursor] = m.textInput.Value()
			m.typing = false
			return m, tea.Quit
		}
		cmd = newCmd
		return m, cmd
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
		case "enter":
			// Toggle selection
			choice := m.Choices[m.cursor]
			if _, ok := m.Selected[m.cursor]; ok {
				delete(m.Selected, m.cursor)
			} else {
				if m.SelectionLimit == 0 || len(m.Selected) < m.SelectionLimit {
					m.Selected[m.cursor] = choice
				}
			}
			if choice == m.OtherString && m.Other {
				m.typing = true
				m.textInput.Focus()
				return m, nil
			}
			return m, nil
		case " ":
			// Return the selected items when spacebar is pressed
			return m, tea.Quit
		}
	}
	return m, cmd
}


func (m ShowMenuMultipleModel) View() string {
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
	s += fmt.Sprintf("\n Hit SPACE or SPACEBAR when done")
	return s
}
