package dao

import (
	"log"

	"gin_blogs_first/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 接口
type Manager interface {
	//用户
	RegisterUser(user *model.User) bool //拿到了model里面的User模型
	Login(username string) APIUser
	//博客
	AddEditor(editor *model.Editor) bool
	GetAllEditor() []model.Editor
	GetEditor(pid int) model.Editor
}

type manager struct {
	db *gorm.DB
}

//这个变量是上面接口的类型
var Mgr Manager

//初始化数据库
func init() {

	dsn := "blogs_first:kM7XaRNxnKjwAHPH@tcp(101.200.243.101:3306)/blogs_first?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		log.Panic("连接mysql数据库失败", err)
	}

	//根据结构体创建表
	// 用户数据表
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Panic("创建User表失败", err)
	}
	//博客文章数据表
	if err := db.AutoMigrate(&model.Editor{}); err != nil {
		log.Panic("创建Editor表失败", err)
	}

	//初始化Mgr
	Mgr = &manager{db}
}
