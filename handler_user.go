package main;

import "net/http"

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}

	decoder:=  
}
