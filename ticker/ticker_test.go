package ticker

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	stop := make(chan string, 1)
	RunSecondNine().RunWithChannel(func(now time.Time) {
		fmt.Println("prepare target, ", now.Format("15:04:05.000"), "\trun at:", time.Now().Format("15:04:05.000"))
		stop <- "stop"
	}, stop)
}
