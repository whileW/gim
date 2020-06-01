package msg

import (
	"errors"
)

//消息结构
type confirm_msg_struct struct {
	ReqId 		string			`json:"req_id"`
}

//消息处理
func ConfirmMsgHand(data map[string]string) error {
	if data["req_id"] == "" {
		return errors.New("please write req_id")
	}
	received_push_msg(data["req_id"])
	return nil
}