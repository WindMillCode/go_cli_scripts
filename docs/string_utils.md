I apologize for the oversight. Let's provide the detailed markdown reference tables for each function and property in the `string-utils.go` file.

## [Reverse]
### Usage
Reverses the order of elements in an array of strings.

### Reference
| Parameter   | Type        | Description                         |
|-------------|-------------|-------------------------------------|
| targetArray | []string    | The array of strings to be reversed.|

| Returns     | Type        | Description                         |
|-------------|-------------|-------------------------------------|
| newArray    | []string    | The reversed array of strings.      |

## [UnicodeUnquote]
### Usage
Converts Unicode escape sequences in a byte slice to their corresponding characters.

### Reference
| Parameter | Type    | Description                                     |
|-----------|---------|-------------------------------------------------|
| bs        | []byte  | The byte slice containing Unicode escape codes. |

| Returns   | Type    | Description                                 |
|-----------|---------|---------------------------------------------|
| result    | []byte  | The byte slice with escape codes converted. |

## [TruncateStringByRegex]
### Usage
Truncates a string based on matches from a regular expression, guided by a predicate.

### Reference
| Parameter    | Type                              | Description                                           |
|--------------|-----------------------------------|-------------------------------------------------------|
| options      | TruncateStringByRegexOptions      | Struct containing input string, regex pattern, and predicate. |

| Returns      | Type     | Description                             |
|--------------|----------|-----------------------------------------|
| modifiedString| string  | The resulting string after truncation.  |

## [CreateStringObject]
### Usage
Generates a string object that provides various string transformation methods.

### Reference
| Parameter    | Type                        | Description                              |
|--------------|-----------------------------|------------------------------------------|
| myStr        | string                      | The string to create the object from.    |
| entitySuffix | string                      | A suffix to be considered in transformations. |

| Returns      | Type                         | Description                               |
|--------------|------------------------------|-------------------------------------------|
| result       | CreateStringObjectType       | An object containing string transformation methods. |

## [ContainsAny]
### Usage
Checks if the provided string contains any of the substrings from the given list.

### Reference
| Parameter | Type        | Description                              |
|-----------|-------------|------------------------------------------|
| s         | string      | The string to check.                     |
| substrs   | []string    | The list of substrings to look for.      |

| Returns   | Type    | Description                               |
|-----------|---------|-------------------------------------------|
| found     | bool    | True if any substring is found, else false. |

Each function and property in the `string-utils.go` file is now accompanied by a detailed markdown reference table.
