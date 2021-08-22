package model

// BlogArticle 文章管理
type BlogArticle struct {
	ID         uint   `gorm:"primaryKey;column:id;type:int unsigned;not null" json:"-"`
	TagID      uint   `gorm:"column:tag_id;type:int unsigned;default:0" json:"tagId"` // 标签ID
	Title      string `gorm:"column:title;type:varchar(100);default:''" json:"title"` // 文章标题
	Desc       string `gorm:"column:desc;type:varchar(255);default:''" json:"desc"`   // 简述
	Content    string `gorm:"column:content;type:text" json:"content"`
	CreatedOn  int    `gorm:"column:created_on;type:int" json:"createdOn"`
	CreatedBy  string `gorm:"column:created_by;type:varchar(100);default:''" json:"createdBy"`   // 创建人
	ModifiedOn uint   `gorm:"column:modified_on;type:int unsigned;default:0" json:"modifiedOn"`  // 修改时间
	ModifiedBy string `gorm:"column:modified_by;type:varchar(255);default:''" json:"modifiedBy"` // 修改人
	DeletedOn  uint   `gorm:"column:deleted_on;type:int unsigned;default:0" json:"deletedOn"`
	State      uint8  `gorm:"column:state;type:tinyint unsigned;default:1" json:"state"` // 状态 0为禁用1为启用
}

// TableName get sql table name.获取数据库表名
func (m *BlogArticle) TableName() string {
	return "blog_blog_article"
}

// BlogArticleColumns get sql column name.获取数据库列名
var BlogArticleColumns = struct {
	ID         string
	TagID      string
	Title      string
	Desc       string
	Content    string
	CreatedOn  string
	CreatedBy  string
	ModifiedOn string
	ModifiedBy string
	DeletedOn  string
	State      string
}{
	ID:         "id",
	TagID:      "tag_id",
	Title:      "title",
	Desc:       "desc",
	Content:    "content",
	CreatedOn:  "created_on",
	CreatedBy:  "created_by",
	ModifiedOn: "modified_on",
	ModifiedBy: "modified_by",
	DeletedOn:  "deleted_on",
	State:      "state",
}

// BlogAuth [...]
type BlogAuth struct {
	ID       uint   `gorm:"primaryKey;column:id;type:int unsigned;not null" json:"-"`
	Username string `gorm:"column:username;type:varchar(50);default:''" json:"username"` // 账号
	Password string `gorm:"column:password;type:varchar(50);default:''" json:"password"` // 密码
}

// TableName get sql table name.获取数据库表名
func (m *BlogAuth) TableName() string {
	return "blog_blog_auth"
}

// BlogAuthColumns get sql column name.获取数据库列名
var BlogAuthColumns = struct {
	ID       string
	Username string
	Password string
}{
	ID:       "id",
	Username: "username",
	Password: "password",
}

// BlogTag 文章标签管理
type BlogTag struct {
	ID         uint   `gorm:"primaryKey;column:id;type:int unsigned;not null" json:"-"`
	Name       string `gorm:"column:name;type:varchar(100);default:''" json:"name"`              // 标签名称
	CreatedOn  uint   `gorm:"column:created_on;type:int unsigned;default:0" json:"createdOn"`    // 创建时间
	CreatedBy  string `gorm:"column:created_by;type:varchar(100);default:''" json:"createdBy"`   // 创建人
	ModifiedOn uint   `gorm:"column:modified_on;type:int unsigned;default:0" json:"modifiedOn"`  // 修改时间
	ModifiedBy string `gorm:"column:modified_by;type:varchar(100);default:''" json:"modifiedBy"` // 修改人
	DeletedOn  uint   `gorm:"column:deleted_on;type:int unsigned;default:0" json:"deletedOn"`
	State      uint8  `gorm:"column:state;type:tinyint unsigned;default:1" json:"state"` // 状态 0为禁用、1为启用
}

// TableName get sql table name.获取数据库表名
func (m *BlogTag) TableName() string {
	return "blog_blog_tag"
}

// BlogTagColumns get sql column name.获取数据库列名
var BlogTagColumns = struct {
	ID         string
	Name       string
	CreatedOn  string
	CreatedBy  string
	ModifiedOn string
	ModifiedBy string
	DeletedOn  string
	State      string
}{
	ID:         "id",
	Name:       "name",
	CreatedOn:  "created_on",
	CreatedBy:  "created_by",
	ModifiedOn: "modified_on",
	ModifiedBy: "modified_by",
	DeletedOn:  "deleted_on",
	State:      "state",
}
