package controllers

import (
	"encoding/json"
	"errors"
	"github.com/udistrital/administrativa_amazon_api/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// ResolucionController operations for Resolucion
type ResolucionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ResolucionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("CancelarResolucion", c.CancelarResolucion)
	c.Mapping("RestaurarResolucion", c.RestaurarResolucion)
	c.Mapping("GenerarResolucion", c.GenerarResolucion)
}

// Put ...
// @Title Cancelar
// @Description update the Resolucion
// @Param	id		path 	string	true		"The id you want to update"
// @Success 200 {object} models.Resolucion
// @Failure 403 :id is not int
// @router /CancelarResolucion/:id [put]
func (c *ResolucionController) CancelarResolucion() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Resolucion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.CancelarResolucion(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Put ...
// @Title Restaurar
// @Description update the Resolucion
// @Param	id		path 	string	true		"The id you want to update"
// @Success 200 {object} models.Resolucion
// @Failure 403 :id is not int
// @router /RestaurarResolucion/:id [put]
func (c *ResolucionController) RestaurarResolucion() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Resolucion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.RestaurarResolucion(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Post ...
// @Title Post
// @Description create Resolucion
// @Success 201 {int} models.Resolucion
// @Failure 403 body is empty
// @router /GenerarResolucion [post]
func (c *ResolucionController) GenerarResolucion() {
	var v models.Resolucion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.GenerarResolucion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Post ...
// @Title Post
// @Description create Resolucion
// @Param	body		body 	models.Resolucion	true		"body for Resolucion content"
// @Success 201 {int} models.Resolucion
// @Failure 403 body is empty
// @router / [post]
func (c *ResolucionController) Post() {
	var v models.Resolucion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddResolucion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Resolucion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Resolucion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ResolucionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetResolucionById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Resolucion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Resolucion
// @Failure 403
// @router / [get]
func (c *ResolucionController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllResolucion(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Resolucion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Resolucion	true		"body for Resolucion content"
// @Success 200 {object} models.Resolucion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ResolucionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Resolucion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateResolucionById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Resolucion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ResolucionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteResolucion(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// VigenciaResolucion ...
// @Title VigenciaResolucion
// @Description create VigenciaResolucion
// @Param	body		body 	models.VigenciaResolucion	true	"body for VigenciaResolucion content"
// @Success 201 {int}
// @Failure 403 body is empty
// @router /vigencia_resolucion
func (c *ResolucionController) VigenciaResolucion() {
	respuesta := models.GetVigenciaResolucion()
	if len(respuesta) == 0 {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = "Error leyendo las vigencias"
		c.ServeJSON()
	} else {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
		c.ServeJSON()
	}
}


// ResolucionEstado ...
// @Title ResolucionEstado
// @Description create ResolucionEstado
// @Param	body		body 	models.Resolucion	true	"body for Resolucion content"
// @Success 201 {int}
// @Failure 403 body is empty
// @router /resolucion_por_estado/:vigencia/:estado
func (c *ResolucionController) ResolucionEstado() {
	vigencia := c.Ctx.Input.Param(":vigencia")
	estado := c.Ctx.Input.Param(":estado")
	respuesta := models.GetResolucionEstadoVigencia(vigencia,estado)
	if len(respuesta) == 0 {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = "Error leyendo las vigencias"
		c.ServeJSON()
	} else {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
		c.ServeJSON()
	}
}