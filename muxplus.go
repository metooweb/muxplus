package muxplus

import (
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"github.com/go-playground/form"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
)

type FuncVal struct {
	Func *Func
	Args []reflect.Value
	Outs []reflect.Value
}

func HandleFuncPlus(router *mux.Router, path string, function interface{}, handler Handler) *mux.Route {

	//解析func
	f := FuncParse(function)

	return router.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {

		var (
			err error
			fv  = FuncVal{Func: f}
		)

		if err = req.ParseForm(); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		handler.Deal(w, req, &fv)

	})

}

func DefaultArgsParseHandler(handler Handler) Handler {

	return HandlerFunc(func(w http.ResponseWriter, req *http.Request, fv *FuncVal) {
		var (
			err error
			ctx = req.Context()
		)

		fv.Args = append(fv.Args, reflect.ValueOf(ctx))

		if len(fv.Func.In) > 1 {

			opts := reflect.New(fv.Func.In[1].Elem()).Interface()

			contentType := req.Header.Get("Content-Type")

			if strings.Contains(contentType, "application/json") {

				var bytes []byte

				defer req.Body.Close()

				if bytes, err = ioutil.ReadAll(req.Body); err == nil {

					fmt.Println("=====================", string(bytes))

					err = json.Unmarshal(bytes, opts)
				}
			} else {

				err = form.NewDecoder().Decode(opts, req.Form)

			}

			if err != nil {

				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			fv.Args = append(fv.Args, reflect.ValueOf(opts))

		}

		handler.Deal(w, req, fv)

	})
}
