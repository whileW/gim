package service

type MessageHandleAck struct{
	*Connection
}

type MessageAckBody struct {
	State 			int
	Msg 			string
}

const (
	_AckState	=	iota
	Ack_Success				//成功
	Ack_Para_Codec_Err		//参数解码失败

	Ack_Auth_Please_Login	//请先登录
	Ack_Auth_Fail			//认证失败
)


func (c *MessageHandleAck)HandleMessage(msg *Message)  {
	switch msg.Cmd {
	case CmdACK:
		//确认收到消息
	}
}
func (c *MessageHandleAck)SendAckMsg(uuid string,state int,msg string)  {
	c.EnqueueMessage(&Message{Cmd:CmdACK,UUID:uuid,Body:MessageAckBody{State:state,Msg:msg}})
}
