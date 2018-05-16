// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/administrativa_amazon_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/dedicacion",
			beego.NSInclude(
				&controllers.DedicacionController{},
			),
		),

		beego.NSNamespace("/informacion_persona_juridica_tipo_entidad",
			beego.NSInclude(
				&controllers.InformacionPersonaJuridicaTipoEntidadController{},
			),
		),

		beego.NSNamespace("/tipo_entidad",
			beego.NSInclude(
				&controllers.TipoEntidadController{},
			),
		),

		beego.NSNamespace("/resolucion_estado",
			beego.NSInclude(
				&controllers.ResolucionEstadoController{},
			),
		),

		beego.NSNamespace("/tipo_resolucion",
			beego.NSInclude(
				&controllers.TipoResolucionController{},
			),
		),

		beego.NSNamespace("/estado_resolucion",
			beego.NSInclude(
				&controllers.EstadoResolucionController{},
			),
		),

		beego.NSNamespace("/escalafon",
			beego.NSInclude(
				&controllers.EscalafonController{},
			),
		),

		beego.NSNamespace("/componente_resolucion",
			beego.NSInclude(
				&controllers.ComponenteResolucionController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion_docente",
			beego.NSInclude(
				&controllers.ResolucionVinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion",
			beego.NSInclude(
				&controllers.ResolucionVinculacionController{},
			),
		),

		beego.NSNamespace("/escalafon_persona",
			beego.NSInclude(
				&controllers.EscalafonPersonaController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/vinculacion_docente",
			beego.NSInclude(
				&controllers.VinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/informacion_interventor",
			beego.NSInclude(
				&controllers.InformacionInterventorController{},
			),
		),

		beego.NSNamespace("/informacion_persona_natural",
			beego.NSInclude(
				&controllers.InformacionPersonaNaturalController{},
			),
		),

		beego.NSNamespace("/inhabilidad",
			beego.NSInclude(
				&controllers.InhabilidadController{},
			),
		),

		beego.NSNamespace("/parametro_estandar",
			beego.NSInclude(
				&controllers.ParametroEstandarController{},
			),
		),

		beego.NSNamespace("/objeto_contratar",
			beego.NSInclude(
				&controllers.ObjetoContratarController{},
			),
		),

		beego.NSNamespace("/evaluacion",
			beego.NSInclude(
				&controllers.EvaluacionController{},
			),
		),

		beego.NSNamespace("/parametro_tipo_conformacion",
			beego.NSInclude(
				&controllers.ParametroTipoConformacionController{},
			),
		),

		beego.NSNamespace("/contrato_proveedor",
			beego.NSInclude(
				&controllers.ContratoProveedorController{},
			),
		),

		beego.NSNamespace("/contrato_evaluacion",
			beego.NSInclude(
				&controllers.ContratoEvaluacionController{},
			),
		),

		beego.NSNamespace("/parametro_nomenclatura_dian",
			beego.NSInclude(
				&controllers.ParametroNomenclaturaDianController{},
			),
		),

		beego.NSNamespace("/solicitud_cotizacion",
			beego.NSInclude(
				&controllers.SolicitudCotizacionController{},
			),
		),

		beego.NSNamespace("/informacion_proveedor",
			beego.NSInclude(
				&controllers.InformacionProveedorController{},
			),
		),

		beego.NSNamespace("/informacion_sociedad_temporal",
			beego.NSInclude(
				&controllers.InformacionSociedadTemporalController{},
			),
		),

		beego.NSNamespace("/objeto_contratar_actividad_ciiu",
			beego.NSInclude(
				&controllers.ObjetoContratarActividadCiiuController{},
			),
		),

		beego.NSNamespace("/proveedor_actividad_ciiu",
			beego.NSInclude(
				&controllers.ProveedorActividadCiiuController{},
			),
		),

		beego.NSNamespace("/proveedor_telefono",
			beego.NSInclude(
				&controllers.ProveedorTelefonoController{},
			),
		),

		beego.NSNamespace("/unidad",
			beego.NSInclude(
				&controllers.UnidadController{},
			),
		),

		beego.NSNamespace("/informacion_persona_juridica",
			beego.NSInclude(
				&controllers.InformacionPersonaJuridicaController{},
			),
		),

		beego.NSNamespace("/codigo_validacion",
			beego.NSInclude(
				&controllers.CodigoValidacionController{},
			),
		),

		beego.NSNamespace("/informacion_supervisor",
			beego.NSInclude(
				&controllers.InformacionSupervisorController{},
			),
		),

		beego.NSNamespace("/objeto_contratar_nucleo_basico",
			beego.NSInclude(
				&controllers.ObjetoContratarNucleoBasicoController{},
			),
		),

		beego.NSNamespace("/parametro_dependencia",
			beego.NSInclude(
				&controllers.ParametroDependenciaController{},
			),
		),

		beego.NSNamespace("/telefono",
			beego.NSInclude(
				&controllers.TelefonoController{},
			),
		),

		beego.NSNamespace("/ordenador_gasto",
			beego.NSInclude(
				&controllers.OrdenadorGastoController{},
			),
		),

		beego.NSNamespace("/parametro_unidad",
			beego.NSInclude(
				&controllers.ParametroUnidadController{},
			),
		),

		beego.NSNamespace("/proveedor_representante_legal",
			beego.NSInclude(
				&controllers.ProveedorRepresentanteLegalController{},
			),
		),

		beego.NSNamespace("/informacion_sociedad_participante",
			beego.NSInclude(
				&controllers.InformacionSociedadParticipanteController{},
			),
		),

		beego.NSNamespace("/estado_solicitud_necesidad",
			beego.NSInclude(
				&controllers.EstadoSolicitudNecesidadController{},
			),
		),

		beego.NSNamespace("/ordenador_gasto_temporal",
			beego.NSInclude(
				&controllers.OrdenadorGastoTemporalController{},
			),
		),

		beego.NSNamespace("/reduccion",
			beego.NSInclude(
				&controllers.ReduccionController{},
			),
		),

		beego.NSNamespace("/contrato_cancelado",
			beego.NSInclude(
				&controllers.ContratoCanceladoController{},
			),
		),

		beego.NSNamespace("/solicitud_disponibilidad",
			beego.NSInclude(
				&controllers.SolicitudDisponibilidadController{},
			),
		),

		beego.NSNamespace("/rubro_solicitud_necesidad",
			beego.NSInclude(
				&controllers.RubroSolicitudNecesidadController{},
			),
		),

		beego.NSNamespace("/contrato_cps",
			beego.NSInclude(
				&controllers.ContratoCpsController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/cesion",
			beego.NSInclude(
				&controllers.CesionController{},
			),
		),

		beego.NSNamespace("/dependencia_dependencia_temporal",
			beego.NSInclude(
				&controllers.DependenciaDependenciaTemporalController{},
			),
		),

		beego.NSNamespace("/fuente_financiacion_rubro_necesidad",
			beego.NSInclude(
				&controllers.FuenteFinanciacionRubroNecesidadController{},
			),
		),

		beego.NSNamespace("/terminacion_anticipada",
			beego.NSInclude(
				&controllers.TerminacionAnticipadaController{},
			),
		),

		beego.NSNamespace("/plantilla_minuta",
			beego.NSInclude(
				&controllers.PlantillaMinutaController{},
			),
		),

		beego.NSNamespace("/parametros",
			beego.NSInclude(
				&controllers.ParametrosController{},
			),
		),

		beego.NSNamespace("/adicion",
			beego.NSInclude(
				&controllers.AdicionController{},
			),
		),

		beego.NSNamespace("/amparo_contrato",
			beego.NSInclude(
				&controllers.AmparoContratoController{},
			),
		),

		beego.NSNamespace("/jefe_dependencia_temporal",
			beego.NSInclude(
				&controllers.JefeDependenciaTemporalController{},
			),
		),

		beego.NSNamespace("/contrato_estado",
			beego.NSInclude(
				&controllers.ContratoEstadoController{},
			),
		),

		beego.NSNamespace("/marco_legal_necesidad",
			beego.NSInclude(
				&controllers.MarcoLegalNecesidadController{},
			),
		),

		beego.NSNamespace("/novedad_contractual",
			beego.NSInclude(
				&controllers.NovedadContractualController{},
			),
		),

		beego.NSNamespace("/tipo_contratacion",
			beego.NSInclude(
				&controllers.TipoContratacionController{},
			),
		),

		beego.NSNamespace("/dependencia_SIC",
			beego.NSInclude(
				&controllers.DependenciaSICController{},
			),
		),

		beego.NSNamespace("/ordenadores",
			beego.NSInclude(
				&controllers.OrdenadoresController{},
			),
		),

		beego.NSNamespace("/convenio",
			beego.NSInclude(
				&controllers.ConvenioController{},
			),
		),

		beego.NSNamespace("/modalidad_seleccion",
			beego.NSInclude(
				&controllers.ModalidadSeleccionController{},
			),
		),

		beego.NSNamespace("/contrato_arrendamiento",
			beego.NSInclude(
				&controllers.ContratoArrendamientoController{},
			),
		),

		beego.NSNamespace("/modificacion_contractual",
			beego.NSInclude(
				&controllers.ModificacionContractualController{},
			),
		),

		beego.NSNamespace("/poliza",
			beego.NSInclude(
				&controllers.PolizaController{},
			),
		),

		beego.NSNamespace("/servicio",
			beego.NSInclude(
				&controllers.ServicioController{},
			),
		),

		beego.NSNamespace("/supervisor_contrato",
			beego.NSInclude(
				&controllers.SupervisorContratoController{},
			),
		),

		beego.NSNamespace("/contrato_disponibilidad",
			beego.NSInclude(
				&controllers.ContratoDisponibilidadController{},
			),
		),

		beego.NSNamespace("/suspension",
			beego.NSInclude(
				&controllers.SuspensionController{},
			),
		),

		beego.NSNamespace("/grupo_tipo_contrato",
			beego.NSInclude(
				&controllers.GrupoTipoContratoController{},
			),
		),

		beego.NSNamespace("/catalogo_temporal",
			beego.NSInclude(
				&controllers.CatalogoTemporalController{},
			),
		),

		beego.NSNamespace("/cargo_supervisor_temporal",
			beego.NSInclude(
				&controllers.CargoSupervisorTemporalController{},
			),
		),

		beego.NSNamespace("/dependencia_temporal",
			beego.NSInclude(
				&controllers.DependenciaTemporalController{},
			),
		),

		beego.NSNamespace("/marco_legal",
			beego.NSInclude(
				&controllers.MarcoLegalController{},
			),
		),

		beego.NSNamespace("/centro_costos",
			beego.NSInclude(
				&controllers.CentroCostosController{},
			),
		),

		beego.NSNamespace("/actividad_economica_necesidad",
			beego.NSInclude(
				&controllers.ActividadEconomicaNecesidadController{},
			),
		),

		beego.NSNamespace("/contrato_liquidado",
			beego.NSInclude(
				&controllers.ContratoLiquidadoController{},
			),
		),

		beego.NSNamespace("/lugar_ejecucion",
			beego.NSInclude(
				&controllers.LugarEjecucionController{},
			),
		),

		beego.NSNamespace("/contrato_suscrito",
			beego.NSInclude(
				&controllers.ContratoSuscritoController{},
			),
		),

		beego.NSNamespace("/solicitud_necesidad",
			beego.NSInclude(
				&controllers.SolicitudNecesidadController{},
			),
		),

		beego.NSNamespace("/servicio_contrato",
			beego.NSInclude(
				&controllers.ServicioContratoController{},
			),
		),

		beego.NSNamespace("/relacion_parametro",
			beego.NSInclude(
				&controllers.RelacionParametroController{},
			),
		),

		beego.NSNamespace("/acta_inicio",
			beego.NSInclude(
				&controllers.ActaInicioController{},
			),
		),

		beego.NSNamespace("/contrato",
			beego.NSInclude(
				&controllers.ContratoController{},
			),
		),

		beego.NSNamespace("/elemento_acta_recibido",
			beego.NSInclude(
				&controllers.ElementoActaRecibidoController{},
			),
		),

		beego.NSNamespace("/registro_presupuestal_disponibilidad",
			beego.NSInclude(
				&controllers.RegistroPresupuestalDisponibilidadController{},
			),
		),

		beego.NSNamespace("/anulacion",
			beego.NSInclude(
				&controllers.AnulacionController{},
			),
		),

		beego.NSNamespace("/especificacion_tecnica",
			beego.NSInclude(
				&controllers.EspecificacionTecnicaController{},
			),
		),

		beego.NSNamespace("/actividad_solicitud_necesidad",
			beego.NSInclude(
				&controllers.ActividadSolicitudNecesidadController{},
			),
		),

		beego.NSNamespace("/supervisor_solicitud_necesidad",
			beego.NSInclude(
				&controllers.SupervisorSolicitudNecesidadController{},
			),
		),

		beego.NSNamespace("/estado_contrato",
			beego.NSInclude(
				&controllers.EstadoContratoController{},
			),
		),

		beego.NSNamespace("/contrato_general",
			beego.NSInclude(
				&controllers.ContratoGeneralController{},
			),
		),

		beego.NSNamespace("/espaciosfisicos_SIC",
			beego.NSInclude(
				&controllers.EspaciosfisicosSICController{},
			),
		),

		beego.NSNamespace("/sedes_SIC",
			beego.NSInclude(
				&controllers.SedesSICController{},
			),
		),

		beego.NSNamespace("/orden",
			beego.NSInclude(
				&controllers.OrdenController{},
			),
		),

		beego.NSNamespace("/requisito_minimo",
			beego.NSInclude(
				&controllers.RequisitoMinimoController{},
			),
		),

		beego.NSNamespace("/amparo_poliza",
			beego.NSInclude(
				&controllers.AmparoPolizaController{},
			),
		),

		beego.NSNamespace("/puc_temporal",
			beego.NSInclude(
				&controllers.PucTemporalController{},
			),
		),

		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),

		beego.NSNamespace("/centro_costos_centro_costos",
			beego.NSInclude(
				&controllers.CentroCostosCentroCostosController{},
			),
		),

		beego.NSNamespace("/cambio_supervisor",
			beego.NSInclude(
				&controllers.CambioSupervisorController{},
			),
		),

		beego.NSNamespace("/actividades_centro_costos",
			beego.NSInclude(
				&controllers.ActividadesCentroCostosController{},
			),
		),

		beego.NSNamespace("/contenido_resolucion",
			beego.NSInclude(
				&controllers.ResolucionCompletaController{},
			),
		),

		beego.NSNamespace("/persona_escalafon",
			beego.NSInclude(
				&controllers.PersonaEscalafonController{},
			),
		),

		beego.NSNamespace("/precontratado",
			beego.NSInclude(
				&controllers.PrecontratadoController{},
			),
		),

		beego.NSNamespace("/proveedor_contrato_persona",
			beego.NSInclude(
				&controllers.ProveedorContratoPersonaController{},
			),
		),

		beego.NSNamespace("/vigencia_contrato",
			beego.NSInclude(
				&controllers.VigenciaContratoController{},
			),
		),

		beego.NSNamespace("/solicitud_rp",
			beego.NSInclude(
				&controllers.SolicitudRpController{},
			),
		),

		beego.NSNamespace("/disponibilidad_apropiacion_solicitud_rp",
			beego.NSInclude(
				&controllers.DisponibilidadApropiacionSolicitudRpController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
