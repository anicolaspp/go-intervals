package intervals

import (
	"fmt"
	"math"
	"sort"
)

type Interval struct {
	Low  int
	High int
}

func (i *Interval) String() string {
	return fmt.Sprintf("[%v, %v]", i.Low, i.High)
}

func (i *Interval) Intersect(other *Interval) bool {
	if i.High < other.Low || i.Low > other.High {
		return false
	}

	return true
}

func (i *Interval) Merge(other *Interval) *Interval {
	return &Interval{
		Low:  int(math.Min(float64(i.Low), float64(other.Low))),
		High: int(math.Max(float64(i.Low), float64(other.High))),
	}
}

// O(n log n)
func Merge(intervals []*Interval) []*Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Low <= intervals[j].Low
	})

	for i := 0; i+1 < len(intervals); i++ {
		a := intervals[i]
		b := intervals[i+1]

		if a.Intersect(b) {
			c := a.Merge(b)
			xs := append(intervals[:i], c)
			intervals = append(xs, intervals[i+2:]...)
			i--
		}
	}

	return intervals
}
