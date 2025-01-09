package api

import (
	
	"github.com/golang-jwt/jwt/v4"
)

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

type metodoID struct {
	MetodoID int `json: "metodo_id"`

}

