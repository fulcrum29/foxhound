package main

import "net/http"

func (a *api) clientError(w http.ResponseWriter, err error, status int) {

	http.Error(w, err.Error(), status)

}
