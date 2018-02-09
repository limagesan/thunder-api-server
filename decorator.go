package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func IDShouldBeInt(h httprouter.Handle, name string, idNames []string) httprouter.Handle {
	return CommonHeaders(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		for i := 0; i < len(idNames); i++ {
			idParam := ps.ByName(idNames[i])
			_, err := strconv.Atoi(idParam)
			if err != nil {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(500)
				if err := json.NewEncoder(w).Encode(err); err != nil {
					return
				}
				return
			}
		}

		h(w, r, ps)
	}, name)
}

func CommonHeaders(h httprouter.Handle, name string) httprouter.Handle {
	return Logging(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// 開発用
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		h(w, r, ps)
	}, name)
}
