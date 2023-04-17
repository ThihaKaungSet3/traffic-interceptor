package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"traffic/vendors"
)

func main() {
	ips, _ := vendors.GetFreemium()
	fmt.Println(ips)
	workingProxies := []vendors.ProxyConfig{}
	for _, proxy := range ips {
		if ping(proxy.IP, proxy.Port, "https://www.google.com/") {
			workingProxies = append(workingProxies, proxy)
			fmt.Printf("proxy %s:%d is working!\n", proxy.IP, proxy.Port)
		} else {
			fmt.Printf("failed to connect to proxy %s:%d\n", proxy.IP, proxy.Port)
		}
	}

	// ping("185.162.251.76", 80, "https://www.google.com/")
}

func ping(host string, port int, target string) bool {
	address := fmt.Sprintf("%s//%s:%s", "http:", host, strconv.Itoa(port))
	println(address)
	proxyUrl, _ := url.Parse(address)
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
		DialContext: (&net.Dialer{
			Timeout:   4 * time.Second,
			KeepAlive: 4 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
	res, err := client.Get(target)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	if res.StatusCode == http.StatusOK {
		return true
	}
	return false
}
