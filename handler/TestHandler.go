package handler

import (
	"net/http"
	"model/test"
)

func CreateTest(response http.ResponseWriter, request *http.Request, content string) (bool){

	response.Header().Set("content-type","application/json")

	var Test test

	return true
}
