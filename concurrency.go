// get content type of sites

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {

	channelAndSelect()

	timeOut()

	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

// implementation of channel and select
func channelAndSelect() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}
	fmt.Println("----")
}

// implementation of channel and select with time out
func timeOut() {
	out := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("time out")
	}
}
