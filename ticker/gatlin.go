package ticker

import (
	"time"
)

type Book struct {
	Second  time.Time
	Segment []int64
}

func Gatlin(begin time.Time, task func()) {
	GatlinWithLimit(begin, 1000, task)
}

func GatlinWithLimit(begin time.Time, limit int, task func()) {
	sst := SecondTicker{}
	sst.SetPoint([]int{0})
	mcs := make(chan string, 1)

	book := Book{}

	records := 0
	sst.RunWithCondition(mcs, func(now time.Time) bool {
		milli := time.Now().UnixNano() / 1000000
		melli := begin.UnixNano() / 1000000
		if milli < melli {
			return false
		}

		pos := milli % 1000

		if book.Segment == nil || book.Second.Unix() != now.Unix() {
			book.Second = now
			book.Segment = []int64{}
		}

		level := pos / 5
		for _, vv := range book.Segment {
			if vv == level {
				return false
			}
		}
		book.Segment = append(book.Segment, level)
		return true

	}, func() {
		if records > limit {
			mcs <- "gatlin exceeding limit, exit."
			return
		}
		task()
		records = records + 1
	})
}
