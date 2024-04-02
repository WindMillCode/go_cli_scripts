# Docs

## Usage

The `array_utils.go` file in the `utils` package offers a variety of functions to assist with array operations in Go, providing functionality for element removal, existence checks, filtering, and type conversion.

## Reference

### RemoveElementsNotInSource
#### Description
Removes elements from the `toRemove` array that are not present in the `source` array.
#### Usage
```go
result := RemoveElementsNotInSource(sourceArray, toRemoveArray)
```
#### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| source | `[]T` | The source array to compare against. |
| toRemove | `[]T` | The array containing elements to be removed if not found in the source. |
| return value | `[]T` | An array containing elements not present in the source. |

### ArrayContainsAny
#### Description
Determines if any element in `targetSlice` is present in `contentSlice`.
#### Usage
```go
found := ArrayContainsAny(contentSlice, targetSlice)
```
#### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| contentSlice | `[]string` | The array to search within. |
| targetSlice | `[]string` | The array containing target elements to find. |
| return value | `bool` | `true` if any target element is found, otherwise `false`. |

### FilterArray
#### Description
Filters an array based on a provided condition function.
#### Usage
```go
filtered := FilterArray(array, conditionFunc)
```
#### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| arr | `[]T` | The array to be filtered. |
| condition | `func(interface{}, int) bool` | The function used to determine if an element should be included. |
| return value | `[]interface{}` | An array containing elements that meet the condition. |

The subsequent functions (`ConvertToInterfaceArray`, `ConvertToIntArray`, `ConvertToUintArray`, etc.) follow a similar pattern, converting arrays to different types based on the function's purpose, with their usage and reference details structured similarly to the above examples.
