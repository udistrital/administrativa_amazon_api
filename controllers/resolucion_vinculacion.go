package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_amazon_api/models"
)

type ResolucionVinculacionController struct {
	beego.Controller
}

func (c *ResolucionVinculacionController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)

}

// GetAll ...
// @Title Get All
// @Description get ResolucionVinculacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ResolucionVinculacionDocente
// @Failure 403
// @router / [get]
func (c *ResolucionVinculacionController) GetAll() {
	listaResoluciones := models.GetAllResolucionVinculacion()
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = listaResoluciones
	c.ServeJSON()
}
