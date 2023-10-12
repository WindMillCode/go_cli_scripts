package utils

func OverwriteMap[T any](targetMap, newMap map[string]T) {
	for key, value := range newMap {
			targetMap[key] = value
	}
}

func FilterMap[T any](originalMap map[string]T, predicate func(string,T) bool) map[string]interface{} {
	filteredMap := make(map[string]interface{})

	for key, value := range originalMap {
			if predicate(key,value) {
					filteredMap[key] = value
			}
	}

	return filteredMap
}

func ConvertToStringIntMap(originalMap map[string]interface{}) map[string]int {
	intMap := make(map[string]int)

	for key, value := range originalMap {
			if intValue, ok := value.(int); ok {
					intMap[key] = intValue
			}
	}

	return intMap
}

func ConvertToStringInt8Map(originalMap map[string]interface{}) map[string]int8 {
	int8Map := make(map[string]int8)

	for key, value := range originalMap {
			if intValue, ok := value.(int8); ok {
					int8Map[key] = intValue
			}
	}

	return int8Map
}

func ConvertToStringInt16Map(originalMap map[string]interface{}) map[string]int16 {
	int16Map := make(map[string]int16)

	for key, value := range originalMap {
			if intValue, ok := value.(int16); ok {
					int16Map[key] = intValue
			}
	}

	return int16Map
}

func ConvertToStringInt32Map(originalMap map[string]interface{}) map[string]int32 {
	int32Map := make(map[string]int32)

	for key, value := range originalMap {
			if intValue, ok := value.(int32); ok {
					int32Map[key] = intValue
			}
	}

	return int32Map
}

func ConvertToStringInt64Map(originalMap map[string]interface{}) map[string]int64 {
	int64Map := make(map[string]int64)

	for key, value := range originalMap {
			if intValue, ok := value.(int64); ok {
					int64Map[key] = intValue
			}
	}

	return int64Map
}

func ConvertToStringUint8Map(originalMap map[string]interface{}) map[string]uint8 {
	uint8Map := make(map[string]uint8)

	for key, value := range originalMap {
			if uintValue, ok := value.(uint8); ok {
					uint8Map[key] = uintValue
			}
	}

	return uint8Map
}

func ConvertToStringUint16Map(originalMap map[string]interface{}) map[string]uint16 {
	uint16Map := make(map[string]uint16)

	for key, value := range originalMap {
			if uintValue, ok := value.(uint16); ok {
					uint16Map[key] = uintValue
			}
	}

	return uint16Map
}

func ConvertToStringUint32Map(originalMap map[string]interface{}) map[string]uint32 {
	uint32Map := make(map[string]uint32)

	for key, value := range originalMap {
			if uintValue, ok := value.(uint32); ok {
					uint32Map[key] = uintValue
			}
	}

	return uint32Map
}

func ConvertToStringUint64Map(originalMap map[string]interface{}) map[string]uint64 {
	uint64Map := make(map[string]uint64)

	for key, value := range originalMap {
			if uintValue, ok := value.(uint64); ok {
					uint64Map[key] = uintValue
			}
	}

	return uint64Map
}

func ConvertToStringFloat32Map(originalMap map[string]interface{}) map[string]float32 {
	float32Map := make(map[string]float32)

	for key, value := range originalMap {
			if floatValue, ok := value.(float32); ok {
					float32Map[key] = floatValue
			}
	}

	return float32Map
}

func ConvertToStringFloat64Map(originalMap map[string]interface{}) map[string]float64 {
	float64Map := make(map[string]float64)

	for key, value := range originalMap {
			if floatValue, ok := value.(float64); ok {
					float64Map[key] = floatValue
			}
	}

	return float64Map
}

func ConvertToStringComplex64Map(originalMap map[string]interface{}) map[string]complex64 {
	complex64Map := make(map[string]complex64)

	for key, value := range originalMap {
			if complexValue, ok := value.(complex64); ok {
					complex64Map[key] = complexValue
			}
	}

	return complex64Map
}

func ConvertToStringComplex128Map(originalMap map[string]interface{}) map[string]complex128 {
	complex128Map := make(map[string]complex128)

	for key, value := range originalMap {
			if complexValue, ok := value.(complex128); ok {
					complex128Map[key] = complexValue
			}
	}

	return complex128Map
}

func ConvertToStringBoolMap(originalMap map[string]interface{}) map[string]bool {
	boolMap := make(map[string]bool)

	for key, value := range originalMap {
			if boolValue, ok := value.(bool); ok {
					boolMap[key] = boolValue
			}
	}

	return boolMap
}

func ConvertToStringInterfaceMap(originalMap map[string]interface{}) map[string]string {
	stringMap := make(map[string]string)

	for key, value := range originalMap {
			if strValue, ok := value.(string); ok {
					stringMap[key] = strValue
			}
	}

	return stringMap
}

func ConvertToStringStringMap(originalMap map[string]interface{}) map[string]string {
	stringMap := make(map[string]string)

	for key, value := range originalMap {
			if strValue, ok := value.(string); ok {
					stringMap[key] = strValue
			}
	}

	return stringMap
}

func ConvertToStringErrorMap(originalMap map[string]interface{}) map[string]error {
	errorMap := make(map[string]error)

	for key, value := range originalMap {
			if errValue, ok := value.(error); ok {
					errorMap[key] = errValue
			}
	}

	return errorMap
}

