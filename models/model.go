package models

import(
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "github.com/astaxie/beedb"
   "time"
)

type Blog struct{
	Id int `PK`
	Title string
	Content string
	Created time.Time
}

func GetLink() beedb.Model {
	db, err := sql.Open("mysql", "root:root@/sel")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}

func GetAll() (blogs []Blog) {
	db := GetLink()
	db.FindAll(&blogs)
    return
}

func GetBlog(id int) (blog Blog) {
	db := GetLink()
	db.Where("id=?", id).Find(&blog)
	return
}

func SaveBlog(blog Blog) {
	db := GetLink()
	db.Save(&blog)
	return 
}

func DelBlog(id int) {
	db := GetLink()
	db.SetTable("blog").Where("id=?", id).DeleteRow()
	return
}