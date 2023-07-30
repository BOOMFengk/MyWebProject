package service

import (
	"MyWebProject/dao"
	"MyWebProject/models"
	"MyWebProject/utils"
	"errors"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "mszlu")
	user := dao.GetUser(userName, passwd)

	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid

	token, err := utils.Award(&uid)
	if err != nil {
		if user == nil {
			return nil, errors.New("token未生成")
		}
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, err

}
