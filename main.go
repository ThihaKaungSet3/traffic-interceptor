package main

import (
	"fmt"
	"net/http"
	"time"
	"traffic/api"
	"traffic/scheduler"
	"traffic/vendors"

	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
)

func main() {
	job := cron.New()
	c := cache.New(5*time.Minute, 10*time.Minute)
	api.SetUpRoutes(c)
	scheduler.RunJobs(job, c)
	job.Start()
	cou := vendors.GetRandomCountry()
	fmt.Println(cou)
	http.ListenAndServe(":3333", nil)
	select {}
}
