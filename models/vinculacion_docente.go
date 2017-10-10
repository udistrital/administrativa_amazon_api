package models

import "time"

type VinculacionDocente struct {
	Id_RENAME            int       `orm:"column(id)"`
	NumeroContrato       string    `orm:"column(numero_contrato);null"`
	Vigencia             int       `orm:"column(vigencia);null"`
	IdPersona            float64   `orm:"column(id_persona)"`
	NumeroHorasSemanales int       `orm:"column(numero_horas_semanales)"`
	NumeroSemanas        int       `orm:"column(numero_semanas)"`
	IdPuntoSalarial      int       `orm:"column(id_punto_salarial);null"`
	IdSalarioMinimo      int       `orm:"column(id_salario_minimo);null"`
	IdResolucion         int       `orm:"column(id_resolucion)"`
	IdDedicacion         int       `orm:"column(id_dedicacion)"`
	IdProyectoCurricular int16     `orm:"column(id_proyecto_curricular)"`
	Estado               bool      `orm:"column(estado)"`
	FechaRegistro        time.Time `orm:"column(fecha_registro);type(date)"`
}
