package utils

import (
	"context"
	"encoding/json"
	"fmt"
)

type IpApiResp struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func GetIPLocationOnline(ctx context.Context, ip string) string {
	if ip == "" || ip == "127.0.0.1" {
		return "未知地址"
	}
	url := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)
	resp, err := RemoteDO(ctx, url, "GET", nil, nil)
	if err != nil {
		return "未知地址"
	}
	var data IpApiResp
	if err = json.Unmarshal(resp, &data); err != nil {
		return "未知地址"
	}
	return data.RegionName + "/" + data.City
}
