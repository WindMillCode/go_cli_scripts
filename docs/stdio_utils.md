## [TakeVariableArgsStruct]
### Reference
| Field      | Type   | Description                                        |
|------------|--------|----------------------------------------------------|
| Prompt     | string | The message displayed to the user for input.       |
| ErrMsg     | string | The error message displayed if input validation fails. |
| Default    | string | The default value if no input is provided.         |
| Delimiter  | string | The delimiter used to separate multiple inputs.    |

## [TakeVariableArgsResultStruct]
### Reference
| Field        | Type     | Description                                      |
|--------------|----------|--------------------------------------------------|
| InputString  | string   | The raw input string collected from the user.    |
| InputArray   | []string | The input split into an array based on Delimiter.|

## [GetInputFromStdinStruct]
### Reference
| Field    | Type     | Description                                       |
|----------|----------|---------------------------------------------------|
| Prompt   | []string | An array where the first element is the input prompt. |
| ErrMsg   | string   | Error message displayed if input is invalid.      |
| Default  | string   | Default value returned if no input is provided.   |

## [CommandOptions]
### Reference
| Field                | Type      | Description                                          |
|----------------------|-----------|------------------------------------------------------|
| CmdObj               | *exec.Cmd | The command object, used internally.                 |
| Self                 | *CommandOptions | Reference to itself, used for chaining.           |
| Command              | string    | The command to execute.                              |
| Args                 | []string  | Arguments for the command.                           |
| TargetDir            | string    | Directory where the command will be executed.        |
| GetOutput            | bool      | If true, captures and returns the command output.    |
| PrintOutput          | bool      | If true, prints the command output to stdout.        |
| PrintOutputOnly      | bool      | If true, only prints the output without capturing.   |
| PanicOnError         | bool      | If true, panics on command execution error.          |
| NonBlocking          | bool      | If true, executes the command without blocking.      |
| IsInputFromProgram   | bool      | If true, indicates the command input is from the program. |
