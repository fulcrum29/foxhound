package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"
)

func (s *api) handlerCreateService() http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		v := UHC{}

		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			s.logger.Error(err.Error())
		}
		s.logger.Info("UHC struct values", "name", v.Name, "replicaCount", v.ReplicaCount, "CPU limits", v.Resources.Limits.CPU)
		tmpl, err := template.ParseFiles("uhc.tmpl")
		if err != nil {
			s.logger.Error(err.Error())
		}

		buf := bytes.Buffer{}
		err = tmpl.ExecuteTemplate(&buf, "uhc", v)
		if err != nil {
			s.logger.Error(err.Error())
		}
		s.logger.Info("UHC template as string", "uhc", buf.String())
	}

	return http.HandlerFunc(fn)
}
