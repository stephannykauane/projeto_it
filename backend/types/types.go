package types

import (
	"github.com/golang-jwt/jwt/v4"
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

type Analise struct {
	MetodoID int     `json:"id_metodo"`
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

var jwtKey = []byte("801a28164703e91603002e1a64bd675f6788dc1adef7ce1ae407fa155a52ccd09ca3eb969398378b6cb04a7f2ae756fdbf24b0b340e548439ff0eba371430532")

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
}

type Metodo struct {
	MetodoID int  `json:"id_metodo"`
}