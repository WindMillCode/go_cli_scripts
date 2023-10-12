// 19 types
package utils

func FilterArray[T any](arr []T, condition func(interface{},int) bool) []interface{} {
	var filtered []interface{}
	for index, element := range arr {
			if condition(element,index) {
					filtered = append(filtered, element)
			}
	}
	return filtered
}

func ConvertToInterfaceArray[T any](input []T) []interface{} {
	var result []interface{}
	for _, v := range input {
			result = append(result, v)
	}
	return result
}

func ConvertToIntArray(input []interface{}) []int {
    var result []int
    for _, v := range input {
        if num, ok := v.(int); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToUintArray(input []interface{}) []uint {
    var result []uint
    for _, v := range input {
        if num, ok := v.(uint); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToInt8Array(input []interface{}) []int8 {
    var result []int8
    for _, v := range input {
        if num, ok := v.(int8); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToUint8Array(input []interface{}) []uint8 {
    var result []uint8
    for _, v := range input {
        if num, ok := v.(uint8); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToInt16Array(input []interface{}) []int16 {
    var result []int16
    for _, v := range input {
        if num, ok := v.(int16); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToUint16Array(input []interface{}) []uint16 {
    var result []uint16
    for _, v := range input {
        if num, ok := v.(uint16); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToInt32Array(input []interface{}) []int32 {
    var result []int32
    for _, v := range input {
        if num, ok := v.(int32); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToUint32Array(input []interface{}) []uint32 {
    var result []uint32
    for _, v := range input {
        if num, ok := v.(uint32); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToInt64Array(input []interface{}) []int64 {
    var result []int64
    for _, v := range input {
        if num, ok := v.(int64); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToUint64Array(input []interface{}) []uint64 {
    var result []uint64
    for _, v := range input {
        if num, ok := v.(uint64); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToFloat32Array(input []interface{}) []float32 {
    var result []float32
    for _, v := range input {
        if num, ok := v.(float32); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToFloat64Array(input []interface{}) []float64 {
    var result []float64
    for _, v := range input {
        if num, ok := v.(float64); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToComplex64Array(input []interface{}) []complex64 {
    var result []complex64
    for _, v := range input {
        if num, ok := v.(complex64); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToComplex128Array(input []interface{}) []complex128 {
    var result []complex128
    for _, v := range input {
        if num, ok := v.(complex128); ok {
            result = append(result, num)
        }
    }
    return result
}

func ConvertToBoolArray(input []interface{}) []bool {
    var result []bool
    for _, v := range input {
        if b, ok := v.(bool); ok {
            result = append(result, b)
        }
    }
    return result
}

func ConvertToStringArray(input []interface{}) []string {
	var result []string
	for _, v := range input {
			if str, ok := v.(string); ok {
					result = append(result, str)
			}
	}
	return result
}

func ConvertToByteArray(input []interface{}) []byte {
	var result []byte
	for _, v := range input {
			if b, ok := v.(byte); ok {
					result = append(result, b)
			}
	}
	return result
}

func ConvertToRuneArray(input []interface{}) []rune {
	var result []rune
	for _, v := range input {
			if r, ok := v.(rune); ok {
					result = append(result, r)
			}
	}
	return result
}
