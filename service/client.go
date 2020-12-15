package service

import (
	"container/list"
	"fmt"
	"time"
)

type Client struct {
	Connection//必须放在结构体首部
	*MessageHandleAck
}

func NewClient(conn iconn) *Client {
	client := new(Client)

	client.conn = conn
	client.def_codec = codecs["json"]
	client.is_live = true
	client.wt = make(chan *Message, 300)
	client.messages = list.New()

	//ping、ack
	client.MessageHandleAck = &MessageHandleAck{&client.Connection}

	return client
}
func (c *Client)Run()  {
	go c.Write()
	go c.Read()
}

func (c *Client)Write()  {
	for c.is_live {
		select {
		case msg :=<- c.wt:
			m,err := c.def_codec.Encode(msg)
			if err != nil {
				fmt.Println("write msg encode err:",err)
				break
			}
			c.conn.Write(m)
		}
	}
}
func (c *Client)Read()  {
	for c.is_live {
		t1 := time.Now().Unix()
		msgb := c.conn.Read()
		t2 := time.Now().Unix()
		if t2 - t1 > 6*60 {
			fmt.Println("client:%d socket read timeout:%d %d", c.uid, t1, t2)
		}
		if msgb != nil {
			msg := Message{}
			if err := c.def_codec.Decode(msgb,&msg);err != nil{
				//消息解码失败
				c.MessageHandleAck.SendAckMsg("",Ack_Para_Codec_Err,"消息解码失败："+err.Error())
				continue
			}
			//处理消息
			msg.timestamp = time.Now().Unix()
		}
	}
}

func (c *Client)HandleMessage(msg *Message)  {
	c.MessageHandleAck.HandleMessage(msg)
	switch msg.Cmd {
	case CmdAuth:
		if str,ok := msg.Body.(string);ok {
			t,err := ParseTokenUser(str)
			if err != nil {
				c.SendAckMsg(msg.UUID,Ack_Auth_Fail,err.Error())
				break
			}
			c.is_auth = true
			c.uid = t.ID
		}else {
			c.SendAckMsg(msg.UUID,Ack_Para_Codec_Err,"auth body must be string")
		}
	}
}