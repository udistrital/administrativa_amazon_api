package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ActividadesCentroCostosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AdicionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AmparoPolizaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:AnulacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CambioSupervisorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CargoSupervisorTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CatalogoTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosCentroCostosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CentroCostosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CesionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoArrendamientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCanceladoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoCpsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "InsertarContratos",
			Router: `/InsertarContratos`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoLiquidadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ContratoSuscritoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaSICController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:DependenciaTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspaciosfisicosSICController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EstadoSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:GrupoTipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:JefeDependenciaTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ModificacionContractualController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:NovedadContractualController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadorGastoTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:OrdenadoresController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroNomenclaturaDianController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroTipoConformacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametroUnidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PersonaEscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PersonaEscalafonController"],
		beego.ControllerComments{
			Method: "GetAllPregrado",
			Router: `persona_escalafon_pregrado/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PersonaEscalafonController"],
		beego.ControllerComments{
			Method: "GetAllPosgrado",
			Router: `persona_escalafon_postgrado/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PlantillaMinutaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetAllContratado",
			Router: `/Contratado/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idResolucion/:idPersona`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:PucTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ReduccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "ResolucionTemplate",
			Router: `/ResolucionTemplate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "CancelarResolucion",
			Router: `/CancelarResolucion/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "RestaurarResolucion",
			Router: `/RestaurarResolucion/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GenerarResolucion",
			Router: `/GenerarResolucion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:RubroSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SedesSICController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:SuspensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TerminacionAnticipadaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "InsertarVinculaciones",
			Router: `/InsertarVinculaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_amazon_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
