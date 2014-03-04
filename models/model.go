package models

import(
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "github.com/astaxie/beego/orm"
   "time"
)

type Blog struct{
	Id int `PK`
	Title string
	Content string
	Created time.Time
}

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/sel")
	orm.RegisterModel(new(Blog))

	name := "default"
	force := false
	verbose := false

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func GetOrm() orm.Ormer {
	o := orm.NewOrm()
	return o
}

func GetAll() (blogs []Blog) {
	db := GetOrm()
	db.QueryTable("blog").All(&blogs)
    return
}

func GetBlog(id int) (blog Blog) {
	db := GetOrm()
	db.QueryTable("blog").Filter("id", id).One(&blog)
	return
}

func SaveBlog(blog Blog) {
	db := GetOrm()
	db.Insert(&blog)
	return 
}

func DelBlog(id int) {
	db := GetOrm()
	db.QueryTable("blog").Filter("id", id).Delete()
	return
}