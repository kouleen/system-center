package message

import (
	"context"
	"encoding/json"

	"github.com/kouleen/common/message"
)

type DDLLog struct{}

func (DDLLog) Queue() string {
	return "workflow.data.listener"
}

func (DDLLog) HandleMessage(ctx context.Context, msg *message.Message) error {
	var req BinlogMsg
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}
	return nil
}

type BinlogMsg struct {
	DB     string         `json:"db"`
	Table  string         `json:"table"`
	Action string         `json:"action"`
	Before map[string]any `json:"before,omitempty"`
	After  map[string]any `json:"after,omitempty"`
}

func init() {
	message.MustRegisterConsumer(new(DDLLog))
}
