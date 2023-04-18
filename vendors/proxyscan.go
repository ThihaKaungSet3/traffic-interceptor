package vendors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetProxyScan() ([]ProxyConfig, error) {
	res, err := http.Get("https://www.proxyscan.io/api/proxy?ping=300&limit=20&type=http")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var proxies []proxy
	er := json.NewDecoder(res.Body).Decode(&proxies)
	if er != nil {
		fmt.Println(er)
		return []ProxyConfig{}, nil
	}
	var proxyConfigs []ProxyConfig
	for _, p := range proxies {
		port := 0
		fmt.Sscanf(strconv.Itoa(p.Port), "%d", &port)
		proxyConfig := ProxyConfig{IP: p.Ip, Port: port}
		proxyConfigs = append(proxyConfigs, proxyConfig)
	}
	return proxyConfigs, nil
}

type proxy struct {
	Ip   string `json:"Ip"`
	Port int    `json:"Port"`
}
