/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 15:21:49
 * @LastEditTime: 2020-03-06 13:11:53
 * @FilePath: /XGBlog/app/model/articles.go
 * @Description:
 */

package model

// Articles 文章表
type Articles struct {
	ArticleID   int    `gorm:"primary_key;column:articleID;type:int(10);not null" json:"ArticleID" ` // 文章ID
	CateID      int    `gorm:"column:cateID;type:int(10);not null" json:"CateID"`                    // 所属分类ID
	Title       string `gorm:"column:title;type:varchar(128);not null" json:"Title"`                 // 标题
	Status      int    `gorm:"column:status;type:tinyint(2);not null" json:"Status"`                 // 状态：1-公开、2-私密
	Description string `gorm:"column:description;type:varchar(512);not null" json:"Description"`     // 简介、描叙
	Keywords    string `gorm:"column:keywords;type:varchar(256);not null" json:"Keywords"`           // 关键词，用英文逗号隔开
	ImgPath     string `gorm:"column:imgPath;type:varchar(256);not null" json:"ImgPath"`             // 图片
	OpID        int    `gorm:"column:opID;type:int(10);not null" json:"OpID"`                        // 操作人id
	OpUser      string `gorm:"column:opUser;type:varchar(32);not null" json:"OpUser"`                // 操作人显示帐号
	CreateTime  string `gorm:"column:createTime;type:datetime;not null" json:"CreateTime"`           // 创建时间
	ModifyTime  string `gorm:"column:modifyTime;type:datetime;not null" json:"ModifyTime"`           // 修改时间
}

// TableName 数据库表的名称
func (Articles) TableName() string {
	return "mb_articles"
}

// ArticlesCate 文章类型表
type ArticlesCate struct {
	CateID      int    `gorm:"primary_key;column:cateID;type:int(10);not null" json:"CateID"` // 分类ID
	Name        string `gorm:"column:name;type:varchar(64);not null" json:"Name"`             // 分类名
	CName       string `gorm:"column:cName;type:varchar(64);not null" json:"CName"`           // 分类中文名
	Description string `gorm:"column:description;type:varchar(512);not null"`                 // 简介、描叙
	ParentID    int    `gorm:"column:parentID;type:int(10);not null"`                         // 上级ID
	Orderby     int    `gorm:"column:orderby;type:int(10);not null"`                          // 排序
	OpID        int    `gorm:"column:opID;type:int(10);not null"`                             // 操作人id
	OpUser      string `gorm:"column:opUser;type:varchar(32);not null"`                       // 操作人显示帐号
	CreateTime  string `gorm:"column:createTime;type:datetime;not null"`                      // 创建时间
	ModifyTime  string `gorm:"column:modifyTime;type:datetime;not null"`                      // 修改时间
}

// TableName 不知道什么玩意
func (ArticlesCate) TableName() string {
	return "mb_articles_cate"
}

// ArticlesContents 文章内容表
type ArticlesContents struct {
	ID        int    `gorm:"primary_key;column:id;type:int(10);not null" `                  // 自增ID
	ArticleID int    `gorm:"index;column:articleID;type:int(10);not null" json:"ArticleID"` // 文章ID
	ShowType  int    `gorm:"column:showType;type:tinyint(2);not null" json:"ShowType"`      // 内容展示类型：1-html、2-markdown
	Contents  string `gorm:"column:contents;type:text;not null" json:"Contents"`            // 文章内容
}

// TableName 不知道啥玩意
func (ArticlesContents) TableName() string {
	return "mb_articles_contents"
}

//GetShowTypeName 获取文章显示类型
func (a ArticlesContents) GetShowTypeName() string {
	showTypeName := ""
	switch a.ShowType {
	case 1:
		showTypeName = "html"
	case 2:
		showTypeName = "markdown"
	default:
		showTypeName = ""
	}
	return showTypeName
}
