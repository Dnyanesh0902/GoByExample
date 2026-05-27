package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int){
	fmt.Println("worker %d starting \n",id)
	time.Sleep(time.Second)
	fmt.Println("worker %d done\b", id)
}

func main() {
	var wg sync.WaitGroup


	for i := 1; i <= 5; i++ {
		wg.Go(func ()  {
			worker(i)
		})
	}
	wg.Wait()
}