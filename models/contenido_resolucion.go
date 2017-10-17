package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Paragrafo struct {
	Id     int
	Numero int
	Texto  string
}

type Articulo struct {
	Id         int
	Numero     int
	Texto      string
	Paragrafos []Paragrafo
}

type ResolucionCompleta struct {
	Vinculacion   ResolucionVinculacionDocente
	Consideracion string
	Preambulo     string
	Vigencia      int
	Numero        string
	Id            int
	Articulos     []Articulo
}

func GetOneResolucionCompleta(idResolucion string) (resolucion ResolucionCompleta) {
	o := orm.NewOrm()
	var temp []Resolucion
	_, err := o.Raw("SELECT * FROM administrativa.resolucion WHERE administrativa.resolucion.id_resolucion=" + idResolucion + ";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	resolucionCompleta := ResolucionCompleta{Id: temp[0].Id, Consideracion: temp[0].ConsideracionResolucion, Preambulo: temp[0].PreambuloResolucion, Vigencia: temp[0].Vigencia, Numero: temp[0].NumeroResolucion}

	var arts []ComponenteResolucion
	_, err2 := o.Raw("SELECT * FROM administrativa.componente_resolucion WHERE resolucion_id=" + idResolucion + " AND tipo_componente like 'Articulo' ORDER BY numero asc;").QueryRows(&arts)
	if err2 == nil {
		fmt.Println("Consulta exitosa")
	}

	var articulos []Articulo

	for _, art := range arts {
		articulo := Articulo{Id: art.Id, Numero: art.Numero, Texto: art.Texto}

		var pars []ComponenteResolucion
		_, err3 := o.Raw("SELECT * FROM administrativa.componente_resolucion WHERE resolucion_id=" + idResolucion + " AND tipo_componente like 'Paragrafo' AND componente_padre=" + strconv.Itoa(articulo.Id) + " ORDER BY numero asc;").QueryRows(&pars)
		if err3 == nil {
			fmt.Println("Consulta exitosa")
		}

		var paragrafos []Paragrafo

		for _, par := range pars {
			paragrafo := Paragrafo{Id: par.Id, Numero: par.Numero, Texto: par.Texto}
			paragrafos = append(paragrafos, paragrafo)
		}

		articulo.Paragrafos = paragrafos

		articulos = append(articulos, articulo)
	}
	resolucionCompleta.Articulos = articulos
	return resolucionCompleta
}

func UpdateResolucionCompletaById(m *ResolucionCompleta) (err error) {
	o := orm.NewOrm()
	v := Resolucion{Id: m.Id}
	if err = o.Read(&v); err == nil {
		v.NumeroResolucion = m.Numero
		_, err = o.Update(&v)
	} else {
		return
	}
	idResolucionStr := strconv.Itoa(m.Id)
	r := m.Vinculacion
	fmt.Println(r.Id)
	a := ResolucionVinculacionDocente{Id: r.Id}
	if err = o.Read(&a); err == nil {
		_, err = o.Update(&r)
	} else {
		return
	}
	if err = o.Read(&v); err == nil {
		v.ConsideracionResolucion = m.Consideracion
		v.PreambuloResolucion = m.Preambulo
		v.NumeroResolucion = m.Numero
		fmt.Println(v)
		if err := UpdateResolucionById(&v); err != nil {
		}

		resolucionCompleta := GetOneResolucionCompleta(idResolucionStr)

		for _, articulo := range resolucionCompleta.Articulos {
			if articulo.Paragrafos != nil {
				for _, paragrafo := range articulo.Paragrafos {
					if err := DeleteComponenteResolucion(paragrafo.Id); err != nil {
					}
				}
			}
			if err := DeleteComponenteResolucion(articulo.Id); err != nil {
			}
		}

		for indexArticulo, articulo := range m.Articulos {
			componenteArticulo := ComponenteResolucion{ResolucionId: &Resolucion{Id: m.Id}, Texto: articulo.Texto, Numero: indexArticulo + 1, TipoComponente: "Articulo"}
			if _, err := AddComponenteResolucion(&componenteArticulo); err == nil {
				if articulo.Paragrafos != nil {
					for indexParagrafo, paragrafo := range articulo.Paragrafos {
						componenteParagrafo := ComponenteResolucion{ResolucionId: &Resolucion{Id: m.Id}, Texto: paragrafo.Texto, Numero: indexParagrafo + 1, TipoComponente: "Paragrafo", ComponentePadre: &ComponenteResolucion{Id: componenteArticulo.Id}}
						if _, err := AddComponenteResolucion(&componenteParagrafo); err == nil {

						}
					}
				}
			}
		}
	}
	return
}

func GetTemplateResolucion() (resolucion ResolucionCompleta) {
	resolucion = ResolucionCompleta{Consideracion: "Que mediante Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció un nuevo régimen salarial y prestacional de los docentes de las Universidades estatales u oficiales del Orden Nacional, Departamental, Municipal y Distrital.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes en las modalidades de hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional, en virtud del artículo 13 del Acuerdo 011 de noviembre 15 de 2002. Termino fijo por periodos académicos.\n\nQue la presente rige y se aplica únicamente a los docentes de Vinculación Especial Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional en Pregrado en lo pertinente únicamente al valor del punto.\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución (Resoluciones 318 de septiembre de 2006, 317 de septiembre 8 de 2006, Ley 30 de  1992 y Acuerdo 003 de 1997, articulo 49 y Ley 4 de 1992).\n\nQue el artículo 128 de la Carta Política consigna que nadie podrá desempeñar simultáneamente más de  un empleo público ni recibir más de una asignación que provenga del tesoro público salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4 de 1992 determino como excepción los honorarios percibidos por concepto hora cátedra.\n\nQue mediante Acuerdo 002 de enero 31 de 2003 se modifica y reglamenta el Acuerdo 001 de enero 17 de  2003.\n\nQue para efectos de pago de Salarios y prestaciones sociales, el periodo académico corresponde de conformidad con la Resolución No 071 de Julio 12 de 2016, por medio de la cual se modifica y se expide el calendario académico 2016 comprendido entre el 16 de agosto 2016 al 17 de diciembre del  2016 equivalentes  a 4,5 meses y el mes comprenderá (30) días.\n\nQue de conformidad con el artículo 2° del Decreto 215 de 12 de febrero de 2016 establece “A partir del 1° de enero del 2016, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en doce mil ciento veinte pesos ($12.120) moneda corriente”.\n\nQue de acuerdo a lo establecido en el artículo 5° del Decreto 215 del 12 de febrero de 2016, los Rectores Universitarios expedirán los correspondientes actos administrativos.\n\nQue mediante Resolución 001 del 15 de Febrero del 2012 emitida por la Vicerrectoría Académica establece el proceso de selección y vinculación de Docentes en la modalidad de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional.\n\nQue mediante Resolucion279 del 2016 emitida por la Rectoría acoge y aplica el Decreto 215 de 12 de febrero de 2016, artículo 2°, en lo pertinente única y expresamente al valor del punto para los docentes de vinculación Especial Hora Cátedra, Hora Cátedra Honorarios, Medio Tiempo Ocasional y Tiempo Completo  Ocasional en pregrado.\n\nQue para efectos presupuestales de la presente resolución se hará  a cargo de la Disponibilidad Presupuestal No 548 del 25 de Enero de 2016.", Preambulo: "El decano de la Facultad INGENIERIA  de la Universidad Distrital Francisco José de Caldas en uso de sus facultades legales y estatuarias y"}
	var articulos []Articulo
	articulo := Articulo{Texto: "Vincular para el Tercer Periodo Académico del año 2016, como docentes de Tiempo completo Ocasional (Vinculación Especial) en el escalafón y dedicación establecidos en la tabla, a los siguientes docentes:"}
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "El pago de los servicios prestados por los profesores de Vinculación Especial Según su escalafón, se cancelaran previa certificación de las horas efectivamente dictadas, expedida por el Decano y/o Directora o quien haga las veces de Director."}
	paragrafo := Paragrafo{Texto: "El valor del punto en pesos para el reconocimiento y pago de los docentes de Hora Cátedra, será el que fije el Gobierno Nacional, mediante Decreto cada año y que la Universidad acoja mediante acto administrativo para los docentes de Vinculación Especial."}
	articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "El docente deberá cumplir con las obligaciones inherentes a la naturaleza del servicio, contempladas en la Ley, en los Reglamentos de la Universidad y en los Planes de Trabajo entregados por el Profesor y aprobados por el Decano y/o Directora."}
	paragrafo = Paragrafo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad mediante acto administrativo hará la liquidación con corte a la fecha del cumplido expedido por el Decano y se cancelaran las prestaciones sociales en la última liquidación del periodo académico que efectué la División de Recursos Humanos."}
	articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "El gasto que ocasione la presente resolución se hará con cargo al presupuesto de la actual vigencia, previa certificación de disponibilidad presupuestal."}
	paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias de la secretaria de Hacienda Distrital."}
	articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "Comunicar de la presente resolución al Docente, quien manifestara bajo la gravedad de juramento que no se encuentra incurso en inhabilidades e incompatibilidades de ley para aceptar dicha vinculación especial, que no tienen con cruces de horarios ni ostentan otra vinculación de carácter público, diferentes a hora cátedra en entidades de educación oficiales, siempre y cuando los honorarios sumados no correspondan a más de ocho(8) horas diarias de trabajo, que no ostente vinculación de tiempo Completo en dos entidades ni Medio Tiempo Ocasional únicamente a excepción de hora cátedra."}
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "En caso de declaratoria de suspensión de actividades académicas, por parte de los máximos organismos de gobierno de la Universidad y autorización previa del Ministerio de Trabajo, cesara para el docente la vinculación especial la obligación de prestar el servicio y para la Universidad la de pagar los salarios, pero persistirá para esta última, la de continuar efectuando los respectivos aportes a salud y pensión en el porcentaje que le corresponda."}
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "La presente Resolución rige a partir de la fecha de su expedición y surte efectos a partir del 16 de agosto del 2016, para el tercer periodo académico del año 2016."}
	articulos = append(articulos, articulo)
	resolucion.Articulos = articulos //articulos//articulos
	//resolucion.Vinculacion=nil
	return
}
