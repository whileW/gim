package examples

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"im/msg"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func main()  {
	flag.Parse()
	r := gin.Default()
	r.GET("/websocket",websocket_g)
	//r.GET("/GetDialogueListByUser",gin2.GetDialogueListByUser)
	r.GET("/test", func(c *gin.Context) {
		c.Writer.Write([]byte("213"))
	})
	r.Run(":"+*addr)
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}
func websocket_g(c *gin.Context)  {
	page := c.Query("page")
	page_para := c.Query("page_para")
	platform := c.Query("platform")
	token := c.Request.Header.Get("Sec-Websocket-Protocol")
	//升级get请求为webSocket协议
	header := http.Header{}
	header.Add("Sec-Websocket-Protocol",token)
	conn, err := upGrader.Upgrade(c.Writer, c.Request, header)
	if err != nil {
		return
	}
	msg.RegisterWebSocket(conn,token,page,page_para,platform)
}