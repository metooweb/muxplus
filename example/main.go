package main

import (
	"github.com/metooweb/muxplus"
	"github.com/gorilla/mux"
	"net/http"
	"context"
)

func main() {

	router := mux.NewRouter()
	handler := muxplus.DefaultArgsParseHandler(
		Test(
			muxplus.DefaultServerHandler{},
		),
	)

	muxplus.HandleFuncPlus(router, "/", func(ctx context.Context) (int, error) {

		return 1, nil

	}, handler)

	http.ListenAndServe(":81", router)

}

func Test(next muxplus.Handler) muxplus.Handler {

	return muxplus.HandlerFunc(func(w http.ResponseWriter, req *http.Request, fv *muxplus.FuncVal) {

		next.Deal(w, req, fv)

		if !fv.Outs[0].IsNil() {

		}

	})

}
