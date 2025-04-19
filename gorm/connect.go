package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var mydb *gorm.DB
var mydbLogger logger.Interface

func init() {

	userName := "root"          //用户名
	password := "root"          //密码
	address := "127.0.0.1:3306" //数据库地址
	dbName := "gormLearning"    //数据库名称
	//日志初始化
	mydbLogger = logger.Default.LogMode(logger.Info) //设置默认日志等级为info

	//数据库连接初始化
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, address, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{

		NamingStrategy: schema.NamingStrategy{
			//对命名进行控制
			SingularTable: false, //单数表名
			NoLowerCase:   false, //禁止小写转换
		},
		//[日志设置-1] 全局logger设置
		//Logger: mydbLogger,
	})
	if err != nil {
		fmt.Println("数据库连接错误")
		return
	}
	mydb = db

}

/*
func main() {
	[日志设置-2] 通过创建session选择是否写日志
	mydb = mydb.Session(&gorm.Session{
		Logger: mydbLogger,
	})
	//[日志设置-3] debug 模式
	//mydb.Debug().AutoMigrate(&Student{})
	//根据 Student 结构体的定义自动创建或更新数据库中的表结构
	err := mydb.AutoMigrate(&Student{}) //AutoMigrate 只增加不修改
	if err != nil {
		fmt.Println("Migrate err", err)
		return
	}
}
*/
