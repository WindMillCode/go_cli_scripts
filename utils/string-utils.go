package utils



func Reverse(targetArray []string) []string {
	newArray := make([]string, 0, len(targetArray))
	for i := len(targetArray)-1; i >= 0; i-- {
		newArray = append(newArray, targetArray[i])
	}
	return newArray
}
