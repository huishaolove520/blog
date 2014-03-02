package controllers

import(
  "github.com/astaxie/beego"
  "strconv"

  "blog/models"
)

type DeleteController struct{
	beego.Controller
}

func (this *DeleteController) Get() {
	//inputs := this.Input()
	id, _ := strconv.Atoi(this.Ctx.Input.Params[":id"])
	models.DelBlog(id)
	this.Ctx.Redirect(302, "/")
}