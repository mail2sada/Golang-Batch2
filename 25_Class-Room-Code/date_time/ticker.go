package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Tickers")

	tck := time.NewTicker(1 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case val := <-tck.C:
			fmt.Println("Ticker", val)
		}
	}

	tck.Stop()

}
