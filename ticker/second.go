package ticker

import (
	"log"
	"time"
)

type LoopConditionTicker struct {
	Ticker        time.Ticker
	ConditionFunc func(now time.Time) (int64, int64)
	Context       map[string]interface{}
}

func (this *LoopConditionTicker) RunWithChannel(task func(now time.Time), stopChannel chan string) {
	this.Context = map[string]interface{}{}
	for {
		select {
		case <-this.Ticker.C:
			now := time.Now()
			unixTime, nanoTime := this.ConditionFunc(now)
			if unixTime == 0 {
				continue
			}

			nextTime := time.Unix(unixTime, nanoTime)
			if cached, _has := this.Context["nextTime"]; _has && cached == nextTime {
				continue
			}

			this.Context["nextTime"] = nextTime

			waitFor := nextTime.Sub(now)
			time.AfterFunc(waitFor, func() {
				task(now)
			})
		case message := <-stopChannel:
			log.Println(message)
		}
	}
}

func (this *LoopConditionTicker) Run(task func(now time.Time)) {
	this.RunWithChannel(task, make(chan string, 1))
}

func RunSecondNine() *LoopConditionTicker {
	lc := LoopConditionTicker{
		Ticker: *time.NewTicker(time.Duration(3) * time.Second),
		ConditionFunc: func(n time.Time) (int64, int64) {
			delta := int64(10 - n.Unix()%10)
			if delta < 1 {
				return 0, 0
			}
			return n.Unix() + delta - 1, 0
		},
	}
	return &lc
}
