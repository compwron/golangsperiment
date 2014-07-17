package batch

import (
	"fmt"
	// "time"
)

func WriteToFileAfterReceivingMessages(messageCount int) {

	ccgn := channelContainingGivenNumbers(1, 2, 3, 1, 2, 3)
	for n := range ccgn {
		fmt.Println(n)
	}
	// ci := make(chan int, messageCount)

	// for i := 0; i < messageCount; i++ {
	// 	<-ci
	// }
}

func channelContainingGivenNumbers(nums ...int) <-chan int {
	out := make(chan int, 2)
	go func() {
		for a, b := range nums {
			fmt.Println("printing a and b", a, b)
			out <- b
			
		}
		close(out)
	}()
	return out
}
