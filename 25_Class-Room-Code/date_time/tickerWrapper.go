package main

import (
	"fmt"
	"time"
)

type MyTicker struct {
	Duration time.Duration
	tck      *time.Ticker
	fn       func()
}

func (tck *MyTicker) Start(dur time.Duration, fn func()) {
	tck.Duration = dur
	tck.fn = fn
	tck.tck = time.NewTicker(tck.Duration)
	go func() {
		for {
			select {
			case <-tck.tck.C:
				tck.fn()
			}
		}
	}()
}
func (tck *MyTicker) Stop() {
	tck.tck.Stop()
}

func main() {
	var ticker MyTicker = MyTicker{}
	ticker.Start(1*time.Second, func() {
		fmt.Println("Ticker tic")
	})
	time.Sleep(20 * time.Second)
}
