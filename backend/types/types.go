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
	SatD     float64 `json:"sat_desejada"`
	Potassio float64 `json:"potassio"`
	Magnesio float64 `json:"magnesio"`
	Calcio   float64 `json:"Calcio"`
	Prnt     int     `json:"prnt"`
	Ctc      float64 `json:"ctc"`
}


type Usuario struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
	Nome string `json:"nome"`
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
	Potassio    float64 `json:"potassio"`
	Magnesio    float64 `json:"magnesio"`
	Aluminio    float64 `json:"aluminio"`
	Calcio      float64 `json:"calcio"`
    Argila      float64 `json:"argila"`
	SatD        float64 `json:"sat_desejada"`
	Prnt        int     `json:"prnt"`
	Ctc         float64 `json:"ctc"`
	Resultado   float64 `json:"resultado"`
}

type Metodo struct {
	MetodoID int  `json:"id_metodo"`
}

type CalculoDetalhes struct {
	Resultado    float64   `json:"resultado"`
	DataCalculo  time.Time `json:"data_calculo"`
	MetodoID     int       `json:"id_metodo"`
	Potassio     float64   `json:"potassio"`
	Magnesio     float64   `json:"magnesio"`
	Aluminio     float64   `json:"aluminio"`
	Calcio       float64   `json:"calcio"`
	SatDesejada  float64   `json:"sat_desejada"`
	Prnt         float64   `json:"prnt"`
	Ctc          float64   `json:"ctc"`
	Argila       float64   `json:"argila"`
}



