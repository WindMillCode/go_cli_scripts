package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type KillPortsOptions struct {
	Ports          []string
	ProgramName    string
	OutputFile     string
	OpenOutputFile bool
	DryRun         bool
}

func KillPorts(options KillPortsOptions) {
  var findCmdOptions, killCmdOptions CommandOptions

  switch runtime.GOOS {
  case "windows":
    findCmdOptions = CommandOptions{
      Command:   "netstat",
      Args:      []string{"-nao"},
      GetOutput: true,
    }
  case "darwin", "linux", "freebsd", "openbsd", "netbsd", "dragonfly":
    findCmdOptions = CommandOptions{
      Command:   "lsof",
      Args:      []string{"-i", "-P"},
      GetOutput: true,
    }
  case "aix", "solaris", "illumos":
    findCmdOptions = CommandOptions{
      Command:   "netstat",
      Args:      []string{"-an"},
      GetOutput: true,
    }
  default:
    log.Printf("Unsupported OS: %s", runtime.GOOS)
    return
  }

  output, err := RunCommandWithOptions(findCmdOptions)
  if err != nil {
    log.Printf("Error finding processes: %v\n", err)
    return
  }

  pids := make(map[string]bool)
  lines := strings.Split(output, "\n")
	var processesToDelete []string
  for _, line := range lines {
    for _, port := range options.Ports {
      if runtime.GOOS == "windows" {
        if strings.Contains(line, fmt.Sprintf(":%s", port)) {
          fields := strings.Fields(line)
          if len(fields) > 4 {
            pid := fields[len(fields)-1]
            pids[pid] = true
						processesToDelete = append(processesToDelete,line)
          }
        }
      } else {
        if strings.Contains(line, fmt.Sprintf(":%s", port)) {
          fields := strings.Fields(line)
          if len(fields) > 1 {
            pid := fields[1]
            pids[pid] = true
						processesToDelete = append(processesToDelete,line)
          }
        }
      }
    }
  }

  if len(pids) == 0 {
    log.Println("No processes found on the specified ports")
    return
  }

  var pidsToDelete []string
  for pid := range pids {
    pidsToDelete = append(pidsToDelete, pid)
  }

  if options.OutputFile != "" {
    file, err := os.Create(ConvertPathToOSFormat(options.OutputFile))
    if err != nil {
      log.Printf("Failed to create output file: %v\n", err)
      return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    defer writer.Flush()

    writer.WriteString("Matched Processes:\n")
    for _, line := range lines {
      writer.WriteString(line + "\n")
    }

    writer.WriteString("\nProcesses to Delete:\n")
    for _, pid := range processesToDelete {
      writer.WriteString(pid + "\n")
    }

    log.Printf("Process details saved to %s\n", options.OutputFile)
    if options.OpenOutputFile {
      vscodeOpenFileOptions := CommandOptions{
        Command:     "code",
        Args:        []string{options.OutputFile},
        NonBlocking: true,
      }
      RunCommandWithOptions(vscodeOpenFileOptions)
    }
  }

  if options.DryRun {
    return
  }

  var killArgs []string
  if runtime.GOOS == "windows" {
    killArgs = append(killArgs, "/F")
    for _, pid := range pidsToDelete {
      killArgs = append(killArgs, "/PID", pid)
    }
    killCmdOptions = CommandOptions{
      Command: "taskkill",
      Args:    killArgs,
    }
  } else {
    killArgs = append(killArgs, "-9")
    for _, pid := range pidsToDelete {
      killArgs = append(killArgs, pid)
    }
    killCmdOptions = CommandOptions{
      Command: "kill",
      Args:    killArgs,
    }
  }

  _, err = RunCommandWithOptions(killCmdOptions)
  if err != nil {
    log.Printf("Failed to kill processes: %v\n", err)
  } else {
    log.Println("Killed processes on the specified ports")
  }
}


type TakeVariableArgsStruct struct {
	Prompt    string
	ErrMsg    string
	Default   string
	Delimiter string
}

type TakeVariableArgsResultStruct struct {
	InputString string
	InputArray  []string
}

func TakeVariableArgs(obj TakeVariableArgsStruct) TakeVariableArgsResultStruct {
	var innerScriptArguments []string
	prompt0 := obj.Prompt

	if obj.Delimiter == "" {
		obj.Delimiter = " "
	}
	if obj.Default != "" {
		prompt0 = fmt.Sprintf("%s (Default is %s)", obj.Prompt, obj.Default)
	}
	fmt.Println(prompt0)
	fmt.Println("Enter the arguments to pass to the script (press ENTER to enter another argument, leave blank and press ENTER once done):")
	for {
		var argument string
		fmt.Scanln(&argument)

		if strings.TrimSpace(argument) == "" {
			break
		}

		innerScriptArguments = append(innerScriptArguments, argument)
	}
	input := strings.Join(innerScriptArguments, obj.Delimiter)
	if input == "" && obj.ErrMsg != "" {
		panic(obj.ErrMsg)
	} else if input == "" && obj.Default != "" {
		input = obj.Default
		innerScriptArguments = strings.Split(obj.Default, obj.Delimiter)
	}
	return TakeVariableArgsResultStruct{
		InputString: input,
		InputArray:  innerScriptArguments,
	}

}

type GetInputFromStdinStruct struct {
	Prompt  []string
	ErrMsg  string
	Default string
}

func GetInputFromStdin(obj GetInputFromStdinStruct) string {
	if len(obj.Prompt) == 0 {
		obj.Prompt = []string{"Enter your input: "} // Default value
	} else {
		obj.Prompt[0] += " "
	}

	// Create a new scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	if obj.Default != "" {
		fmt.Printf("%s (Default is %s) ", obj.Prompt[0], obj.Default)
	} else {
		fmt.Print(obj.Prompt[0])
	}

	// Read the next line of input from stdin
	if !scanner.Scan() && scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
		return ""
	}
	input := scanner.Text()

	if input == "" && obj.Default != "" {
		input = obj.Default
	} else if input == "" && obj.ErrMsg != "" {
		panic(obj.ErrMsg)
	}

	return input
}

