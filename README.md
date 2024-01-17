# Changelog

## v1.0.0
* utils.GetItemsInFolder now has just one argument for convenience

## v1.0.1

* added FilterArray which makes it more convenient to filter an array given an array and predicate(condition function)
* added methods that would cast the interface to the appropriate type

## v1.0.2
* made a slight change

## v1.0.3
* added index to element

## v1.1.0
* added FilterMapByKeys all all the respective casting fns

## v1.1.1
* added func OverwriteMap which is like js Object.assign

## v1.1.2
* added func OverwriteMap which is like js Object.assign

## v2.0.0
* replaced FilterMapByKeys with FilterMap where its predicate fn accepts key,val


## v2.0.1
* added FilterJSONByPredicate

## v2.0.2
* added WriteCustomFormattedJSONToFile to help format files

## v2.0.5
* WriteCustomFormattedJSONToFile supports bytes and interface strcutures, added UnicodeUnquote which will remove all unicode from a string when writing from bytes to a file,

## v2.0.6
* added AddContentToEachLineInFile AddContentToFile, which takes a predicate function and updates the file based on  the return of the predicate fn

## v2.1.1
* added MergeDirectories fn which would merge all files and folders from target dir into source dir w/o overrting anything

## v2.1.2
* added RunCommandWithOptions fn which supports optional target dir,optional get output, and panic on error

## v2.1.3
* added TraverseDirectory with a predicate fn

## v2.1.4
* added TruncateStringByRegex fn which allows the end user to provide a regex and has a predicate fn for every match in the pattern matcher which if returns true removes the substr from the array


## v2.1.5
* additional updates

## v3.0.0
* added cli and changed name to go_cli_scripts

## v3.1.0
* updated RunCommandWithOptions to print out standard err along with the reason why the command failed

## v3.1.1
* indicated all RunCommand fns are deprecated and RunCommandWithOptions should be used instead

## v3.1.2
* added ProcessFoldersMatchingPattern just like ProcessFilesMatchingPattern



## v3.1.5
added CreateStringObject an object that will give all sorts of cases
  camelCase
  kebab
  snakeCase
  classify
* fixied CreateStringObject

## v3.1.6
added IsRunningInDocker to see whether the given go script is running in a docker container or not

## v3.2.0
* [PATCH] updated GetInputFromStdin and ShowMenu to immediately cancel when the user hits Ctrl+C

## v3.2.1 [1/5/24]
* [PATCH] build mechanism seems to unexpectedly add unwanted code this patch should fix that

## v3.3.0 [1/6/24]
* [PATCH] fixed issue with RunCommandWithOptions where output gets returned from the program and printed to console
* added Uppercase method to CreateStringObjectType

## v3.3.1 [1/11/24]
* [PATCH] IsRunningInDocker will able to detect for macbooks as well

## v3.3.3 [1/17/24]
* [UPDATE] added 	ProxyURLs string `json:"proxyURLs"`
for the VSCodeSettings.ExtensionPack struct

## v4.0.0 [1/17/24]
* [BREAKING CHANGE] TakeVariableArgs returns TakeVariableArgsResultStruct
```go
type TakeVariableArgsResultStruct struct{
	inputString  string
	inputArray   []string
}
```
