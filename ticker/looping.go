package ticker

import (
	"log"
	"sync"
	"time"
)

type SecondTicker struct {
	milliPoint []int // 需要执行的列表
	history    []int
	lock       sync.Mutex // 锁
	roundTime  time.Time  // 当前迭代轮
}

func (st *SecondTicker) SetPoint(points []int) {
	st.milliPoint = []int{}
	if points == nil {
		st.milliPoint = []int{0}
	} else {
		for _, n := range points {
			st.milliPoint = append(st.milliPoint, (n)%1000)
		}
	}
}

func (st *SecondTicker) Run(stopChan chan string, f func()) {
	if st.milliPoint == nil || len(st.milliPoint) <= 0 {
		return
	}

	t := time.NewTicker(time.Duration(100) * time.Microsecond)
	for {
		select {
		case <-t.C:
			now := time.Now()
			milli := now.UnixNano() / (1000 * 1000)
			calcMilli := milli % 1000
			if st.getLock(now, calcMilli) {
				go f()
			}

		case signal := <-stopChan:
			log.Println(signal)
			return
		}
	}
}

func (st *SecondTicker) KeepRun(f func()) {
	if st.milliPoint == nil || len(st.milliPoint) <= 0 {
		return
	}

	t := time.NewTicker(time.Duration(100) * time.Microsecond)
	for {
		select {
		case <-t.C:
			now := time.Now()
			mic := now.UnixNano() / 1000
			if (mic%1000)/100 > 2 {
				continue
			} else if st.getLock(now, mic/1000) {
				go f()
			}
		}
	}
}

func (st *SecondTicker) RunWithCondition(stopChan chan string, condition func(time.Time) bool, task func()) {
	t := time.NewTicker(time.Duration(1) * time.Millisecond)
	for {
		st.lock.Lock()
		select {
		case <-t.C:
			now := time.Now()
			if condition(now) {
				go task()
			}
		case signal := <-stopChan:
			log.Println(signal)
			return
		}
		st.lock.Unlock()
	}
}

func (st *SecondTicker) getLock(now time.Time, milli int64) bool {
	st.lock.Lock()
	defer st.lock.Unlock()

	if st.roundTime.Unix() != now.Unix() {
		st.history = st.milliPoint
		st.roundTime = now
	}

	if len(st.history) <= 0 {
		return false
	}

	line := int(now.UnixNano() / (1000 * 1000) % 1000)
	if line >= st.history[0] {
		length := len(st.history)
		for length > 0 && st.history[0] <= line {
			st.history = st.history[1:length]
			length = len(st.history)
		}
		return true
	}
	return false
}
