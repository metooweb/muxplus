package main

import (
	"github.com/metooweb/muxplus"
	"github.com/gorilla/mux"
	"net/http"
	"context"
	"encoding/json"
)

type OptionHello struct {
	Name string `json:"name" form:"name"`
}

type ResHello struct {
	Msg string
}

func Hello(ctx context.Context, options *OptionHello) (res *ResHello, err error) {

	res = &ResHello{}

	res.Msg = "hello world " + options.Name

	return
}

func ResponseHandler(next muxplus.Handler) muxplus.Handler {

	return muxplus.HandlerFunc(func(w http.ResponseWriter, req *http.Request, fv *muxplus.FuncVal) {

		next.Deal(w, req, fv)

		if resErr := fv.Outs[ 1 ]; !resErr.IsNil() {

			http.Error(w, resErr.String(), http.StatusInternalServerError)

		} else {

			if err := json.NewEncoder(w).Encode(fv.Outs[0].Interface()); err != nil {
				panic(err)
			}

		}

	})

}

func main() {

	handler := muxplus.DefaultArgsParseHandler(
		ResponseHandler(
			muxplus.DefaultServerHandler{},
		),
	)

	router := mux.NewRouter()

	muxplus.HandleFuncPlus(router, "/", Hello, handler)

	http.ListenAndServe(":8081", router)
}
