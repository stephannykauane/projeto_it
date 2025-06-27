package calculator

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/stephannykauane/projeto_it/backend/types"
	"math"
)

type ResultadoCalculo struct {
	Result   sql.NullFloat64
	SatAtual sql.NullFloat64
}

func CalculoSatBases(a types.SatBases) (sql.NullFloat64, sql.NullFloat64) {

	NC := ((a.SatD - a.SatA) * a.Ctc) / float64(a.Prnt)

	ResultValor := math.Round(NC*100)/100


	var Result, SatAtual sql.NullFloat64

	if math.IsNaN(a.SatA) {
		SatAtual.Valid = false
	} else {
		SatAtual.Float64 = a.SatA
		SatAtual.Valid = true
	}

	if math.IsNaN(ResultValor) {
		Result.Valid = false
	} else {
		Result.Float64 = ResultValor
		Result.Valid = true
	}



	return Result, SatAtual

}

func CalculoAluminio(b types.Aluminio) sql.NullFloat64 {
	var NC float64

	if b.Ctc > 4.0 && b.Argila > 15 && (b.Calcio+b.Magnesio) < 2.0 {

		NC = ((2 * b.Aluminio) + 2 - (b.Calcio + b.Magnesio)) * float64(b.Prnt)

	} else {
		NC = (2 * b.Aluminio) * float64(b.Prnt)

	}
	ResultValor := math.Round(NC*100) / 100

	var Result sql.NullFloat64

	if math.IsNaN(ResultValor) {
		Result.Valid = false
	} else {
		Result.Float64 = ResultValor
		Result.Valid = true
	}
   
	return Result
}

func CalculoSatGeneric(
	ctc float64,
	teor float64,
	desejado float64,
	conversao float64,
	kgBase float64,
	oxido float64,
	prnt float64,
) sql.NullFloat64 {

	cmolNecessario := (ctc * (desejado / 100)) - teor
	kgElemento := kgBase * cmolNecessario
	kgCorretivo := kgElemento * conversao
	xDoKgCorretivo := RegraDeTres(oxido, 100, kgCorretivo)
	resultado := math.Round(RegraDeTres(prnt, xDoKgCorretivo, 100)*100) / 100

	var result sql.NullFloat64

	if math.IsNaN(resultado) {
		result.Valid = false
	} else {
		result.Float64 = resultado
		result.Valid = true
	}

	return result
}

func CalculoSatCalcio(c types.SatCalcio) sql.NullFloat64 {
	return CalculoSatGeneric(
		c.Ctc,
		c.TeorCa,
		c.CaDesejada,
		1.4,
		400.78,
		c.CaO,
		float64(c.Prnt),
	)
}

func CalculoSatMagnesio(c types.SatMagnesio) sql.NullFloat64 {
	return CalculoSatGeneric(
		c.Ctc,
		c.TeorMg,
		c.MgDesejada,
		1.6667,
		243.05,
		c.MgO,
		float64(c.Prnt),
	)
}

func RegraDeTres(a float64, b float64, c float64) float64 {
	return (b * c) / a
}

func Calculando(jsonData []byte, MetodoID int) (ResultadoCalculo, error) {

	var err error
	var resultado ResultadoCalculo

	switch MetodoID {
	case 1:
		var satbases types.SatBases
		err = json.Unmarshal(jsonData, &satbases)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para SatBases: %v", err)
		}

		resultado.Result, resultado.SatAtual = CalculoSatBases(satbases)

	case 2:
		var aluminio types.Aluminio
		err = json.Unmarshal(jsonData, &aluminio)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para Aluminio: %v", err)
		}

		resultado.Result = CalculoAluminio(aluminio)
	case 3:
		var Calcio types.SatCalcio
		err = json.Unmarshal(jsonData, &Calcio)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para Calcio: %v", err)
		}

		resultado.Result = CalculoSatCalcio(Calcio)

	case 4:
		var Magnesio types.SatMagnesio
		err = json.Unmarshal(jsonData, &Magnesio)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para Magnesio: %v", err)
		}

		resultado.Result = CalculoSatMagnesio(Magnesio)

	default:

		return resultado, fmt.Errorf("metodoID desconhecido")
	}

	return resultado, nil

}
