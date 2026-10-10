package handle

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
)

var httpClient *client.Client

func init() {
	var err error
	httpClient, err = client.NewClient(
		client.WithDialer(standard.NewDialer()),
		client.WithDialTimeout(2*time.Second),
		client.WithClientReadTimeout(30*time.Second),
		client.WithWriteTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
}

func TestBulletinHandle(t *testing.T) {
	testGet()
}

func testGet() {
	request := protocol.AcquireRequest()
	response := protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(request)
		protocol.ReleaseResponse(response)
	}()
	request.SetMethod("GET")
	request.SetRequestURI("https://www.baidu.cn")

	err := httpClient.Do(context.Background(), request, response)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode(), string(response.Body()))
}
