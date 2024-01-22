package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/agenda_mid_prueba_v1/helpers"
)

// Contacto_parametrosController operations for Contacto_parametros
type Contacto_parametrosController struct {
	beego.Controller
}

// URLMapping ...
func (c *Contacto_parametrosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Contacto_parametros
// @Param	body		body 	models.Contacto_parametros	true		"body for Contacto_parametros content"
// @Success 201 {object} models.Contacto_parametros
// @Failure 403 body is empty
// @router / [post]
func (c *Contacto_parametrosController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Contacto_parametros by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Contacto_parametros
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Contacto_parametrosController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Contacto_parametros
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Contacto_parametros
// @Failure 403
// @router / [get]
func (c *Contacto_parametrosController) GetAll() {
	defer helpers.ErrorController(c.Controller, "Contacto_parametrosController")

	if v, err := helpers.ListarContactosParametros(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status": "200", "Message": "Request succesful", "Data": v}
	}else {
		panic(err)
	}

	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Contacto_parametros
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Contacto_parametros	true		"body for Contacto_parametros content"
// @Success 200 {object} models.Contacto_parametros
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Contacto_parametrosController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Contacto_parametros
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Contacto_parametrosController) Delete() {

}
