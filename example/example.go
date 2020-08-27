package main

import (
	"fmt"

	"github.com/nknorg/go-fast"
)

func main() {
	fastCom, err := fast.New("")
	if err != nil {
		panic(err)
	}

	// init
	err = fastCom.Init()
	if err != nil {
		panic(err)
	}

	// get urls
	urls, err := fastCom.GetUrls()
	if err != nil {
		panic(err)
	}

	// measure
	KbpsChan := make(chan float64)

	go func() {
		for Kbps := range KbpsChan {
			fmt.Printf("%.2f Kbps %.2f Mbps\n", Kbps, Kbps/1000)
		}

		fmt.Println("done")
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		panic(err)
	}

	fmt.Println("exit")
}
