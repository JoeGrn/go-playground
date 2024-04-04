package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	p(then)

	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Hour())
	p(now.Minute())
	p(now.Second())
	p(now.Nanosecond())
	p(now.Location())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
