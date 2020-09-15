package main

import (
	"context"
	"fmt"
	"time"
)

func Yield(ctx context.Context, cb func(string)) {
	for {
		tCh := time.After(3 * time.Second)
		select {
		case <-ctx.Done():
			return
		case <-tCh:
			cb(fmt.Sprintf("%+v", time.Now()))
		}
	}
}

func main() {
	Yield(context.Background(), func(a string) {
		fmt.Println(a)
	})
}
