package msg

import (
	"github.com/whileW/im/user"
	"github.com/whileW/im/utils"
	"github.com/google/uuid"
	"time"
)

type RespMsgPush struct {
	ReqId 				string			`json:"req_id"`			//请求id
	Msg 				string			`json:"msg"`			//信息
	SendUser 			RespMsgSendUser	`json:"send_user"`
	Data 				interface{}		`json:"data"`			//数据
}
type RespMsgSendUser struct {
	UserName 		string			`json:"user_name"`
	UserId 			string			`json:"user_id"`
}

func RespMsgPushByte(req_id,msg string,user_id,user_name string,data interface{}) *msg_resp_struct {
	resp_msg := &msg_resp_struct{
		Function:function_push,
		Data:&RespMsgPush{
			ReqId:req_id,
			Msg:msg,
			Data:data,
			SendUser:RespMsgSendUser{
				UserId:user_id,
				UserName:user_name,
			},
		},
	}
	return resp_msg
}

type push_msg struct {
	req_id			string
	msg 			[]byte
	h 				*hub_struct
	retry_count 	int
	ok_chan 		chan int
}
var push_msg_pool = make(map[string]*push_msg)

func (h *hub_struct)AddPushRespMsg(msg string,data interface{},send_user user.UserInter){
	req_id := uuid.New().String()
	pmp := &push_msg{
		req_id:req_id,
		msg:RespMsgPushByte(req_id,msg,send_user.GetUserId(),send_user.GetUserName(),data).ToByte(),
		h:h,
		retry_count:1,
		ok_chan:make(chan int,1),
	}
	push_msg_pool[req_id] = pmp
	go send_push_resp_msg(pmp)
}
func send_push_resp_msg(pmp *push_msg)  {
	pmp.h.send_msg(pmp.msg)
	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-pmp.ok_chan:
			close(pmp.ok_chan)
			timer.Stop()
			return
		case <-timer.C:
			if pmp.retry_count < 5 {
				if pmp.h.State == 1 {
					return
				}
				pmp.retry_count ++
				pmp.h.send_msg(pmp.msg)
				timer.Reset(2 * time.Second)
			}else {
				utils.ZlLoggor.Error("send err:",string(pmp.msg))
				timer.Stop()
				return
			}
		}
	}
}
func received_push_msg(req_id string)  {
	if push_msg_pool[req_id] != nil {
		push_msg_pool[req_id].ok_chan<-1
		delete(push_msg_pool, req_id)
	}
}