package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	t := time.NewTimer(3 * time.Second)

	fmt.Printf("Start waiting as %v\n", time.Now().Format(time.UnixDate))
	<-t.C
	fmt.Printf("Code executed at %v\n", time.Now().Format(time.UnixDate))

	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Printf("Start waiting form AfterFunc at %v\n", time.Now().Format(time.UnixDate))
	time.AfterFunc(3*time.Second, func() {
		fmt.Printf("Code executed for AfterFunc at %v\n", time.Now().Format(time.UnixDate))
		wg.Done()
	})
	wg.Wait()
	fmt.Printf("Waiting on time.After at %v\n", time.Now().Format(time.UnixDate))
	<-time.After(3 * time.Second)
	fmt.Printf("Code resumed for at %v\n", time.Now().Format(time.UnixDate))
}
