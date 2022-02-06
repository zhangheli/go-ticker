package ticker

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSecondTicker(t *testing.T) {
	my := SecondTicker{}
	my.SetPoint([]int{0, 1, 2, 3, 4, 5})
	my.Run(make(chan string, 1), func() {
		tn := time.Now()
		fmt.Println(tn.Format(time.StampMicro))
		// fmt.Print("")
	})

	<-make(chan int)
}

func TestSecondTickerWithBusy(t *testing.T) {
	my := SecondTicker{}
	my.SetPoint([]int{1, 2, 3})
	my.Run(make(chan string, 1), func() {
		begin := time.Now().Format("15:04:05.000000")
		str := "123"

		message := fmt.Sprintf("done, %d", len(str))

		fmt.Printf("%s\t%s\n", begin, message)

		if strings.Contains(begin, ":00.") {
			os.Exit(0)
		}
	})
}
