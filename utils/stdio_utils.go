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
}

func TakeVariableArgs(obj TakeVariableArgsStruct) string {
	var innerScriptArguments []string

	fmt.Println(obj.Prompt)
	fmt.Println("Enter the arguments to pass to the script (press ENTER to enter another argument, leave blank and press ENTER once done):")
	for {
		var argument string
		fmt.Scanln(&argument)

		if strings.TrimSpace(argument) == "" {
			break
		}

		innerScriptArguments = append(innerScriptArguments, argument)
	}
	input := strings.Join(innerScriptArguments," ")
	if(input == "" && obj.ErrMsg != ""){
		panic(obj.ErrMsg)
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
	}
}

func RunCommandAndGetOutput(command string,args []string) string {

	fullCommand :=  fmt.Sprintf("Running command: %s %s", command,strings.Join(args," "))
	fmt.Println(fullCommand)
	output,err := exec.Command(command, args...).Output()
	if err != nil {
		msg := fmt.Sprintf("Could not run command %s %s \n This was the err %s", command,strings.Join(args," "),err.Error())
		panic(msg)

	}
	return string(output)
}




