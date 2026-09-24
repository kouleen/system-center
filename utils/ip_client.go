package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bytedance/gopkg/util/logger"
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
	client := http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)
	resp, err := client.Get(url)
	if err != nil {
		return "未知地址"
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			logger.CtxErrorf(ctx, "GetIPLocationOnline Close err: %v", err)
		}
	}(resp.Body)
	var data IpApiResp
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "未知地址"
	}
	return data.RegionName + "/" + data.City
}
