package models

import (
	"encoding/json"
	"net/http"
)

var ErrorResponse = func(err interface{}) string {
	jsonString, _ := json.Marshal(err)
	return string(jsonString)
}

type Response struct {
	response interface{} `json:"response"`
}

func (r *Response) Set(i interface{}) {
	r.response = i
}

func (r Response) Get() interface{} {
	return r.response
}

func (r Response) marshal() ([]byte, error) {
	marshal, err := json.Marshal(r.Get())
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (r Response) ToJson(w http.ResponseWriter) {

	reply, err := r.marshal()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Write(reply)
}
