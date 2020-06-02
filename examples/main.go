package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"github.com/whileW/im/msg"
	"log"
	"net/http"
	"net/url"
)

var addr = flag.String("addr", ":8080", "http service address")
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket_g(w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//var upGrader = websocket.Upgrader{
//	CheckOrigin: func (r *http.Request) bool {
//		return true
//	},
//}
func websocket_g(w http.ResponseWriter, r *http.Request)  {
	page := valuesGetDefault(r.Form,"page","im_chat_room")
	page_para := valuesGetDefault(r.Form,"page_para","1")
	platform := valuesGetDefault(r.Form,"platform","pc")
	token := r.Header.Get("Sec-Websocket-Protocol")
	//升级get请求为webSocket协议
	header := http.Header{}
	header.Add("Sec-Websocket-Protocol",token)
	conn, err := upgrader.Upgrade(w, r, header)
	if err != nil {
		return
	}
	msg.RegisterWebSocket(conn,token,page,page_para,platform)
}

func valuesGetDefault(values url.Values, key, defaultValue string) string {
	v := values.Get(key)
	if v == "" {
		return defaultValue
	} else {
		return v
	}
}