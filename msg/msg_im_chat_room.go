package msg

import (
	"im/user"
	"errors"
	"im/utils"
	"strings"
)

//消息结构
type im_chat_room_msg_struct struct {
	Msg 			string			`json:"msg"`			//消息
}

//消息处理
func ImChatRoomMsgHand(useri user.UserInter,data map[string]string) error {
	if data["msg"] == "" {
		return errors.New("please write msg")
	}
	users,err := user.FindUserByPageInRedis(useri.GetUserPage(),useri.GetUserPagePara())
	if err != nil {
		return err
	}
	//user_ids := []string{}
	for _,t := range users {
		u := strings.Split(t,":")
		if len(u)<2 {
			continue
		}
		if u[1] == utils.LocalIP() {
			us := user.GetUserById(u[0],useri.GetUserPage(),useri.GetUserPagePara())
			for _,t := range *us {
				h := FindHubByReqId(t.GetReqId())
				h.AddPushRespMsg(data["msg"],nil,useri)
			}
		}else {
			//负载均衡
		}
	}
	return nil
}
