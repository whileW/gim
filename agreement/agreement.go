package agreement

type AgreementInter interface {
	Send([]byte)						//发送消息
	GetReadChan()	chan []byte			//获取读取消息的chan
	GetDeathChan() 	chan int			//获取连接死亡通知chan
	Disconnect()						//主动关闭连接
}