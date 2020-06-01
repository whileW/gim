package user

import (
	"im/utils"
)

const (
	Page_Im_Chat_Room = "im_chat_room"		//聊天室
)

func RegisterUserByPage(page,page_para string,user_id string) error {
	var err error
	switch page {
	case Page_Im_Chat_Room:
		err = AddUserByImChatRoomInRedis(page,page_para,user_id,utils.LocalIP())
		break
	}
	return err
}
func OfflineUserByPage(page,page_para string,user_id string) error {
	var err error
	switch page {
	case Page_Im_Chat_Room:
		err = DelUserByImChatRoomInRedis(page,page_para,user_id,utils.LocalIP())
		break
	}
	return err
}

func AddUserByImChatRoomInRedis(page,page_para,user_id,local_ip string) error {
	r := utils.GetRedisPool().Get()
	_,err := r.Do("SADD",page+":"+page_para,
		user_id+":"+local_ip)
	return err
}
func DelUserByImChatRoomInRedis(page,page_para,user_id,local_ip string) error {
	r := utils.GetRedisPool().Get()
	_,err := r.Do("SREM",page+":"+page_para,
		user_id+":"+local_ip)
	return err
}
func FindUserByPageInRedis(page,page_para string) ([]string,error) {
	r := utils.GetRedisPool().Get()
	return utils.Strings(r.Do("SMEMBERS",page+":"+page_para))
}