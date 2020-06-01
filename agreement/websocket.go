package agreement

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type websocket_struct struct {
	conn 			*websocket.Conn
	send_chan		chan []byte
	read_chan 		chan []byte
	death_chan 		chan int
}

const AgreementName_WebSocket  =	"websocket"
const (
	maxMessageSize = 512
)
var(
	newline = []byte{'\n'}
)
//注册websocket
func RegisterWebSocket(conn *websocket.Conn) *websocket_struct {
	websocket := &websocket_struct{
		conn:conn,
		send_chan:make(chan []byte,100),
		read_chan:make(chan []byte,100),
		death_chan:make(chan int,100),
	}
	go websocket.write_pump()
	go websocket.read_pump()
	return websocket
}
func (c *websocket_struct)Send(msg []byte)  {
	c.send_chan<-msg
}
func (c *websocket_struct)GetReadChan() chan []byte {
	return c.read_chan
}
func (c *websocket_struct)GetDeathChan() chan int {
	return c.death_chan
}
func (c *websocket_struct)Disconnect() {
	write_all_msg(c)
	if c.conn != nil {
		c.conn.Close()
	}
}
//读消息
func (c *websocket_struct) read_pump() {
	defer func() {
		//u.User.OffOnline()
		c.death_chan <- 1
		fmt.Println("read close")
	}()
	c.conn.SetReadLimit(maxMessageSize)
	//c.conn.SetReadDeadline(time.Now().Add(pongWait))
	//c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		//fmt.Println(123)
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error:%v",err)
				//log.Printf("error: %v", err)
			}
			break
		}
		//处理消息
		fmt.Println(string(message))
		c.read_chan<-message
	}
}
//发送消息
func (c *websocket_struct) write_pump() {
	timer := time.NewTimer(time.Second)
	defer func() {
		//UnRegisterUser(u,"WritePump")
		c.death_chan <- 1
		timer.Stop()
		fmt.Println("write close")
	}()
	for {
		select {
		case <-timer.C:
			//u.Conn.SetWriteDeadline()
			n := len(c.send_chan)
			if n>0 {
				w, err := c.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}
				for i := 0; i < n; i++ {
					w.Write([]byte(<-c.send_chan))
					if i != n-1 {
						w.Write(newline)
					}
				}
				if err := w.Close(); err != nil {
					fmt.Println("close:"+err.Error())
					return
				}
			}
			timer.Reset(time.Second)
		}
	}
}
func write_all_msg(c *websocket_struct)  {
	n := len(c.send_chan)
	if n>0 {
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		for i := 0; i < n; i++ {
			w.Write([]byte(<-c.send_chan))
			if i != n-1 {
				w.Write(newline)
			}
		}
		if err := w.Close(); err != nil {
			fmt.Println("close:"+err.Error())
			return
		}
	}
}