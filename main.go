package main

import (
	"fmt"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/initialize"
	"im/model"
	"im/router"
	"im/service"
	"runtime"
)

var(
	http_port = global.GVA_CONFIG.SysSetting.HttpAddr
)

func init()  {
	init_redis()
	init_db_tables()
}
func init_redis()  {
	if global.IsHaveRedis() {
		initialize.Redis()
	}else {
		panic("请配置redis地址")
	}
}

// 注册数据库表
func init_db_tables() {
	initialize.Db()
	db := global.GVA_DB.Get(model.ImDB)
	db.SingularTable(true)		//禁用表名复数
	db.AutoMigrate(model.User{},model.GroupMsg{},model.GroupUser{})
	global.GVA_LOG.Info("register table success")
}

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go run_http()
	run_print()
	select {}
}

func run_http() {
	r := global.InitGin()
	r.GET("ws",service.WebSocketHander)		//websocket

	api := r.Group("api")
	router.InitUserRouter(api)					//用户

	if err := r.Run(":"+http_port);err != nil{
		service.HttpHealth = false
		fmt.Println("http listen err",err)
	}
}

func run_print() {
	if service.HttpHealth {
		fmt.Println("http start listen to",http_port)
	}
}