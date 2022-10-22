package config

import "time"

type AutoSync struct {
	Interval time.Duration `yaml:"interval"`
	LastRun  time.Time     `yaml:"last_run"`
}

func (a *AutoSync) ShouldRun() bool {
	return a.Interval > 0 && time.Now().Sub(a.LastRun) >= a.Interval
}

func (a *AutoSync) Run() {
	a.LastRun = time.Now()
}
