package muxplus

import (
	"net/http"
)

type Handler interface {
	Deal(http.ResponseWriter, *http.Request, *FuncVal)
}

type HandlerFunc func(http.ResponseWriter, *http.Request, *FuncVal)

func (f HandlerFunc) Deal(w http.ResponseWriter, r *http.Request, fv *FuncVal) {
	f(w, r, fv)
}

type DefaultServerHandler struct {
}

func (DefaultServerHandler) Deal(w http.ResponseWriter, req *http.Request, fv *FuncVal) {

	fv.Outs = fv.Func.Func.Call(fv.Args)

}
