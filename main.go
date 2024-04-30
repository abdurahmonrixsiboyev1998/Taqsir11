package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	res := make(chan int)
	var wg sync.WaitGroup
	num := func(ID int) {
		defer wg.Done()
		time.Sleep(time.Duration(ID) * time.Second) 
		res <- ID                                  
	}

	wg.Add(3)

	go num(4)
	go num(5)
	go num(6)

	go func() {
		wg.Wait()
		close(res)
	}()

	for res := range res {
		fmt.Println(res)
	}
}