package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stephannykauane/projeto_it/backend/types"
	"github.com/xuri/excelize/v2"
)

func SaveExcel(w http.ResponseWriter, r *http.Request) {
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var values types.AnaliseRequest
	if err := json.Unmarshal(body, &values); err != nil {
		http.Error(w, "Erro ao deserializar JSON", http.StatusBadRequest)
		return
	}

	templateFile := "RecomendaçãoCalagem.xlsx"

	f, err := excelize.OpenFile(templateFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir arquivo: %v", err), http.StatusInternalServerError)
		return
	}


	sheetName := "Recomendação de Calagem"


	if index, err := f.GetSheetIndex(sheetName); err != nil || index == -1 {
		f.NewSheet(sheetName)
	}

	
	data := map[string]interface{}{
		"A10": values.Aluminio,
		"B10": values.Potassio,
		"C10": values.Magnesio,
		"D10": values.Calcio,
		"E10": values.Argila,
		"F10": values.SatD,
		"G10": values.Ctc,
	}

	for cell, value := range data {
		if err := f.SetCellValue(sheetName, cell, value); err != nil {
			fmt.Printf("Erro ao definir valor na célula %s: %v\n", cell, err)
		}
	}


	outputFile := "RecomendaçãoCalagem.xlsx"
	if err := f.SaveAs(outputFile); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao salvar arquivo: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Planilha gerada com sucesso: %s\n", outputFile)


	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Planilha gerada com sucesso: %s", outputFile)))
}
