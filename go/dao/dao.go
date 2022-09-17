package dao

import (
	"log"

	"gin_blogs_first/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 接口
type Manager interface {
	RegisterUser(user *model.User) bool //拿到了model里面的User模型
	Login(username string) model.User
}

type manager struct {
	db *gorm.DB
}

//这个变量是上面接口的类型
var Mgr Manager

//初始化数据库
func init() {

	dsn := "blogs_first:kM7XaRNxnKjwAHPH@tcp(101.200.243.101:3306)/blogs_first?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("连接mysql数据库失败", err)
	}

	//初始化Mgr
	Mgr = &manager{db}

	//根据结构体创建表
	if err2 := db.AutoMigrate(&model.User{}); err2 != nil {
		log.Panic("创建User表失败", err2)
	}

}

//添加用户
func (mgr *manager) RegisterUser(user *model.User) bool {

	//必须输入
	if user.Username == "" || user.Password == "" {
		return false
	}

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
func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username = ?", username).First(&user)
	return user
}
