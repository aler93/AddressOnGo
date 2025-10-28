package main

import (
	"encoding/json"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	query := r.URL.Query()
	cepParam := query.Get("cep")

	if cepParam == "" {
		http.Error(w, "Parâmetro 'cep' não encontrado", http.StatusBadRequest)
		return
	}

	if get(cepParam) {
		renderJson(JsonResp{w, r, result, 200})
	} else {
		doReq(cepParam, w, r)
	}
}

type JsonResp struct {
	w      http.ResponseWriter
	r      *http.Request
	Data   any `json:"data"`
	Status int `json:"status"`
}

func renderJson(res JsonResp) {
	res.w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(res.Data)
	if err != nil {
		res.w.WriteHeader(http.StatusInternalServerError)
		res.w.Write([]byte("{\"status\": \"error\"}"))
		return
	}

	if res.Status <= 0 {
		res.Status = 200
	}

	res.w.WriteHeader(res.Status)
	res.w.Write(response)
}
