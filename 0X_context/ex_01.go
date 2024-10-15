package main

import (
	"fmt"
	"time"
)

func main() {

	for counter := range 5 {
		go fmt.Println(counter)
	}

	//give some time for other goroutine to conclude
	time.Sleep(1 * time.Second)
}
