package calculator

type Aluminio struct {
	Ctc      float64 `json:"ctc"`
	Argila   int     `json:"argila"`
	Calcio   float64 `json:"calcio"`
	Magnesio float64 `json:"magnesio"`
	Aluminio float64 `json:"aluminio"`
	Prnt     int     `json:"prnt"`
}

type SatBases struct {
	SatD     float64 `json:"sat_desejada"`
	Potassio float64 `json:"potassio"`
	Magnesio float64 `json:"magnesio"`
	Calcio   float64 `json:"Calcio"`
	Prnt     int     `json:"prnt"`
	Ctc      float64 `json:"ctc"`
}

type Analise struct {
	MetodoID int     `json:"id_metodo"`
}
