package calculator

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"

	"github.com/stephannykauane/projeto_it/backend/types"
)

type ResultadoCalculo struct {
	Result   sql.NullFloat64
	SatAtual sql.NullFloat64
}

func CalculoSatBases(a types.SatBases) (sql.NullFloat64, sql.NullFloat64) {

	NC := ((a.SatD - a.SatA) * a.Ctc) / float64(a.Prnt)

	ResultValor := math.Round(NC*1000) / 1000

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

    fmt.Println("valores recebidos: ", b.Aluminio)
	fmt.Println("valores recebidos: ", b.Calcio)
	fmt.Println("valores recebidos: ", b.Magnesio)
	fatorCorrecao := 100 / b.Prnt


	nc1 := (2 * b.Aluminio) * float64(fatorCorrecao)
	nc2 := 2 - (b.Calcio+b.Magnesio)*float64(fatorCorrecao)

	maiorValor := nc1
	if nc2 > nc1 {
		maiorValor = nc2
	}

	resultValor := math.Round(maiorValor*1000) / 1000 

	var result sql.NullFloat64
	if math.IsNaN(resultValor) {
		result.Valid = false
	} else {
		result.Float64 = resultValor
		result.Valid = true
	}

	return result
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

	resultadoKg := RegraDeTres(prnt, xDoKgCorretivo, 100)
	resultadoTon := resultadoKg / 1000

	resultado := math.Round(resultadoTon*1000) / 1000

	var result sql.NullFloat64
	if math.IsNaN(resultado) {
		result.Valid = false
	} else {
		result.Float64 = resultado
		result.Valid = true
	}
	fmt.Println("cmolNecessario:", cmolNecessario)
	fmt.Println("kgElemento:", kgElemento)
	fmt.Println("kgCorretivo:", kgCorretivo)
	fmt.Println("xDoKgCorretivo:", xDoKgCorretivo)
	fmt.Println("resultadoTon (final em toneladas):", resultadoTon)

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
	
		var Magnesio types.SatMagnesio
		err = json.Unmarshal(jsonData, &Magnesio)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para Magnesio: %v", err)
		}

		resultado.Result = CalculoSatMagnesio(Magnesio)
	case 3:

	
		var Calcio types.SatCalcio
		err = json.Unmarshal(jsonData, &Calcio)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para Calcio: %v", err)
		}

		resultado.Result = CalculoSatCalcio(Calcio)

	case 4:
		var aluminio types.Aluminio
		err = json.Unmarshal(jsonData, &aluminio)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para Aluminio: %v", err)
		}

		resultado.Result = CalculoAluminio(aluminio)

	default:

		return resultado, fmt.Errorf("metodoID desconhecido")
	}
	fmt.Println(">>> Resultado final enviado:", resultado.Result.Float64)

	return resultado, nil

}
