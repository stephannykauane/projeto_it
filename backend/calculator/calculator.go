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
	SatExtra sql.NullFloat64
	RelacaoCaMg sql.NullFloat64

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


	return result

}

func CalcOptional(
	tonCorretivo float64,  
	teorElemento float64,  
	fatorCorrecao float64, 
	fatorCorretivo float64,
	Elemento float64,
	ctc float64, 
	prnt float64,          
) sql.NullFloat64 {

	
	CorretivoReagido := (tonCorretivo * prnt) / 100

	
	QtdElementoCorretivo := CorretivoReagido * (Elemento/100)

	ElementoConvertidoMg := (QtdElementoCorretivo * fatorCorretivo)

	ElementoConvertidoMg = ElementoConvertidoMg * 1000


	ConvertidoFinal := (ElementoConvertidoMg / fatorCorrecao) + teorElemento

	ResultadoSaturacao := (ConvertidoFinal / ctc) * 100

	ResultadoSaturacao = math.Round(ResultadoSaturacao * 10) / 10


	var result sql.NullFloat64
	if math.IsNaN(ResultadoSaturacao) {
		result.Valid = false
	} else {
		result.Float64 = ResultadoSaturacao
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
	
		var Magnesio types.SatMagnesio
		err = json.Unmarshal(jsonData, &Magnesio)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para Magnesio: %v", err)
		}

		resultado.Result = CalculoSatMagnesio(Magnesio)

		if Magnesio.CaO > 0 && Magnesio.Ctc > 0 {
			resultado.SatExtra = CalcOptional(
				resultado.Result.Float64,
				Magnesio.TeorCa,
				400.78,
				0.7149,
				Magnesio.CaO,
				Magnesio.Ctc,
				float64(Magnesio.Prnt),
			)
		}

		if Magnesio.TeorCa > 0 {
			relacaoCaMg := (resultado.SatExtra.Float64/10) / ((Magnesio.MgDesejada/100) * Magnesio.Ctc)
			relacaoCaMg = math.Round(relacaoCaMg * 100) / 100
			resultado.RelacaoCaMg = sql.NullFloat64{
				Float64: relacaoCaMg,
				Valid: true,
			}
	
		}
	case 3:

	
		var Calcio types.SatCalcio


		err = json.Unmarshal(jsonData, &Calcio)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para Calcio: %v", err)
		}

		resultado.Result = CalculoSatCalcio(Calcio)

		if Calcio.MgO > 0 && Calcio.Ctc > 0 {
			resultado.SatExtra = CalcOptional(
				resultado.Result.Float64,
				Calcio.TeorMg,
				243.05,
				0.603,
				Calcio.MgO,
				Calcio.Ctc,
				float64(Calcio.Prnt),
			)

			
		}

		if Calcio.TeorMg > 0 {
			fmt.Println("teor ca", Calcio.CaO)
			relacaoCaMg := ((Calcio.CaDesejada / 100) * Calcio.Ctc) / (resultado.SatExtra.Float64/10)
			relacaoCaMg = math.Round(relacaoCaMg * 100)/100
			resultado.RelacaoCaMg = sql.NullFloat64{
				Float64: relacaoCaMg,
				Valid: true,
			}

		}

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
