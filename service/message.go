package service

type Message struct {
	UUID 		string		//uuid
	Cmd 		int			//命令
	//flag 		int
	timestamp	int64		//时间戳

	from 		string
	To 			string

	Body 		interface{}
	//BodyData	string
}

const (
	_cmd 		=	iota

	CmdHello 			//say hello
	CmdPing
	CmdPong

	CmdAuth 			//认证

	//CmdSendMsg		//发送普通消息
	CmdGroupMsg			//发送群组消息

	CmdACK				//确认消息
)