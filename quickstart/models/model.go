package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	Name string
	Pwd string
	Artilce []*Article `orm:"rel(m2m)"`
}

type Article struct{
	Id int `orm:"pk;auto"`
	Aname string `orm:"size(20);null"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string
	Aimg string
	ArticleType *ArticleType `orm:"rel(fk)"`
	User []*User `orm:"reverse(many)"`
}

type ArticleType struct {
	Id int `orm:"pk;auto"`
	TypeName string `orm:size(20)`
	Articles []*Article `orm:"reverse(many)"`
}

func init()  {
	orm.RegisterDataBase("default","mysql","root:sxj54686521@@tcp(miqilin5212.mysql.rds.aliyuncs.com:3306)/test?charset=utf8")
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	orm.RunSyncdb("default",false,true)
}
