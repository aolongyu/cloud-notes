package Data

import "github.com/jinzhu/gorm"

//账户内部信息
type Account_info struct {
	UID             int32
	Account_lv      int32
	Account_exp     int32
	Weapon_shard    int32
	Star_stage      int32
	Choosen_hero_id int32
	Weapon_id       int32 //Hero_weapon
}

type Hero struct {
	Id         int32
	UID        int32
	Hero_id    int32
	Hero_lv    int32
	Hero_exp   int32
	Hero_skill int32 //choosen_skill
}
type User_info struct {
	//gorm.Model
	UID int32 `gorm:"column:uid"`
	//`gorm:"default:'1';primary_key"` //;AUTO_INCREMENT"`
	Account_name     string
	Account_password string
	User_name        string
}
type Weapon struct {
	gorm.Model
	Id              int32
	Weapon_id       int32
	Weapon_name     string
	Weapon_describe string
	Weapon_RPM      int32
	Weapon_magazine int32
	Weapon_reload   float32
	Weapon_range    int32
	Weapon_impact   int32
}
type Bangding struct {
	Id        int32
	Uid       int32
	Weapon_id int32
}
