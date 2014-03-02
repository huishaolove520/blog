package controllers

import(
  "github.com/astaxie/beego"
  "strconv"

  "blog/models"
)

type ViewController struct{
	beego.Controller
}

func (this *ViewController) Get() {
	//inputs := this.Input
	id, _ := strconv.Atoi(this.Ctx.Input.Params[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplNames = "view.tpl"
}