package batch

import (
	"fmt"
)

func WriteToFileAfterReceivingMessages(messageCount int) {
	fmt.Println("the intended count", messageCount)
	c := gen(1, 2, 3, 4, 5, 6, 7, 8, 9)
	batchedStuff, out := batchStuffUp(c)

	fmt.Println("the batched stuff", batchedStuff) 
	<-out
	<-out
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func batchStuffUp(in <-chan int) ([]string, <-chan int) {
	out := make(chan int)
	messsageCount := 5
	batchOfInput := []string{}
	go func() {
		
		currentBatch := ""
		count := 0

		fmt.Println("\ninside go func() but before n:= range")
		for n := range in {
			out <- n * n
			fmt.Println("\ninside for n := range in { but before if/else")
			if count >= messsageCount {
				fmt.Println("in if before", currentBatch, count)
				currentBatch = "foo"
				count = 0
				fmt.Println("in if after", currentBatch, count)
			} else {
				fmt.Println("in else before", currentBatch, count)
				currentBatch += string(n)
				count += 1
				fmt.Println("in else after", currentBatch, count)
			}
		}
		close(out)
	}()
	return batchOfInput, out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
