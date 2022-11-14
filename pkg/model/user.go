package model

type User struct {
	Id       int    `gorm:"column:id;primary_key;AUTO_INCREMENT;not null;comment:用户ID" json:"id"`
	Username string `gorm:"column:username;comment:用户名" json:"username"`
	Password string `gorm:"column:password;comment:密码" json:"password"`
	Email    string `gorm:"column:email;comment:邮箱;size:128" json:"email"`
	NickName string `gorm:"column:nick_name;comment:用户昵称;size:128" json:"nick_name"`
	Status   *bool  `gorm:"type:tinyint(1);default:true;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"`
}

func (u User) TableName() string {
	return "users"
}
