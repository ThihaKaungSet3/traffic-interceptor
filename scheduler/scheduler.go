package scheduler

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"traffic/vendors"

	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
)

func RunJobs(cron *cron.Cron, memcache *cache.Cache) {
	go pingAndSave(memcache)
	cron.AddFunc("@every 5m", func() {
		memcache.Flush()
		pingAndSave(memcache)
	})
}

func pingAndSave(memcache *cache.Cache) {
	ips, _ := vendors.GetProxyScrapeFreemium()
	workingProxies := []vendors.ProxyConfig{}
	for _, proxy := range ips {
		if ping(proxy.IP, proxy.Port, "https://www.google.com/") {
			workingProxies = append(workingProxies, proxy)
			memcache.Set("proxies", workingProxies, cache.NoExpiration)
			fmt.Printf("proxy %s:%d is working!\n", proxy.IP, proxy.Port)
		} else {
			// fmt.Printf("failed to connect to proxy %s:%d\n", proxy.IP, proxy.Port)
		}
	}
}

func ping(host string, port int, target string) bool {
	address := fmt.Sprintf("%s//%s:%s", "http:", host, strconv.Itoa(port))
	proxyUrl, _ := url.Parse(address)
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
		DialContext: (&net.Dialer{
			Timeout:   4 * time.Second,
			KeepAlive: 4 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}
	res, err := client.Get(target)
	if err != nil {
		return false
	}
	if res.StatusCode == http.StatusOK {
		return true
	}
	return false
}
