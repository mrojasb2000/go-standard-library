package main

import (
	"fmt"
	"time"
)

/*
 The timeout for long-running operation in the code is implemented
 with the use of the time.After function, which provides the channel
 delivering the tick after the given period.

 The operation itself is wrapped to select a statement which
 chooses between the time.After channel and the default option,
 which executes the operation.

 Note that you need to allow the code to read from the time.After
 channel periodically to fin out whether the timeout is exceeded or not.
 Otherwise, if the default code branch blocks the execution entirely,
 there is no way how to fin out if the tiemout has already elapsed.
*/
func main() {

	to := time.After(3 * time.Second)
	list := make([]string, 0)
	done := make(chan bool, 1)

	fmt.Println("Starting to insert items")

	go func() {
		defer fmt.Println("Exiting goroutine")
		for {
			select {
			case <-to:
				fmt.Println("The time is up")
				done <- true
				return
			default:
				list = append(list, time.Now().String())
			}
		}
	}()
	<-done
	fmt.Printf("Managed to insert %d items\n", len(list))
}
