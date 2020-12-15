package model

import (
	"github.com/whileW/enze-global/utils"
	"time"
)

//群组消息
type GroupMsg struct {
	ID 				int			`gorm:"primary_key"`
	MsgId 			string							//msg id
	GId 			string		`sql:"index"`		//group id
	SenderId		string		`sql:"index"`
	Time			time.Time			//send time
	Content 		string				//send msg
}

//群组用户表
type GroupUser struct {
	utils.BaseModel
	GId 			string		`sql:"index"`	//group id
	UId 			string		`sql:"index"`	//user id
	LastAckMsgId	string			//最后一条收到的消息id
}
