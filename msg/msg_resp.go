package msg

import (
	"encoding/json"
)

type RespMsgInter interface {
	ToByte()	[]byte
} 

type msg_resp_struct struct {
	Function 			string			`json:"function"`		//功能
	Data				interface{}	 	`json:"data"`			//数据
}

const (
	function_answer		=		"answer"		//应答
	function_push 		=		"push"			//推送
	function_os			=		"os"			//系统
)

//回答结构体
type RespMsgAnswer struct {
	ReqId 				string			`json:"req_id"`			//请求id
	Result				string 			`json:"result"`			//结果
	Msg 				string			`json:"msg"`			//信息
	Data 				interface{}		`json:"data"`			//数据
}
type RespMsgOs struct {
	Msg 				string			`json:"msg"`			//信息
	Data 				interface{}		`json:"data"`			//数据
} 

func RespMsgAnswerByte(req_id ,result,msg string,data interface{}) *msg_resp_struct {
	resp_msg := &msg_resp_struct{
		Function:function_answer,
		Data:&RespMsgAnswer{
			ReqId:req_id,
			Result:result,
			Msg:msg,
			Data:data,
		},
	}
	return resp_msg
}
func RespMsgOsByte(msg string,data interface{}) *msg_resp_struct {
	resp_msg := &msg_resp_struct{
		Function:function_push,
		Data:&RespMsgOs{
			Msg:msg,
			Data:data,
		},
	}
	return resp_msg
}

func (r *msg_resp_struct)ToByte() []byte {
	str,_ := json.Marshal(r)
	return str
}