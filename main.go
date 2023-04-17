package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"traffic/vendors"
)

func main() {
	vendors.GetFreemium()
	// ping("185.162.251.76", 80, "https://www.google.com/")
}

func ping(host string, port int, target string) {
	address := fmt.Sprintf("%s//%s:%s", "http:", host, strconv.Itoa(port))
	println(address)
	proxyUrl, _ := url.Parse(address)
	transport := &http.Transport{
		Proxy:               http.ProxyURL(proxyUrl),
		TLSHandshakeTimeout: 3 * time.Second,
		IdleConnTimeout:     3 * time.Second,
	}
	client := &http.Client{Transport: transport}
	resp, err := client.Get(target)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	body := make([]byte, 1024)
	resp.Body.Read(body)
	fmt.Println("Body:", string(body))
}
