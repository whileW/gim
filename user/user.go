package user

import (
	"im/agreement"
	"sync"
	"time"
)

type UserInter interface {
	GetReqId() 			string
	GetUserName() 		string
	GetUserId()			string
	GetUserPage()		string
	GetUserPagePara()	string
	GetUserPatform()	string
	OffLine()
}

type user struct {
	ReqId 			string			`json:"req_id"`
	Token 			string			`json:"token"`				//token
	UserId 			string			`json:"user_id"`			//用户id
	UserName 		string			`json:"user_name"`			//用户名称
	OnLineTime 		time.Time		`json:"go_in_time"`			//在线时间
	Page 			string			`json:"page"`				//在哪个页面
	PagePara 		string			`json:"page_para"`			//页面参数
	Platform		string			`json:"platform"`			//平台
	SignalName 		string			`json:"signal_name"`		//通信协议名称
	Conn 			agreement.AgreementInter	`json:"conn"`				//通信链接
	State 			int 			`json:"state"`				//状态 0未登录 1正常 2下线
}

var user_pool = make(map[string]*[]user)
var user_pool_lock = new(sync.RWMutex)
//注册用户在线
func RegisterUser(ReqId,UserId,UserName,Page,PagePara,Platform,SignalName string,Conn agreement.AgreementInter) (*user,error) {
	user_pool_lock.Lock()
	defer user_pool_lock.Unlock()

	u := &user{ReqId:ReqId,UserId:UserId,UserName:UserName,PagePara:PagePara,
		OnLineTime:time.Now(),Page:Page,Platform:Platform,
		SignalName:SignalName,Conn:Conn}
	if user_pool[UserId] == nil {
		user_pool[UserId] = &[]user{}
	}
	*user_pool[UserId] = append(*user_pool[UserId], *u)
	err := RegisterUserByPage(Page,PagePara,UserId)
	return u,err
}
func (u *user) GetUserName() 	string{
	return u.UserName
}
func (u *user) GetReqId()string{
	return u.ReqId
}
func (u *user) GetUserId()string{
	return u.UserId
}
func (u *user) GetUserPage()string{
	return u.Page
}
func (u *user) GetUserPagePara()	string{
	return u.PagePara
}
func (u *user) GetUserPatform()string {
	return u.Platform
}
//下线
func (u *user) OffLine() {
	user_pool_lock.Lock()
	defer user_pool_lock.Unlock()
	if user_pool[u.UserId] == nil {
		return
	}
	OfflineUserByPage(u.GetUserPage(),u.GetUserPagePara(),u.GetUserId())
	for i:=0;i< len(*user_pool[u.UserId]);i++ {
		user := (*user_pool[u.UserId])[i]
		if user.ReqId == u.ReqId {
			user.Conn.Disconnect()
			user.State = 2
			*user_pool[u.UserId] = append((*user_pool[u.UserId])[:i], (*user_pool[u.UserId])[i+1:]...)
		}
	}
}

func GetUserById(id string,page,page_para string) *[]user {
	users := &[]user{}
	user_pool_lock.RLock()
	defer user_pool_lock.RUnlock()
	if user_pool[id] != nil {
		for _,t :=range *user_pool[id] {
			if t.Page == page && (page_para == "" || t.PagePara == page_para) {
				*users = append(*users, t)
			}
		}
	}
	return users
}