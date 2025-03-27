package calculator

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/stephannykauane/projeto_it/backend/types"
)



func Calcular(i interface{}) {
	switch v := i.(type) {
	case types.SatBases:
		CalculoSatBases(v)
    case types.Aluminio:
		CalculoAluminio(v)
	default:
		fmt.Println("Tipo invalido")
	}
}

func CalculoSatBases(a types.SatBases) (float64, float64) {

	s := a.Calcio + a.Magnesio + a.Potassio
	V1 := (s / a.Ctc) * 100

	SatAtual := math.Round(V1)


	fmt.Println("valor saturação atual: ", SatAtual)
	fmt.Println("valor s atual: ", s)
	fmt.Println("valor saturação desejada: ", a.SatD)

	NC := ((a.Ctc * (a.SatD - SatAtual)) / 10) * float64(a.Prnt)

	Result := math.Round(NC)
	
	fmt.Println("valor ctc: ", a.Ctc)
	fmt.Println("valor Magnesio: ", a.Magnesio)
	fmt.Println("valor calcio: ", a.Calcio)
	fmt.Println("valor Potassio: ", a.Potassio)
	fmt.Println("valor prnt: ", a.Prnt)

	return Result, SatAtual

}


func CalculoAluminio(b types.Aluminio) float64 {
	var NC float64

	if b.Ctc > 4.0 && b.Argila > 15 && (b.Calcio+b.Magnesio) < 2.0 {

		NC = ((2 * b.Aluminio) + 2 - (b.Calcio + b.Magnesio)) * float64(b.Prnt)
		fmt.Println("esse foi utilizado!")

	} else {
		NC = (2 * b.Aluminio) * float64(b.Prnt)
		fmt.Println("esse AQUI foi utilizado!")
	}
	Result := math.Round(NC)

	fmt.Println("valor ctc: ", b.Ctc)
	fmt.Println("valor argila: ", b.Argila)
	fmt.Println("valor calcio: ", b.Calcio)
	fmt.Println("valor magnesio: ", b.Magnesio)
	fmt.Println("valor aluminio: ", b.Aluminio)
	fmt.Println("valor prnt: ", b.Prnt)
	fmt.Println("nc: ", NC)

	return Result
}



func Calculando(jsonData []byte, MetodoID int) (float64, float64, error) {
	
	var err error
	var Resultado float64
    var SatAtual float64
	switch MetodoID {
	case 1:
		var satbases types.SatBases
		err = json.Unmarshal(jsonData, &satbases)

		if err != nil {

			return 0, 0, fmt.Errorf("erro ao decodificar JSON para SatBases: %v", err)
		}

		Resultado, SatAtual = CalculoSatBases(satbases)

	case 2:
		var aluminio types.Aluminio
		err = json.Unmarshal(jsonData, &aluminio)

		if err != nil {

			return 0, 0, fmt.Errorf("erro ao decodificar JSON para Aluminio: %v", err)
		}

		Resultado = CalculoAluminio(aluminio)

	default:

		return 0, 0, fmt.Errorf("metodoID desconhecido")
	}

	return Resultado, SatAtual, nil

}
