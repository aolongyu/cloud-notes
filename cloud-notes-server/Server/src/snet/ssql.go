package snet

import (
	//Data "data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


var SDB *gorm.DB
/*
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//path := strings.Join([]string{"root:@tcp(172.20.80.42:3306)/g4?charset=utf8"})

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = gorm.Open("mysql", "root:123456@tcp(172.20.80.42:3306)/server?charset=utf8")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}
*/
/*
type test1 struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Password string `gorm:"not null;"`
	Status   uint   `gorm:"default:0"`
}

func main() {
	DB, err := gorm.Open("mysql", "root:123456@tcp(172.20.80.42:3306)/server?charset=utf8")

	if err != nil {
		println(err)
	}
	defer DB.Close()
	DB.SingularTable(true)
	DB.AutoMigrate(&Data.Account_info{})
	user := Data.Account_info{
		Model:       gorm.Model{},
		Id:          1,
		Uid:         1,
		Account_exp: 1,
		Account_lvl: 1,
		Hero_shard:  1,
	}

	DB.Create(&user)
	DB.Save(&user)
}
*/
