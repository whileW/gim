package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-global/utils/resp"
	"gim/model"
	"gim/model/request"
	"gim/service"
)

func Login(c *gin.Context)  {
	reqStruct := &request.UserLocalLoginReqStruct{}
	c.ShouldBindJSON(reqStruct)

	u := model.User{}
	if err := u.Login(reqStruct.LoginName,reqStruct.Password);err != nil{
		resp.FailWithMessage(c,err.Error())
		return
	}
	token,err := service.GenerateTokenUser(u.UUID,u.Name,u.HeadImg)
	if err != nil {
		resp.FailWithMessage(c,err.Error())
		return
	}
	resp.OkWithData(c,token)
	return
}

func Reg(c *gin.Context)  {
	reqStruct := &request.UserRegReqStruct{}
	c.ShouldBindJSON(reqStruct)
	u := model.User{
		Phone:reqStruct.LoginName,
		Password:reqStruct.Password,
		Name:reqStruct.NickName,
		HeadImg:reqStruct.HeadImg,
	}
	if err := u.Reg();err != nil {
		resp.FailWithMessage(c,err.Error())
		return
	}
	token,err := service.GenerateTokenUser(u.UUID,u.Name,u.HeadImg)
	if err != nil {
		resp.Ok(c)
		return
	}
	resp.OkWithData(c,token)
	return
}