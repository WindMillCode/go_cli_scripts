## [ShowMenuModel]
### Usage
`ShowMenuModel` is used to create an interactive menu using Bubble Tea, which allows users to make a single selection from a list of choices.

### Fields in `ShowMenuModel`

| Field         | Type                  | Description                                               |
|---------------|-----------------------|-----------------------------------------------------------|
| Default       | string                | The default selected item in the menu.                    |
| Other         | bool                  | Flag to include an "Other" option for custom input.       |
| OtherString   | string                | The string displayed for the "Other" option.              |
| Prompt        | string                | The prompt message displayed above the choices.           |
| Choices       | []string              | The list of choices available in the menu.                |
| Selected      | map[int]string        | The map tracking which items are selected.                |
| cursor        | int                   | The cursor's current position in the choice list.         |
| typing        | bool                  | Flag indicating if the user is currently typing.          |
| textInput     | textinput.Model       | The model for text input when "Other" is selected.        |

## [ShowMenu]
### Usage
`ShowMenu` displays an interactive menu using the `ShowMenuModel` and returns the user's selection as a string.

```go
menu := ShowMenuModel{
    Prompt: "Select an option:",
    Choices: []string{"Option 1", "Option 2"},
    Other: true,
}
selected := ShowMenu(menu, nil)
fmt.Println("Selected:", selected)
```

## [ShowMenuMultipleModel]
### Usage
`ShowMenuMultipleModel` is similar to `ShowMenuModel` but allows for multiple selections.

### Fields in `ShowMenuMultipleModel`

| Field              | Type                  | Description                                                   |
|--------------------|-----------------------|---------------------------------------------------------------|
| Defaults           | []string              | The default selections.                                       |
| Prompt             | string                | The prompt message displayed above the choices.               |
| Choices            | []string              | The list of choices available in the menu.                    |
| Selected           | map[int]string        | The map tracking which items are selected.                    |
| cursor             | int                   | The cursor's current position in the choice list.             |
| typing             | bool                  | Flag indicating if the user is currently typing.              |
| textInput          | textinput.Model       | The model for text input when "Other" is selected.            |
| Other              | bool                  | Flag to include an "Other" option for custom input.           |
| OtherString        | string                | The string displayed for the "Other" option.                  |
| SelectionLimit     | int                   | The maximum number of selections allowed.                     |
| SelectedDelimiter  | string                | The delimiter for selected items in the returned string.      |

## [ShowMenuMultipleOptions]
### Usage
`ShowMenuMultipleOptions` displays an interactive menu for multiple selections using the `ShowMenuMultipleModel`.

```go
multiMenu := ShowMenuMultipleModel{
    Prompt: "Select options:",
    Choices: []string{"Option 1", "Option 2", "Option 3"},
    Other: true,
}
selectedItems := ShowMenuMultipleOptions(multiMenu, nil)
fmt.Println("Selected items:", selectedItems)
```

These menu models and functions create interactive, terminal-based menus for single or multiple selections, with support for custom user input.
