package main

import (
	"time"
	"traffic/scheduler"

	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
)

func main() {
	job := cron.New()
	c := cache.New(5*time.Minute, 10*time.Minute)
	scheduler.RunJobs(job, c)
}
