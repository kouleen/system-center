package utils

import (
	"context"
	"time"

	"github.com/bytedance/gopkg/util/logger"
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

func RemoteDO(ctx context.Context, requestURI, method string, headers map[string]string, body []byte) ([]byte, error) {
	request := protocol.AcquireRequest()
	response := protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(request)
		protocol.ReleaseResponse(response)
	}()
	request.SetMethod(method)
	request.SetHeaders(headers)
	request.SetRequestURI(requestURI)
	request.SetBody(body)
	if err := httpClient.Do(ctx, request, response); err != nil {
		logger.CtxErrorf(ctx, "remote do err: %v", err)
		return nil, err
	}
	response.Body()
	return response.Body(), nil
}
