package controllers

import (
	"github.com/udistrital/administrativa_amazon_api/models"
	"github.com/astaxie/beego"
)

//  ProveedorContratoPersonaController operations for Proveedor_contrato_persona
type ProveedorContratoPersonaController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProveedorContratoPersonaController) URLMapping() {
	c.Mapping("Proveedor_Contrato_Persona", c.ProveedorContratoPersona)
}


// ProveedorContratoPersona ...
// @Title ProveedorContratoPersona
// @Description create ProveedorContratoPersona
// @Param	body		body 	models.ProveedorContratoPersona	true		"body for ProveedorContratoPersona content"
// @Success 201 {int} models.ProveedorContratoPersona
// @Failure 403 body is empty
// @router /:vigencia [get]
func (c *ProveedorContratoPersonaController) ProveedorContratoPersona() {
	vigencia := c.Ctx.Input.Param(":vigencia")
	respuesta, err := models.VigenciaProveedorContratoPersona(vigencia)
	if err != nil{	
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = respuesta
	}
	c.ServeJSON()
}
