package config

import "time"

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
