package calculator

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"

	"github.com/stephannykauane/projeto_it/backend/types"
)

type ResultadoCalculo struct {
	Result sql.NullFloat64
	SatAtual sql.NullFloat64
}


func CalculoSatBases(a types.SatBases) (sql.NullFloat64, sql.NullFloat64) {

	s := a.Calcio + a.Magnesio + a.Potassio
	V1 := (s / a.Ctc) * 100

	SatAtualValor := math.Round(V1)

	NC := ((a.Ctc * (a.SatD - SatAtualValor)) / 10) * float64(a.Prnt)

	ResultValor := math.Round(NC)


	var Result, SatAtual sql.NullFloat64

	if math.IsNaN(SatAtualValor) {
		SatAtual.Valid = false
	} else {
		SatAtual.Float64 = SatAtualValor
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

func CalculoAluminio(b types.Aluminio) (sql.NullFloat64) {
	var NC float64

	if b.Ctc > 4.0 && b.Argila > 15 && (b.Calcio+b.Magnesio) < 2.0 {

		NC = ((2 * b.Aluminio) + 2 - (b.Calcio + b.Magnesio)) * float64(b.Prnt)
		

	} else {
		NC = (2 * b.Aluminio) * float64(b.Prnt)
		
	}
	ResultValor := math.Round(NC)
    
   
	fmt.Println("valor ctc: ", b.Ctc)
	fmt.Println("valor argila: ", b.Argila)
	fmt.Println("valor calcio: ", b.Calcio)
	fmt.Println("valor magnesio: ", b.Magnesio)
	fmt.Println("valor aluminio: ", b.Aluminio)
	fmt.Println("valor prnt: ", b.Prnt)
	fmt.Println("nc: ", NC)
	var Result sql.NullFloat64

    if math.IsNaN(ResultValor) {
        Result.Valid = false
    } else {
        Result.Float64 = ResultValor
        Result.Valid = true
    }
   

	fmt.Println("resultado aluminio:", Result)

	return Result
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

		resultado.Result, resultado.SatAtual =  CalculoSatBases(satbases) 

	case 2:
		var aluminio types.Aluminio
		err = json.Unmarshal(jsonData, &aluminio)

		if err != nil {

			return resultado, fmt.Errorf("erro ao decodificar JSON para Aluminio: %v", err)
		}

		resultado.Result = CalculoAluminio(aluminio)

	default:

		return resultado, fmt.Errorf("metodoID desconhecido")
	}

	return resultado, nil

}
