package main

import (
	"context"
	"fmt"
	"time"
)

type Yielder struct {
	done bool
	ch   chan string
}

func (y *Yielder) More() bool {
	return !y.done
}

func (y *Yielder) Yield() string {
	a := <-y.ch
	return a
}

func (y *Yielder) Run(ctx context.Context) {
	y.ch = make(chan string)
	go func() {
		for {
			tCh := time.After(3 * time.Second)
			select {
			case <-ctx.Done():
				y.done = true
				return
			case <-tCh:
				y.ch <- fmt.Sprintf("%+v", time.Now())
			}
		}
	}()
	return
}

func main() {
	y := &Yielder{}
	y.Run(context.Background())
	for y.More() {
		a := y.Yield()
		fmt.Println(a)
	}
}
