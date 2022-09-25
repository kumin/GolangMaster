package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kumin/GolangMaster/restful/entities"
)

type HandlerFn func(r *http.Request) (interface{}, error)

// handle http error
func HandlerWrapper(hf HandlerFn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj, err := hf(r)
		if err != nil {
			switch err {
			case entities.ParamInvalid:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			case entities.MethodNotAllowErr:
				http.Error(w, err.Error(), http.StatusMethodNotAllowed)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		d, err := json.Marshal(obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(d)
	}
}
