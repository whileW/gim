package msg

import (
	"im/user"
	"encoding/json"
	"errors"
)

type msg_req_struct struct {
	ReqId 				string				`json:"req_id"`			//请求id
	Function			string				`json:"function"`		//功能
	Data 				map[string]string	`json:"data"`			//数据
}

const (
	function_Chat 		=		"chat"			//im聊天
	function_confirm	=		"confirm"		//确认收到通知
)

var resp_err = errors.New("")

func HandReqMsg(msg []byte,send_user user.UserInter) RespMsgInter {
	msg_struct := &msg_req_struct{}
	if err := json.Unmarshal(msg,msg_struct);err != nil {
		return RespMsgAnswerByte("","false","数据结构错误:msg_struct:"+string(msg),nil)
	}
	var err error
	switch msg_struct.Function {
	case function_Chat:
		//聊天
		switch send_user.GetUserPage() {
		case user.Page_Im_Chat_Room:
			//聊天室
			err = ImChatRoomMsgHand(send_user,msg_struct.Data)
			break
		}
		break
	case function_confirm:
		//确认收到通知
		err = ConfirmMsgHand(msg_struct.Data)
		if err == nil {
			return nil
		}
		break
	}
	if err != nil {
		return RespMsgAnswerByte(msg_struct.ReqId,"false",err.Error(),nil)
	}
	return RespMsgAnswerByte(msg_struct.ReqId,"true","",nil)
}
