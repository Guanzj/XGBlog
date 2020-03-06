/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 15:21:49
 * @LastEditTime: 2020-03-06 15:21:55
 * @FilePath: /XGBlog/app/model/admins.go
 * @Description:
 */

package model

// Admins 后台管理员的结构体
// type Admins struct {
// 	AdminID  int32  `json:"admin_id"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Status   int8   `json:"status"`
// }

// Admins 管理员表
type Admins struct {
	AdminID    int32  `gorm:"primary_key;column:adminID;type:int(10);not null" json:"admin_id"` // ID
	Username   string `gorm:"column:username;type:varchar(64);not null" json:"username"`        // 用户名
	Password   string `gorm:"column:password;type:varchar(64);not null" json:"password"`        // 密码
	Status     int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`             // 状态：1-正常，2-禁用
	OpID       int    `gorm:"column:opID;type:int(10);not null"`                                // 操作人id
	OpUser     string `gorm:"column:opUser;type:varchar(32);not null"`                          // 操作人显示帐号
	CreateTime string `gorm:"column:createTime;type:datetime;not null"`                         // 创建时间
	ModifyTime string `gorm:"column:modifyTime;type:datetime;not null"`                         // 修改时间
}

// TableName 不知道什么玩意？？？
func (Admins) TableName() string {
	return "mb_admins"
}
