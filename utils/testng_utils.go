package utils

import (
	"path/filepath"
)

type GetTestNGArgsStruct struct{
	WorkspaceFolder string
	EnvVarsFile string
	TestNGFolder string
	SuiteFile string
	ParamEnv string
}
func GetTestNGArgs(c GetTestNGArgsStruct) GetTestNGArgsStruct {
	c.EnvVarsFile = GetInputFromStdin(
		GetInputFromStdinStruct{
			Prompt: []string{"script where env vars are set for the app to run relative to workspace root"},
			Default: filepath.Join(c.WorkspaceFolder,"ignore\\Local\\testng_e2e_shared.env"),
		},
	)

	c.TestNGFolder = GetInputFromStdin(
		GetInputFromStdinStruct{
			Prompt: []string{"testng app location"},
			Default: filepath.Join(c.WorkspaceFolder,"apps\\testing\\testng"),
		},
	)

	c.SuiteFile = GetInputFromStdin(
		GetInputFromStdinStruct{
			Prompt: []string{"xml suite file needed for testng (this should be relative to the testng folder)"},
			Default: filepath.Join("src\\test\\resources\\tests.xml"),
		},
	)

	cliInfo := ShowMenuModel{
		Prompt: "the environment to test",
		Choices:[]string{"DEV","PREVIEW","PROD"},
		Default:"DEV",
		Other:true,
	}
	c.ParamEnv = ShowMenu(cliInfo,nil)
	return c
}
