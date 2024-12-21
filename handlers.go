package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (s *api) handlerCreateService() http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		v := UHC{}

		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			s.logger.Error(err.Error())
			s.clientError(w, err, 400)
		}

		s.logger.Info("UHC struct values", "name", v.Name, "replicaCount", v.ReplicaCount, "CPU limits", v.Resources.Limits.CPU)

		buf := bytes.Buffer{}

		err = s.templateCache.ExecuteTemplate(&buf, "uhc", v)
		if err != nil {
			s.logger.Error(err.Error())
		}
		s.logger.Info("UHC template as string", "uhc", buf.String())
	}

	return http.HandlerFunc(fn)
}
