package helpers

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/agenda_mid_prueba_v1/models"
)

func ListarContactosParametros() (contactosParametros models.ContactosParametros, outputError map[string]interface{}){
	defer func ()  {
		if err := recover(); err != nil{
			outputError = map[string]interface{}{"funcion": "ListarContactosParametros", "err": err, "status": "500"}
			panic(err)
		}
	}()

	var contactos []models.Contacto
	var parametros []models.Parametro

	url := "parametro?query=TipoParametroId__CodigoAbreviacion:TV&limit=0"
	if err := GetRequestNew("UrlCrudParametros", url, &parametros); err != nil {
		logs.Error(err)
		panic(err.Error())
	}

	url = "contacto/?limit=0"
	if err := GetRequestNew("UrlCrudContactos", url, &contactos); err != nil {
		logs.Error(err)
		panic(err.Error())
	}

	contactosParametros.Parametros = parametros
	contactosParametros.Contactos = contactos

	fmt.Println("PARAMETROS", parametros)
	fmt.Println("CONTACTOS", contactos)

	return contactosParametros, outputError

	
}
