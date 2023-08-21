package utils

import "fmt"

func DecreaseChannelBatchFn(i int, batchSize int, batchDone chan bool, targetArray []string) {
	if i%batchSize == 0 {

		for j := 0; j < batchSize; j++ {
			<-batchDone
		}
		fmt.Println("Batch complete")
	} else if i == len(targetArray) {

		for j := 0; j < len(targetArray)%batchSize; j++ {
			<-batchDone
		}
		fmt.Println("Batch complete")
	}
}
