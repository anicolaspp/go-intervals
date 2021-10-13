package intervals_test

import (
	"testing"

	intervals "github.com/anicolaspp/intervals/interval"
)

func TestIntervalIntersect(t *testing.T) {
	for _, tc := range []struct {
		name      string
		a         *intervals.Interval
		b         *intervals.Interval
		intersect bool
	}{
		{
			name:      "empty intervals intersect",
			a:         &intervals.Interval{Low: 0, High: 0},
			b:         &intervals.Interval{Low: 0, High: 0},
			intersect: true,
		},
		{
			name:      "exclusive",
			a:         &intervals.Interval{Low: 0, High: 10},
			b:         &intervals.Interval{Low: 11, High: 20},
			intersect: false,
		},
		{
			name:      "same intervals",
			a:         &intervals.Interval{Low: 11, High: 20},
			b:         &intervals.Interval{Low: 11, High: 20},
			intersect: true,
		},
		{
			name:      "within",
			a:         &intervals.Interval{Low: 11, High: 20},
			b:         &intervals.Interval{Low: 12, High: 18},
			intersect: true,
		},
		{
			name:      "reverse within",
			b:         &intervals.Interval{Low: 11, High: 20},
			a:         &intervals.Interval{Low: 12, High: 18},
			intersect: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.a.Intersect(tc.b) != tc.intersect {
				t.Errorf("%v.Intersect(%v) != %v", tc.a, tc.b, tc.intersect)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		name string
	}{} {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
