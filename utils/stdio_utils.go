package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"runtime"
	"strings"
	"syscall"
)

type SetGlobalVarsOptions struct {
	NonInteractive ProcessIfDefaultIsPresentStruct
	Infinity       int
}

var GLOBAL_VARS SetGlobalVarsOptions

func SetGlobalVars(options SetGlobalVarsOptions) {
	GLOBAL_VARS = options
	GLOBAL_VARS.Infinity = 1<<31 - 1
}

type KillPortsOptions struct {
	Ports          []string
	ProgramNames   []string
	OutputFile     string
	OpenOutputFile bool
	DryRun         bool
}

type KillPortsProcessInfo struct {
	ColumnNameIndex int
	PIDIndex        int
	Output          string
	Columns         []string
	Lines           []string
	Rows            []map[string]string
	Regex           *regexp.Regexp
}

func getColumnIndex(headers []string, columnName string) int {
	for i, header := range headers {
		if strings.Contains(strings.ToLower(header), strings.ToLower(columnName)) {
			return i
		}
	}
	return -1
}

func findPIDIndex(info *KillPortsProcessInfo) {
	for i, line := range info.Lines {
		if strings.Contains(line, "PID") {
			fields := info.Regex.Split(line, -1)
			info.ColumnNameIndex = i
			for j, field := range fields {
				if strings.Contains(field, "PID") {
					info.PIDIndex = j
					return
				}
			}
		}
	}
	info.PIDIndex = -1 // set to -1 if "PID" not found
}

func initProcessInfo(processInfo *KillPortsProcessInfo) {
	processInfo.Columns = processInfo.Regex.Split(strings.TrimSpace(processInfo.Lines[processInfo.ColumnNameIndex]), -1)
	for i, line := range processInfo.Lines {

		if i <= processInfo.ColumnNameIndex {
			continue
		}
		columns := processInfo.Columns
		var fields = processInfo.Regex.Split(strings.TrimSpace(line), -1)
		processMap := make(map[string]string)

		for i, column := range columns {
			if i < len(fields) {
				// TODO instead of the key use an additional number map?
				processMap[column] = fields[i]
			} else {
				processMap[column] = ""
			}
		}

		processInfo.Rows = append(processInfo.Rows, processMap)

	}
}

