package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhangheli/go-ticker/ticker"
)

func TestGatlin(t *testing.T) {
	action := func() {
		fmt.Println(time.Now())
	}
	now := time.Now().UnixMilli()
	p := now + (1000 - now%1000)
	runBegin := time.UnixMilli(p)

	ticker.GatlinWithLimit(runBegin, 10, action)
}

func TestTicker(t *testing.T) {
	action := func() {
		fmt.Println(time.Now())
	}
	stopC := make(chan string, 1)

	st := ticker.SecondTicker{}
	st.SetPoint([]int{1, 100, 200, 300, 400, 500})
	st.Run(stopC, action)

	<-make(chan int, 1)
}
