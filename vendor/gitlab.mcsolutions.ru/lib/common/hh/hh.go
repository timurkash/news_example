package hh

import (
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.mcsolutions.ru/lib/common/header"
	"gitlab.mcsolutions.ru/lib/common/json"
	"gitlab.mcsolutions.ru/lib/common/logger"
	"gitlab.mcsolutions.ru/lib/common/status"
	"io"
	"net/http"
	"strings"
)

type (
	Route struct {
		Pattern     string
		Methods     string
		HandlerFunc http.HandlerFunc
	}
	//MSRoute struct {
	//	PathPrefix string
	//	MSUrl      string
	//	MSName     string
	//}
)

//var (
//	msRoutes []MSRoute
//)
//
const (
	OPTIONS = "," + http.MethodOptions
)

func NewRouter(routes *[]Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range *routes {
		handler := route.HandlerFunc
		methods := strings.Split(route.Methods+OPTIONS, ",")
		//if route.MSUrl != "" {
		//	router.
		//		Methods(methods...).
		//		PathPrefix(route.Pattern).
		//		Handler(handler)
		//	msRoutes = append(msRoutes,
		//		MSRoute{
		//			route.Pattern,
		//			route.MSUrl,
		//			route.MSName})
		//} else {
		router.
			Methods(methods...).
			Path(route.Pattern).
			Handler(handler)
		//}
	}
	return router
}

//func GetRoute(pattern, method string, routes *[]Route) *Route {
//	pattern_ := pattern
//	if pattern_ == "" {
//		pattern_ = "/"
//	}
//	for _, route := range *routes {
//		if route.Pattern == pattern_ &&
//			others.Contains(strings.Split(route.Methods+OPTIONS, ","), method) {
//			return &route
//		}
//	}
//	return nil
//}
//
//func GetMSRoute(pattern string) *MSRoute {
//	for _, msRoute := range msRoutes {
//		if msRoute.PathPrefix == pattern {
//			return &msRoute
//		}
//	}
//	return nil
//}
//
func GetHandlerFunc(f func(*http.Request) error, logger *logger.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		header.SetHeader(&w)
		if err := f(r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			status.WriteLoggerStatus(&w, err, logger)
		} else {
			status.WriteStatus(&w, nil)
		}
	}
}

func GetCodeHandlerFunc(f func(*http.Request) (int, error), logger *logger.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		header.SetHeader(&w)
		if code, err := f(r); err != nil {
			w.WriteHeader(code)
			status.WriteLoggerStatus(&w, err, logger)
		} else {
			status.WriteStatus(&w, nil)
		}
	}
}

//func GetBodyHandlerFunc(f func(*http.Request) (*[]byte, error), logger *logger.Logger) func(http.ResponseWriter, *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		header.SetHeader(&w)
//		if body, err := f(r); err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			status.WriteLoggerStatus(&w, err, logger)
//		} else {
//			if body != nil {
//				w.Write(*body)
//			}
//		}
//	}
//}
//
func GetObjHandlerFunc(f func(*http.Request) (interface{}, error), logger *logger.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if obj, err := f(r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			status.WriteLoggerStatus(&w, err, logger)
		} else if obj == nil {
			w.WriteHeader(http.StatusNoContent)
		} else if bytes, ok := obj.([]byte); ok {
			w.Write(bytes)
		} else if bytes, ok := obj.(*[]byte); ok {
			w.Write(*bytes)
		} else {
			header.SetHeader(&w)
			if err := json.Encode(obj, io.Writer(w)); err != nil {
				status.StatusJsonEncodeError(&w, err, logger)
			}
		}
	}
}

func GetCodeObjHandlerFunc(f func(*http.Request) (int, interface{}, error), logger *logger.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if code, obj, err := f(r); err != nil {
			w.WriteHeader(code)
			status.WriteLoggerStatus(&w, err, logger)
		} else if obj == nil {
			w.WriteHeader(http.StatusNoContent)
		} else if bytes, ok := obj.([]byte); ok {
			w.WriteHeader(code)
			w.Write(bytes)
		} else if bytes, ok := obj.(*[]byte); ok {
			w.WriteHeader(code)
			w.Write(*bytes)
		} else {
			header.SetHeader(&w)
			w.WriteHeader(code)
			if err := json.Encode(obj, io.Writer(w)); err != nil {
				status.StatusJsonEncodeError(&w, err, logger)
			}
		}
	}
}

func GetCodeObjFileHandlerFunc(f func(*http.Request) (int, interface{}, string, error), logger *logger.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if code, obj, filename, err := f(r); err != nil {
			w.WriteHeader(code)
			status.WriteLoggerStatus(&w, err, logger)
		} else if obj == nil {
			w.WriteHeader(http.StatusNoContent)
		} else if bytes, ok := obj.([]byte); ok {
			w.WriteHeader(code)
			w.Write(bytes)
		} else if bytes, ok := obj.(*[]byte); ok {
			w.WriteHeader(code)
			w.Write(*bytes)
		} else {
			w.Header().Set("Content-Disposition",
				fmt.Sprintf("attachment; filename=%s", filename))
			http.ServeFile(w, r, filename)
		}
	}
}
