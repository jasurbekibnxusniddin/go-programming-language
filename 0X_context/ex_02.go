package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go SlowTask(ctx)

	time.Sleep(2 * time.Second)
}

func SlowTask(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("slow task done")
	case <-ctx.Done():
		fmt.Println("slow task cancelled")
	}
}
