package msg

import (
	"github.com/whileW/im/agreement"
	"github.com/whileW/im/user"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/url"
	"sync"
)

type hub_struct struct {
	req_id 		string
	user		user.UserInter
	conn 		agreement.AgreementInter
	pong 		chan int
	State 		int				//0正常    1断开
}

var hub_pool = sync.Map{}

func RegisterWebSocket(conn *websocket.Conn,token string,page,page_para,platform string) {
	h := &hub_struct{State:0,pong:make(chan int,1)}
	h.conn = agreement.RegisterWebSocket(conn)
	user_id, user_name, err := CheckAuth(token)
	if err != nil {
		register_err(h, err, "token")
		return
	}
	h.req_id = uuid.New().String()
	h.user, err = user.RegisterUser(h.req_id, user_id, user_name, page, page_para, platform,
		agreement.AgreementName_WebSocket, h.conn)
	if err != nil {
		register_err(h, err, "register user")
		return
	}
	hub_pool.Store(h.req_id, h)
	go h.read_msg()
	go h.conn_death()
	go h.set_ping()
}
func (h *hub_struct)conn_death()  {
	for {
		select {
		case <-h.conn.GetDeathChan():
			//客户端主动断开连接
			h.conn.Disconnect()
			h.State = 1
			h.user.OffLine()
		}
	}
}
func (h *hub_struct)Death()  {
	h.conn.Disconnect()
	h.State = 1
	h.user.OffLine()
}
func (h *hub_struct)read_msg()  {
	read_chan := h.conn.GetReadChan()
	for {
		select {
		case msg := <- read_chan:
			if check_msg_is_pong(msg) {
				h.pong<-1
			}else {
				HandReqMsg(msg,h.user)
			}
		}
	}
}
func register_err(h *hub_struct,err error,data interface{})  {
	h.SendOsRespMsg(err.Error(),data)
	h.conn.Disconnect()
}
//检查权限返回user_id、user_name
//func CheckAuth(token string) (string,string,error) {
//	user,err := logic.PaseToken(token)
//	if err != nil {
//		return "","",err
//	}
//	if user != nil {
//		return user.User.Id,user.User.Name,nil
//	}
//	return "","",errors.New("please login")
//}
func CheckAuth(name string) (string,string,error) {
	name, _ = url.QueryUnescape(name)
	return uuid.New().String(),name,nil
}

//发送消息
func (h *hub_struct) send_msg(msg []byte)  {
	h.conn.Send(msg)
}
func (h *hub_struct)SendOsRespMsg(msg string,data interface{})  {
	msg_b := RespMsgOsByte(msg,data).ToByte()
	h.send_msg(msg_b)
}
func (h *hub_struct)SendAnswerRespMsg(msg RespMsgInter)  {
	if msg != nil {
		h.send_msg(msg.ToByte())
	}
}

//查找hub
func FindHubByReqId(req_id string) *hub_struct {
	h,ok := hub_pool.Load(req_id)
	if ok {
		return h.(*hub_struct)
	}
	return nil
}