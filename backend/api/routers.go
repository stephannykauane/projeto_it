package api

import (
	"net/http"

)

func Routes() {
	http.HandleFunc("/signup", PostSignUp)
	http.HandleFunc("/login", GetLogin)
	
}
