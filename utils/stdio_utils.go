package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type TakeVariableArgsStruct struct{
	Prompt string
	ErrMsg string
	Default string
	Delimiter string
}

func TakeVariableArgs(obj TakeVariableArgsStruct) string {
	var innerScriptArguments []string
	prompt0 := obj.Prompt
	if(obj.Delimiter == ""){
		obj.Delimiter = " "
	}
	if(obj.Default != ""){
		prompt0 =fmt.Sprintf("%s (Default is %s)",obj.Prompt,obj.Default)
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
	input := strings.Join(innerScriptArguments,obj.Delimiter)
	if(input == "" && obj.ErrMsg != ""){
		panic(obj.ErrMsg)
	} else if (input == "" && obj.Default !=""){
		input = obj.Default
	}
	return input
}

type GetInputFromStdinStruct struct{
	Prompt []string
	ErrMsg string
	Default string
}

func GetInputFromStdin(obj GetInputFromStdinStruct) string {
	if len(obj.Prompt) == 0 {
		obj.Prompt = []string{"Enter your input: "} // Default value
	} else  {
		obj.Prompt[0] += " "
	}
	// Create a new scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	if obj.Default != "" {
		fmt.Print(fmt.Sprintf("%s (Default is %s) ",obj.Prompt[0] , obj.Default))
	} else {
		fmt.Print(fmt.Sprintf("%s",obj.Prompt[0]))
	}

	// Read the next line of input from stdin
	scanner.Scan()
	input := scanner.Text()
	if (input == "" && obj.Default != ""){
		input = obj.Default
	} else if(input == "" && obj.ErrMsg != ""){
		panic(obj.ErrMsg)
	}

	return input
}

type ShellCommandOutput struct{}

func (c ShellCommandOutput) Write(p []byte) (int, error) {
	fmt.Println( string(p))
	return len(p), nil
}

func RunCommand(command string,args []string) {

	fullCommand :=  fmt.Sprintf("Running command: %s %s", command,strings.Join(args," "))
	fmt.Println(fullCommand)
	cmd := exec.Command(command, args...)
	// cmd.Stdout = ShellCommandOutput{}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {

		msg := fmt.Sprintf("Could not run command %s %s \n This was the err %s", command,strings.Join(args," "),err.Error())
		fmt.Println(msg)
		// panic(msg)
	}
}

func RunCommandInSpecificDirectory(command string,args []string,targetDir string) {

	fullCommand :=  fmt.Sprintf("Running command: %s %s", command,strings.Join(args," "))
	fmt.Println(fullCommand)
	cmd := exec.Command(command, args...)
	cmd.Dir = targetDir
	// cmd.Stdout = ShellCommandOutput{}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {

		msg := fmt.Sprintf("Could not run command %s %s \n This was the err %s", command,strings.Join(args," "),err.Error())
		fmt.Println(msg)
		// panic(msg)
	}
}

func RunCommandAndGetOutput(command string,args []string) string {

	fullCommand :=  fmt.Sprintf("Running command: %s %s", command,strings.Join(args," "))
	fmt.Println(fullCommand)
	output,err := exec.Command(command, args...).Output()
	if err != nil {
		msg := fmt.Sprintf("Could not run command %s %s \n This was the err %s", command,strings.Join(args," "),err.Error())
		fmt.Println(msg)
		// panic(msg)

	}
	return string(output)
}

func RunCommandInSpecifcDirectoryAndGetOutput(command string,args []string,targetDir string) string {

	fullCommand :=  fmt.Sprintf("Running command: %s %s", command,strings.Join(args," "))
	fmt.Println(fullCommand)
	cmd := exec.Command(command, args...)
	cmd.Dir = targetDir
	output,err := cmd.Output()
	if err != nil {
		msg := fmt.Sprintf("Could not run command %s %s \n This was the err %s", command,strings.Join(args," "),err.Error())
		fmt.Println(msg)
		// panic(msg)

	}
	return string(output)
}

type CommandOptions struct {
	Command      string
	Args         []string
	TargetDir    string
	GetOutput    bool
	PanicOnError bool
}

func RunCommandWithOptions(options CommandOptions) (string, error) {
	fullCommand := fmt.Sprintf("Running command: %s %s", options.Command, strings.Join(options.Args, " "))
	fmt.Println(fullCommand)

	cmd := exec.Command(options.Command, options.Args...)
	if options.TargetDir != "" {
		cmd.Dir = options.TargetDir
	}

	if options.GetOutput {
		output, err := cmd.Output()
		if err != nil {
			msg := fmt.Sprintf("Could not run command %s %s\nThis was the err: %s", options.Command, strings.Join(options.Args, " "), err.Error())
			fmt.Println(msg)

			if options.PanicOnError {
				panic(msg)
			}

			return "", err
		}
		return string(output), nil
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		msg := fmt.Sprintf("Could not run command %s %s\nThis was the err: %s", options.Command, strings.Join(options.Args, " "), err.Error())
		fmt.Println(msg)

		if options.PanicOnError {
			panic(msg)
		}

		return "", err
	}

	return "", nil
}