type ShellCommandOutput struct{}

func (c ShellCommandOutput) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func RunCommand(command string, args []string) {

	fullCommand := fmt.Sprintf("Running command: %s %s", command, strings.Join(args, " "))
	fmt.Println(fullCommand)
	cmd := exec.Command(command, args...)
	// cmd.Stdout = ShellCommandOutput{}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {

		msg := fmt.Sprintf("Could not run command %s %s \n This was the err %s", command, strings.Join(args, " "), err.Error())
		fmt.Println(msg)
		// panic(msg)
	}
}

// Deprecated: This function will be removed in the next major release. Use RunCommandWithOptions instead.
func RunCommandInSpecificDirectory(command string, args []string, targetDir string) {

	fullCommand := fmt.Sprintf("Running command: %s %s", command, strings.Join(args, " "))
	fmt.Println(fullCommand)
	cmd := exec.Command(command, args...)
	cmd.Dir = targetDir
	// cmd.Stdout = ShellCommandOutput{}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {

		msg := fmt.Sprintf("Could not run command %s %s \n This was the err %s", command, strings.Join(args, " "), err.Error())
		fmt.Println(msg)
		// panic(msg)
	}
}

// Deprecated: This function will be removed in the next major release. Use RunCommandWithOptions instead.
func RunCommandAndGetOutput(command string, args []string) string {

	fullCommand := fmt.Sprintf("Running command: %s %s", command, strings.Join(args, " "))
	fmt.Println(fullCommand)
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		msg := fmt.Sprintf("Could not run command %s %s \n This was the err %s", command, strings.Join(args, " "), err.Error())
		fmt.Println(msg)
		// panic(msg)

	}
	return string(output)
}

// Deprecated: This function will be removed in the next major release. Use RunCommandWithOptions instead.
func RunCommandInSpecifcDirectoryAndGetOutput(command string, args []string, targetDir string) string {

	fullCommand := fmt.Sprintf("Running command: %s %s", command, strings.Join(args, " "))
	fmt.Println(fullCommand)
	cmd := exec.Command(command, args...)
	cmd.Dir = targetDir
	output, err := cmd.Output()
	if err != nil {
		msg := fmt.Sprintf("Could not run command %s %s \n This was the err %s", command, strings.Join(args, " "), err.Error())
		fmt.Println(msg)
		// panic(msg)

	}
	return string(output)
}