func KillPorts(options KillPortsOptions) {
	var findProcessOptions, findNameOptions, killCmdOptions CommandOptions

	switch runtime.GOOS {
	case "windows":
		netstatPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "netstat.ps1"))
		if err != nil {
			fmt.Printf("Failed to extract netstat.ps1: %v", err)
		}
		tasklistPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "tasklist.ps1"))
		if err != nil {
			fmt.Printf("Failed to extract tasklist.ps1: %v", err)
		}
		findProcessOptions = CommandOptions{
			Command:   "powershell",
			Args:      []string{"-c", netstatPath},
			GetOutput: true,
		}
		findNameOptions = CommandOptions{
			Command:   "powershell",
			Args:      []string{"-c", tasklistPath},
			GetOutput: true,
		}
	case "darwin":
		netstatPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "netstat_macos.sh"))
		if err != nil {
			fmt.Printf("Failed to extract netstat_macos.sh: %v", err)
		}
		psPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "ps_macos.sh"))
		if err != nil {
			fmt.Printf("Failed to extract ps_macos.sh: %v", err)
		}
		findProcessOptions = CommandOptions{
			Command:   "sh",
			Args:      []string{netstatPath},
			GetOutput: true,
		}
		findNameOptions = CommandOptions{
			Command:   "sh",
			Args:      []string{psPath},
			GetOutput: true,
		}
	case "linux":
		netstatPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "netstat_linux.sh"))
		if err != nil {
			fmt.Printf("Failed to extract netstat_linux.sh: %v", err)
		}
		psPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "ps_linux.sh"))
		if err != nil {
			fmt.Printf("Failed to extract ps_linux.sh: %v", err)
		}
		findProcessOptions = CommandOptions{
			Command:   "sh",
			Args:      []string{netstatPath},
			GetOutput: true,
		}
		findNameOptions = CommandOptions{
			Command:   "sh",
			Args:      []string{psPath},
			GetOutput: true,
		}
	case "freebsd", "openbsd", "netbsd", "dragonfly":
		netstatPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "netstat_bsd.sh"))
		if err != nil {
			fmt.Printf("Failed to extract netstat_bsd.sh: %v", err)
		}
		psPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "ps_bsd.sh"))
		if err != nil {
			fmt.Printf("Failed to extract ps_bsd.sh: %v", err)
		}
		findProcessOptions = CommandOptions{
			Command:   "sh",
			Args:      []string{netstatPath},
			GetOutput: true,
		}
		findNameOptions = CommandOptions{
			Command:   "sh",
			Args:      []string{psPath},
			GetOutput: true,
		}
	case "aix", "solaris", "illumos":
		netstatPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "netstat_unix.sh"))
		if err != nil {
			fmt.Printf("Failed to extract netstat_unix.sh: %v", err)
		}
		psPath, err := GetFilePathFromPackage(JoinAndConvertPathToOSFormat("scripts", "ps_unix.sh"))
		if err != nil {
			fmt.Printf("Failed to extract ps_unix.sh: %v", err)
		}
		findProcessOptions = CommandOptions{
			Command:   "sh",
			Args:      []string{netstatPath},
			GetOutput: true,
		}
		findNameOptions = CommandOptions{
			Command:   "sh",
			Args:      []string{psPath},
			GetOutput: true,
		}
	default:
		fmt.Printf("Unsupported OS: %s", runtime.GOOS)
		return
	}

	findProcessOutput, err := RunCommandWithOptions(findProcessOptions)
	if err != nil {
		fmt.Printf("Error finding processes: %v\n", err)
		return
	}

	findNameOutput, err := RunCommandWithOptions(findNameOptions)
	if err != nil {
		fmt.Printf("Error finding processes: %v\n", err)
		return
	}

	infoFindProcess := KillPortsProcessInfo{
		PIDIndex: -1,
		Output:   findProcessOutput,
		Lines:    strings.Split(findProcessOutput, "\n"),
		Regex:    regexp.MustCompile(`\s{2,}`),
	}

	infoFindName := KillPortsProcessInfo{
		PIDIndex: -1,
		Output:   findNameOutput,
		Lines:    strings.Split(findNameOutput, "\n"),
		Regex:    regexp.MustCompile(`\s{2,}`),
	}

	switch runtime.GOOS {
	case "darwin":
		infoFindProcess.Regex = regexp.MustCompile(`\s{2,}|->`)
	}

	findPIDIndex(&infoFindProcess)

	findPIDIndex(&infoFindName)

	processMap := make(map[string]string)
	for _, line := range strings.Split(findNameOutput, "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			pid := fields[0]
			name := fields[1]
			processMap[pid] = name
		}
	}

	initProcessInfo(&infoFindProcess)
	initProcessInfo(&infoFindName)
	var rowsWithPIDAndName []map[string]string

	for _, row := range infoFindProcess.Rows {

		var nameRow map[string]string
		for _, r := range infoFindName.Rows {
			if r["PID"] == row["PID"] {
				nameRow = r
				break
			}
		}
		unionRow := make(map[string]string)

		for k, v := range row {
			unionRow[k] = v
		}

		for k, v := range nameRow {
			unionRow[k] = v
		}
		rowsWithPIDAndName = append(rowsWithPIDAndName, unionRow)
	}

	var pidsToDelete []string
	for _, row := range rowsWithPIDAndName {
		for _, port := range options.Ports {
			formattedPort := fmt.Sprintf(":%s", port)
			isInTargetProgramNames := true
			if len(options.ProgramNames) != 0 {
				isInTargetProgramNames = ArrayContainsAny(options.ProgramNames, []string{row["Name"]})
			}
			foreignAddressContainsPort := strings.Contains(row["Foreign Address"], formattedPort)
			localAddressContainsPort := strings.Contains(row["Local Address"], formattedPort)
			pidIsNotZero := row["PID"] != "0"
			stateIsListen := strings.Contains(strings.ToLower(row["State"]), "listen")
			stateIsTimeWait := strings.Contains(strings.ToLower(row["State"]), "timewait")

			switch runtime.GOOS {
			// TODO see if you need listening
			case "windows":

				cantKillByName := ArrayContainsAny([]string{"svchost", "SearchHost"}, []string{row["Name"]})

				if ((foreignAddressContainsPort || localAddressContainsPort) && stateIsListen || (stateIsTimeWait && pidIsNotZero && isInTargetProgramNames)) && !cantKillByName {
					pid := row["PID"]
					pidsToDelete = append(pidsToDelete, pid)
				}
			case "darwin":

				cantKillByName := ArrayContainsAny([]string{"svchost", "SearchHost"}, []string{row["Name"]})
				if ((foreignAddressContainsPort || localAddressContainsPort) && stateIsListen || (stateIsTimeWait && pidIsNotZero && isInTargetProgramNames)) && !cantKillByName {
					pid := row["PID"]
					pidsToDelete = append(pidsToDelete, pid)
				}

			default:
				fmt.Printf("Unsupported OS: %s", runtime.GOOS)
				return
			}
		}
	}

	pidsToDelete = RemoveDuplicates(pidsToDelete)
	if options.OutputFile != "" {
		file, err := os.Create(ConvertPathToOSFormat(options.OutputFile))
		if err != nil {
			fmt.Printf("Failed to create output file: %v\n", err)
			return
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush()

		firstRow := rowsWithPIDAndName[0]
		columnNames := make([]string, 0, len(firstRow))
		for key := range firstRow {
			columnNames = append(columnNames, key)
		}

		// Write column names
		writer.WriteString(strings.Join(columnNames, ",") + "\n")

		// Write each row
		for _, row := range rowsWithPIDAndName {
			values := make([]string, len(columnNames))
			for i, columnName := range columnNames {
				values[i] = row[columnName]
			}
			writer.WriteString(strings.Join(values, ",") + "\n")
		}

		var deleteProcessHeader = strings.Join(
			append([]string{"Processes", "to", "Delete:"}, strings.Split(strings.Repeat("==== ", len(columnNames)-1), " ")...), ",",
		) + "\n"
		writer.WriteString(deleteProcessHeader)
		writer.WriteString(strings.Join(columnNames, ",") + "\n")

		for _, row := range rowsWithPIDAndName {
			for _, pid := range pidsToDelete {
				if row["PID"] == pid {
					values := make([]string, len(columnNames))
					for i, columnName := range columnNames {
						values[i] = row[columnName]
					}
					writer.WriteString(strings.Join(values, ",") + "\n")
				}
			}
		}

		fmt.Printf("Process details saved to %s\n", options.OutputFile)
		if options.OpenOutputFile {
			vscodeOpenFileOptions := CommandOptions{
				Command:     "code",
				Args:        []string{options.OutputFile},
				NonBlocking: true,
			}
			RunCommandWithOptions(vscodeOpenFileOptions)
		}
	}

	if len(pidsToDelete) == 0 {
		fmt.Println("No processes found on the specified ports")
		return
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
		killArgs = append(killArgs, pidsToDelete...)

		killCmdOptions = CommandOptions{
			Command: "kill",
			Args:    killArgs,
		}
	}

	_, err = RunCommandWithOptions(killCmdOptions)
	if err != nil {
		fmt.Printf("Failed to kill processes: %v\n", err)
	} else {
		fmt.Println("Killed processes on the specified ports")
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

	var input string
	if obj.Default != "" && GLOBAL_VARS.NonInteractive.Global {
		input = obj.Default
		innerScriptArguments = strings.Split(input, obj.Delimiter)
	} else {
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
	}
	input = strings.Join(innerScriptArguments, obj.Delimiter)
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

	if len(obj.Default) != 0 && GLOBAL_VARS.NonInteractive.Global {
		return obj.Default
	}
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

// Deprecated: Its recommended to use RunCommandWithOptions instead. Wont be moved anytime soon
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

// Deprecated: Its recommended to use RunCommandWithOptions instead. Wont be moved anytime soon
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

// Deprecated: Its recommended to use RunCommandWithOptions instead. Wont be moved anytime soon
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

type CommandOptions struct {
	CmdObj             *exec.Cmd
	Self               *CommandOptions
	Command            string
	Args               []string
	TargetDir          string
	GetOutput          bool
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
			os.Setenv(key, value)
			// cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
		}
		cmd.Env = nil
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

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		if cmd.Process != nil {
			cmd.Process.Signal(sig)
		}
	}()

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

// Deprecated: Its recommended to use RunCommandWithOptions instead. Wont be moved anytime soon
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
