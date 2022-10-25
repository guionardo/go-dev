package config

import (
	"fmt"
	"time"
)

type IntervalRunner struct {
	Interval time.Duration `yaml:"interval"`
	LastRun  time.Time     `yaml:"last_run"`
}

func (a *IntervalRunner) ShouldRun() bool {
	return a.Interval > 0 && time.Now().Sub(a.LastRun) >= a.Interval
}

func (a *IntervalRunner) Run() {
	a.LastRun = time.Now()
}

func (a *IntervalRunner) String() string {
	return fmt.Sprintf("Interval: %s, LastRun: %s", a.Interval, a.LastRun.Format("2006-01-02 15:04:05"))
}
