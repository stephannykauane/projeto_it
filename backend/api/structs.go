package api

type Usuario struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
	Nome string `json:"nome"`
}

type Request struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}