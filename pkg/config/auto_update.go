package config

import "time"

type AutoUpdate struct {
	Interval time.Duration `yaml:"interval"`
	LastRun  time.Time     `yaml:"last_run"`
}
