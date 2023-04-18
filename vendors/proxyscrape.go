package vendors

import (
	"bufio"
	"net/http"
	"strconv"
	"strings"
)

func GetProxyScrapeFreemium() ([]ProxyConfig, error) {
	res, err := http.Get("https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=1000&country=all&ssl=all&anonymity=all")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(res.Body)
	ips := []string{}
	for scanner.Scan() {
		ips = append(ips, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	available := []ProxyConfig{}
	for _, ip := range ips {
		parts := strings.Split(ip, ":")
		ipAddress := parts[0]
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		config := ProxyConfig{
			IP:   ipAddress,
			Port: port,
		}
		available = append(available, config)
	}
	return available, nil
}
