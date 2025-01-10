package calculator

import (
	"encoding/json"
	"fmt"
)

func Calculando (jsonData []byte, MetodoID int)(float64, error){
	var resultado float64
    var err error

	switch MetodoID {
	case 1: 
	   var satbases SatBases
	   err = json.Unmarshal(jsonData, &satbases)

	   if err != nil {
		
		return 0,fmt.Errorf("erro ao decodificar JSON para SatBases: %v", err)		
	   }

	   resultado = gerarCalculo(satbases) 
	
	case 2: 
	   var aluminio Aluminio
	   err = json.Unmarshal(jsonData, &aluminio)

	   if err != nil{
		
		return 0, fmt.Errorf("erro ao decodificar JSON para Aluminio: %v", err)
	   }

	   resultado = gerarCalculo(aluminio)
	
	default:
	   
	   return 0,fmt.Errorf("metodoID desconhecido") 
	}

	return resultado, nil




}