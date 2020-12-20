package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)
type customer_login struct{
	Customer_id int   `gorm:"column:customer_id"`
	Login_name string	`gorm:"column:login_name"`
	Password string	`gorm:"column:password"`
	User_stats int	`gorm:"column:user_stats"`
	Modified_time int	`gorm:"column:modified_time"`
}

func main(){
	var err error
	var SDB1 *gorm.DB
	SDB1, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ybj_userdb?charset=utf8")
	SDB1.SingularTable(true)
	if err != nil {
		fmt.Println("打开数据库失败",err)

		//Logs.Error("数据库来链接失败")
	}
	fmt.Println("打开数据库成功")
	data := make([]customer_login,0)
	SDB1.Debug().Raw("SELECT * FROM customer_login;").Scan(&data)
	for k,v := range data {
		fmt.Println(k,v)
	}
	SendData,_ := json.Marshal(data)
	fmt.Println(string(SendData))
}