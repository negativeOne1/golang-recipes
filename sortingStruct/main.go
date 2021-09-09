package main

import (
	"fmt"
	"sort"
	"time"
)

func absInt64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

type TimeDataArray []TimeData

func (t TimeDataArray) FindClosest(d time.Time) TimeData {
	dst := make([]int64, len(t))
	for i, t := range t {
		dst[i] = absInt64(t.date.Unix() - d.Unix())
	}

	minI := 0
	minD := dst[0]
	for i, d := range dst {
		if d < minD {
			minD = d
			minI = i
		}
	}

	return t[minI]
}

type TimeData struct {
	date     time.Time
	exchange map[string]float64
}

var (
	day       = 24 * time.Hour
	today     = time.Date(2021, time.September, 9, 0, 0, 0, 0, time.UTC)
	yesterday = today.Add(-day)
	tomorrow  = today.Add(day)
	d         = TimeDataArray{
		{
			date: today, exchange: map[string]float64{"USD": 1.1},
		},
		{
			date: tomorrow, exchange: map[string]float64{"USD": 1.1},
		},
		{
			date: yesterday, exchange: map[string]float64{"USD": 1.1},
		},
	}
)

func main() {
	sort.Slice(d, func(i, j int) bool {
		return d[i].date.After(d[j].date)
	})

	fmt.Println(d.FindClosest(today).date)
}

func (t TimeDataArray) FindClosestFast(d time.Time) TimeData {
	d = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)

	minI := 0
	minD := t[0].date.Unix() - d.Unix()

	dst := make([]int64, len(t))
	for i, x := range t {
		dst[i] = absInt64(x.date.Unix() - d.Unix())

		if dst[i] < minD {
			minD = dst[i]
			minI = i
		}

		if dst[i] == 0 {
			return x
		}
	}

	return t[minI]
}
