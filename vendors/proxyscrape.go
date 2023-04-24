package vendors

import (
	"bufio"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetProxyScrapeFreemium() ([]ProxyConfig, error) {
	country := getRandomCountry()
	url := "https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=1000&country=" + country.Code + "&ssl=all&anonymity=all"
	res, err := http.Get(url)
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
			Code: country.Code,
		}
		available = append(available, config)
	}
	return available, nil
}

func getRandomCountry() country {
	rand.Seed(time.Now().UnixNano())
	countries := []country{
		{Name: "US", Code: "us"},
		{Name: "Canada", Code: "ca"},
		{Name: "Iran", Code: "ir"},
		{Name: "Norway,Netherland,Czech,Japan", Code: "no,nl,cz,jp"},
	}
	randomIndex := rand.Intn(len(countries))
	return countries[randomIndex]
}

type country struct {
	Name string
	Code string
}
