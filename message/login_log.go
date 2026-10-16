package message

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/kouleen/common/message"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
	"github.com/kouleen/system-center/utils"
)

type LoginLogConsumer struct{}

func (LoginLogConsumer) Queue() string {
	return common.SYSTEM_LOGIN_LOG_QUEUE
}

func (LoginLogConsumer) HandleMessage(ctx context.Context, msg *message.Message) error {
	var req system.SystemLoginLogRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}
	if err := checkLoginLogReq(&req); err != nil {
		logger.CtxErrorf(ctx, "checkLoginLogReq error: %v", err)
		return nil
	}
	location := utils.GetIPLocationOnline(ctx, req.Ip)
	req.Location = location
	if _, err := service.CreateLoginLog(ctx, &req); err != nil {
		return err
	}
	return nil
}

func checkLoginLogReq(req *system.SystemLoginLogRequest) error {
	if req.GetUsername() == "" {
		return errors.New("username is empty")
	}
	if req.GetIp() == "" {
		return errors.New("ip is empty")
	}
	return nil
}

func init() {
	message.MustRegisterConsumer(new(LoginLogConsumer))
}
