package dao

import "gin_blogs_first/model"

//添加用户
func (mgr *manager) RegisterUser(user *model.User) bool {

	//检测是否已经注册
	var userList model.User
	d := mgr.db.Where("username = ?", user.Username).Find(&userList) //这里的first(user) 为什么要传这个user 可能是表示那个表?
	if d.RowsAffected > 0 {
		return false
	} else {
		//添加数据
		mgr.db.Create(user)
		return true
	}
}

//登录
type APIUser struct {
	Id       string
	Username string
	Password string
}

func (mgr *manager) Login(username string) APIUser {
	apiUser := APIUser{}
	mgr.db.Model(&model.User{}).Where("username = ?", username).First(&apiUser) //高级查询-智能选择字段
	// mgr.db.Select("id", "username", "password").Where("username = ?", username).First(&user) //指定字段 但是结构体里有的值会变零值
	// mgr.db.Where("username = ?", username).First(&user) //不指定字段
	return apiUser
}
