> gorm是一个基于go语言的ORM(object relation mapping)框架

什么是ORM？
- 简单来说就是数据库记录与结构体之间的对应关系

## 连接数据库
一般把连接这种初始化操作放在init函数中
```go
var mydb *gorm.DB //定义全局变量便于操作
func init() {  
  
    userName := "root"          //用户名  
    password := "root"          //密码  
    address := "127.0.0.1:3306" //数据库地址  
    dbName := "gormLearning"    //数据库名称  
    //数据库连接初始化 
    //dsn要严格遵循格式，详情见文档 
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, address, dbName)  
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{  
		//可以用gorm.config对连接进行设置
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
```
根据结构体自动生成表
```go
mydb.AutoMigrate(&Student{})
```
## 日志
>日志会在终端输出每个数据库操作实际的sql语句 

init()时设置
```go
//日志初始化  
//全局生效
mydbLogger = logger.Default.LogMode(logger.Info) //设置默认日志等级为info  

db, err := gorm.Open(mysql.Open(...), &gorm.Config{  
		...
       //[日志设置-1] 全局logger设置  
       //Logger: mydbLogger,  
        ...
    })  
```
利用session设置
```go
//[日志设置-2] 通过创建session选择是否写日志  
mydb = mydb.Session(&gorm.Session{  
    Logger: mydbLogger,})
```
Debug模式
```go
//[日志设置-3] debug 模式  
mydb.Debug().AutoMigrate(&Student{})
```
---
## 单表操作
属性
```go
type Student struct {  
    //利用标签可以设置字段属性  
    //具体有什么属性可以查看文档
    Name   string `gorm:"size:16;comment:姓名"`  
    Age    int    `gorm:"size:3;comment:年龄"`  
    Id     uint   `gorm:"size:10;comment:ID"` //默认为主键  
    Gender bool   `gorm:"size:2"`  
    //指针的作用是让字段可以设置为空值  
    Email *string `gorm:"size:128;comment:邮箱"`  
}
```
勾子函数
> gorm默认提供了几个钩子函数用于重写
```go
func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {  
    fmt.Printf("hook function before create\n")  
    eml := "s213131@qq.com"  
    s.Email = &eml  
    return nil  
}
```
### 查询
单条查询
```go
mydb.Take(&student)  
fmt.Println(student)  
//数据不干净可能导致查询出错 
//重制
student = Student{}  
//?是占位符  
mydb.Take(&student, "name=?", "张3")  
//防止sql注入，切忌自己拼接字符串  
//sql注入原理就是字符串转义  
fmt.Println(student)
```
查询多条记录
- find
```go
var studentList []Student  
count := mydb.Find(&studentList).RowsAffected  //RowsAffected 记录查询数量
for _, student := range studentList {  
    fmt.Println(student)  
}  
fmt.Printf("一共有%d条记录:", count)
```
### 更新
单个记录所有字段的更新
```go
var saveStudent Student 
saveStudent.Name = "帕鲁"  
saveStudent.Id = 1  
saveStudent.Age = 100  
mydb.Save(&saveStudent)  

```
批量更新
```go
mydb.Find(&studentList, "age=?", 10).Update("name", "test")
//将age=10的全部记录都name都修改为test
```
- select
```go
//示例：select选择单个字段更新  
mydb.Find(&studentList, "age=?", 10).Select("email").Updates(map[string]interface{}{  
    "name":  "hello",  
    "age":   18,  
    "email": newEmail})
```
 或者使用struct更新
> 注意 使用 struct 更新时, GORM 将只更新非零值字段。 可以用 map 来更新属性，或者使用 Select 声明字段来更新

删除
```go
var student Student  
mydb.Find(&studentList).Delete(&student)  
  
mydb.Delete(&student, []int{1, 2, 3})  
// DELETE FROM users WHERE id IN (1,2,3);
```

### 插入
记录的单条插入
```go
email := "30872223@qq.com"  
err := mydb.Create(&Student{  
    Name:   "张三",  
    Age:    18,    
    Gender: true,    
    Email:  &email,
    }).Error  
if err != nil {  
    fmt.Println(err)
```
记录的批量插入
```go
var stuList []Student  
for i := 0; i < 10; i++ {  
    stuList = append(stuList, Student{  
       Name:   fmt.Sprintf("张%d", i+11),  
       Age:    i + 1,  
       Gender: true,  
       Email:  &email,  
    })  
}  
err := mydb.Create(&stuList).Error  
if err != nil {  
    fmt.Println(err)  
}
```

## 高级查询
gorm也支持类似sql里where的方法
```go
var studentList []Student  
mydb.Where("name like ?", "张%").Find(&studentList)
//_替换单个内容，%替换串内容
```
