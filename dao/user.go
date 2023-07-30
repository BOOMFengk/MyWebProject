package dao

import (
	"MyWebProject/models"
	"log"
)

//type User struct {
//	Uid      int       `json:"uid"`
//	UserName string    `json:"userName"`
//	Passwd   string    `json:"passwd"`
//	Avatar   string    `json:"avatar"`
//	CreateAt time.Time `json:"create_at"`
//	UpdateAt time.Time `json:"update_at"`
//}
//
//type UserInfo struct {
//	Uid      int    `json:"uid"`
//	UserName string `json:"userName"`
//	Avatar   string `json:"avatar"`
//}

func GetUserNameById(Id int) string {
	rows := DB.QueryRow("select user_name from blog_user where uid = ?", Id)
	if rows.Err() != nil {
		log.Println(rows.Err())
	}

	var userName string
	_ = rows.Scan(&userName)
	return userName
}

func GetUser(userName, passwd string) *models.User {
	rows := DB.QueryRow("select * from blog_user where user_name = ? and passwd = ?", userName, passwd)
	if rows.Err() != nil {
		log.Println(rows.Err())
	}
	var user = &models.User{}
	err := rows.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.UpdateAt, &user.Avatar)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user

}
