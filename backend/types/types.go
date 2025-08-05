package types

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Aluminio struct {
	Ctc      float64 `json:"ctc"`
	Argila   int     `json:"argila"`
	Calcio   float64 `json:"calcio"`
	Magnesio float64 `json:"magnesio"`
	Aluminio float64 `json:"aluminio"`
	Prnt     int     `json:"prnt"`
}

type SatBases struct {
	SatD float64 `json:"sat_desejada"`
	SatA float64 `json:"sat_atual"`
	Prnt int     `json:"prnt"`
	Ctc  float64 `json:"ctc"`
}

type SatCalcio struct {
	Prnt       int     `json:"prnt"`
	Ctc        float64 `json:"ctc"`
	TeorCa     float64 `json:"teor_ca"`
	CaO        float64 `json:"caO"`
	CaDesejada float64 `json:"ca_desejada"`
	TeorMg     float64 `json:"teor_mg"`
	MgO        float64 `json:"mgO"`
}

type SatMagnesio struct {
	Prnt       int     `json:"prnt"`
	Ctc        float64 `json:"ctc"`
	TeorMg     float64 `json:"teor_mg"`
	MgO        float64 `json:"mgO"`
	MgDesejada float64 `json:"mg_desejada"`
	TeorCa     float64 `json:"teor_ca"`
	CaO        float64 `json:"caO"`
}

type Area struct {
	Consultor   string `json:"consultor"`
	Propriedade string `json:"propriedade"`
	Area        string `json:"area"`
}

type Usuario struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
	Nome  string `json:"nome"`
}

type Request struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type AnaliseRequest struct {
	UsuarioID   int     `json:"id_usuario"`
	Consultor   string  `json:"consultor"`
	Propriedade string  `json:"propriedade"`
	Area        string  `json:"area"`
	TeorCa      float64 `json:"teor_ca"`
	SatA        float64 `json:"sat_atual"`
	CaO         float64 `json:"caO"`
	CaDesejada  float64 `json:"ca_desejada"`
	TeorMg      float64 `json:"teor_mg"`
	MgO         float64 `json:"mgO"`
	MgDesejada  float64 `json:"mg_desejada"`
	Magnesio    float64 `json:"magnesio"`
	Aluminio    float64 `json:"aluminio"`
	Calcio      float64 `json:"calcio"`
	Argila      int     `json:"argila"`
	SatD        float64 `json:"sat_desejada"`
	Prnt        int     `json:"prnt"`
	Ctc         float64 `json:"ctc"`
	Resultado   float64 `json:"resultado"`
}

type Metodo struct {
	MetodoID int `json:"id_metodo"`
}

type CalculoDetalhes struct {
	Resultado   float64   `json:"resultado"`
	DataCalculo time.Time `json:"data_calculo"`
	MetodoID    int       `json:"id_metodo"`
	SatExtra    float64   `json:"sat_extra"`
	RelacaoCaMg float64   `json:"relacao_ca_mg"`
	Magnesio    float64   `json:"magnesio"`
	Aluminio    float64   `json:"aluminio"`
	Calcio      float64   `json:"calcio"`
	SatDesejada float64   `json:"sat_desejada"`
	Prnt        float64   `json:"prnt"`
	Ctc         float64   `json:"ctc"`
	Argila      float64   `json:"argila"`
	SatAtual    float64   `json:"sat_atual"`
	TeorCa      float64   `json:"teor_ca"`
	CaO         float64   `json:"caO"`
	MgO         float64   `json:"mgO"`
	TeorMg      float64   `json:"teor_mg"`
	MgDesejada  float64   `json:"mg_desejada"`
	CaDesejada  float64   `json:"ca_desejada"`
	Consultor   string    `json:"consultor"`
	Propriedade string    `json:"propriedade"`
	Area        string    `json:"area"`
}
