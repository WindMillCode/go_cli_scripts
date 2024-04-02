## [FilterJSONByPredicate]
### Usage
`FilterJSONByPredicate` filters the JSON data based on a specified predicate function, allowing conditional inclusion of elements in the resulting JSON.

```go
filteredJSON, err := FilterJSONByPredicate(jsonData, func(key string, value interface{}) bool {
    return key == "desiredKey" // Example predicate: keep only the key "desiredKey"
})
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(filteredJSON))
```

### Reference

| Parameter   | Type                                      | Description                                      |
|-------------|-------------------------------------------|--------------------------------------------------|
| inputJSON   | []byte                                    | The input JSON data as a byte slice.             |
| predicate   | func(key string, value interface{}) bool | The predicate function to evaluate each key-value pair. |

| Returns     | Description                                  |
|-------------|----------------------------------------------|
| []byte      | The filtered JSON data as a byte slice.      |
| error       | An error if the operation fails, otherwise nil. |

## [WriteCustomFormattedJSONToFile]
### Usage
`WriteCustomFormattedJSONToFile` writes data to a file in a custom JSON format, allowing for specified indentation.

```go
err := WriteCustomFormattedJSONToFile(data, "output.json", "  ")
if err != nil {
    log.Fatal(err)
}
```

### Reference

| Parameter     | Type        | Description                                       |
|---------------|-------------|---------------------------------------------------|
| data          | interface{} | The data to be written in JSON format.            |
| filename      | string      | The name of the file to write the data to.        |
| indentString  | string      | The string used for indentation in the JSON output.|

| Returns       | Description                                  |
|---------------|----------------------------------------------|
| error         | An error if the operation fails, otherwise nil. |

## [RemoveComments]
### Usage
`RemoveComments` removes single line (`//`) and multi-line (`/* */`) comments from a byte slice containing JSON data, ensuring the remaining data is valid JSON.

```go
cleanData, err := RemoveComments(jsonDataWithComments)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(cleanData))
```

### Reference

| Parameter     | Type        | Description                                       |
|---------------|-------------|---------------------------------------------------|
| data          | []byte      | The JSON data with comments as a byte slice.      |

| Returns       | Description                                  |
|---------------|----------------------------------------------|
| []byte        | The cleaned JSON data without comments.       |
| error         | An error if the operation fails, otherwise nil. |

These functions provide utilities for processing JSON data, such as filtering elements based on conditions, writing data to a file with custom formatting, and removing comments from JSON content.
