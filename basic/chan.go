package main

import (
	"context"
	"fmt"
	"time"
)

func Yield(ctx context.Context) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			tCh := time.After(3 * time.Second)
			select {
			case <-ctx.Done():
				return
			case <-tCh:
				ch <- fmt.Sprintf("%+v", time.Now())
			}
		}
	}()
	return ch
}

func main() {
	c := Yield(context.Background())
	for {
		a := <-c
		fmt.Println(a)
	}
}
