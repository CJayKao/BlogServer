package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	// "github.com/unrolled/render"
)

// var renderer *render.Render = render.New()

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Login",
		http.HandlerFunc(LoginHandler),
	).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}

type LoginInfo struct {
	Accout   string `json:"accout"`
	Password string `json:"password"`
	Info     string `json:"Info"`
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	// Read request body
	defer req.Body.Close()
	zerolog.TimeFieldFormat = ""
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	loginData := &LoginInfo{}

	if err := json.Unmarshal(b, loginData); err != nil {
		panic(err)
	}

	log.Info().Msg("LoginHandler, " + "Accout" + loginData.Accout + ", Password" + loginData.Password)
	// Response
	loginData.Accout = "333333"
	loginData.Info = "pass OK"
	response, _ := json.Marshal(loginData)
	//w.Header().Add("jjjj", "SSS")
	w.Header().Set("CONTENT-TYPE", "application/json; charset=utf-8")
	w.Write(response)

	// renderer.JSON(w, http.StatusOK, x)
}