type CommandOptions struct {
	CmdObj             *exec.Cmd
	Self               *CommandOptions
	Command            string
	Args               []string
	TargetDir          string
	GetOutput          bool
	// TODO PrintOutput does not play well it says nothing is on the path
	PrintOutput        bool
	PrintOutputOnly    bool
	PanicOnError       bool
	NonBlocking        bool
	IsInputFromProgram bool
	IsElevated         bool
	EnvVars            map[string]string
}

func (c CommandOptions) EndProcess() error {
	var cmd *exec.Cmd
	if c.Self != nil && c.Self.CmdObj != nil {
		cmd = c.Self.CmdObj
	} else {
		cmd = c.CmdObj
	}
	if cmd != nil {
		return cmd.Process.Kill()

	}
	return nil
}

type DualWriter struct {
	TerminalWriter io.Writer
	Buffer         *bytes.Buffer
}

func (w DualWriter) Write(p []byte) (n int, err error) {
	n, err = w.TerminalWriter.Write(p)
	if err != nil {
		return n, err
	}

	// Write to the buffer as well
	bufferBytes, bufferErr := w.Buffer.Write(p)
	if bufferErr != nil {
		return bufferBytes, bufferErr
	}

	return n, nil
}

func RunCommandWithOptions(options CommandOptions) (string, error) {

	if options.IsElevated {
		return "", RunElevatedCommand(options.Command, options.Args)
	}
	fullCommand := fmt.Sprintf("Running command: %s %s\n", options.Command, strings.Join(options.Args, " "))
	fmt.Println(fullCommand)

	cmd := exec.Command(options.Command, options.Args...)
	if options.Self != nil {
		options.Self.CmdObj = cmd
	}
	if options.TargetDir != "" {
		cmd.Dir = options.TargetDir
	}

	if options.IsInputFromProgram != true {
		cmd.Stdin = os.Stdin
	}

	if options.EnvVars != nil {
		for key, value := range options.EnvVars {
			os.Setenv(key,value)
			// cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
		}
		cmd.Env =nil
	}

	// Creating buffers and DualWriters for stdout and stderr
	var stdoutBuffer, stderrBuffer bytes.Buffer
	stdoutWriter := DualWriter{TerminalWriter: os.Stdout, Buffer: &stdoutBuffer}
	stderrWriter := DualWriter{TerminalWriter: os.Stderr, Buffer: &stderrBuffer}

	cmd.Stdout = stdoutWriter
	if options.PrintOutput == false {
		cmd.Stdout = &stdoutBuffer
	}
	if options.PrintOutputOnly == true {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = stderrWriter

	var err error
	if options.NonBlocking {
		err = cmd.Start() // Non-blocking execution if NonBlocking is true
	} else {
		err = cmd.Run() // Default to blocking execution
	}

	if err != nil {
		// Construct error message
		msg := fmt.Sprintf(
			"Could not run command %s %s\n\nThis was the err: %s \n %s\n\n",
			options.Command,
			strings.Join(options.Args, " "),
			err.Error(),
			fmt.Sprintf("Standard Error: %s\n", stderrBuffer.String()),
		)
		fmt.Println(msg)

		if options.PanicOnError {
			panic(msg)
		}

		return "", err
	}

	if options.GetOutput {
		return stdoutBuffer.String(), nil
	}

	return "", nil
}

func RunElevatedCommand(command string, args []string) error {
	var elevatedCommand string
	var elevatedArgs []string

	switch runtime.GOOS {

	case "windows":
		elevatedCommand = "powershell"
		elevatedArgs = []string{"-Command", fmt.Sprintf(`Start-Process cmd -ArgumentList '/c %s %s' -Verb RunAs -Wait`, command, JoinArgs(args))}
	case "darwin", "linux":
		elevatedCommand = "sudo"
		elevatedArgs = append([]string{command}, args...)
	default:
		return fmt.Errorf("unsupported platform")
	}

	options := CommandOptions{
		Command:     elevatedCommand,
		Args:        elevatedArgs,
		GetOutput:   true,
		PrintOutput: true,
		NonBlocking: false,
	}

	_, err := RunCommandWithOptions(options)
	return err
}
