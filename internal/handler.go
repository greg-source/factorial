package internal

import (
	"encoding/json"
	"errors"
	"github.com/greg-source/factorial/pkg"
	"github.com/julienschmidt/httprouter"
	"math/big"
	"net/http"
	"runtime"
)

type request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type response struct {
	A *big.Int `json:"a_factorial"`
	B *big.Int `json:"b_factorial"`
}

func factorialHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request, err := validateRequest(r)
	if err != nil {
		ErrorJson(w, http.StatusBadRequest, err)
		return
	}
	response := response{
		A: pkg.CalculateFactorialConcurrently(request.A, runtime.NumCPU()),
		B: pkg.CalculateFactorialConcurrently(request.B, runtime.NumCPU()),
	}
	WriteJSON(w, http.StatusOK, response)
}

func validateRequest(r *http.Request) (*request, error) {
	decoder := json.NewDecoder(r.Body)
	var t = request{A: -1, B: -1}
	err := decoder.Decode(&t)
	if err != nil {
		return nil, errors.New("incorrect input")
	}
	if t.A < 0 || t.B < 0 {
		return nil, errors.New("incorrect input")
	}
	return &t, err
}

//func middleware(n httprouter.Handle) httprouter.Handle {
//	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)
//		// call registered factorialHandler
//		n(w, r, ps)
//	}
//}
