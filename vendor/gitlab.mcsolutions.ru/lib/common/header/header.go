package header

import (
	"errors"
	"github.com/gorilla/mux"
	"gitlab.mcsolutions.ru/lib/common/consts"
	"gitlab.mcsolutions.ru/lib/reserr"
	"net/http"
	"strings"
)

func SetContentTypeJson(w *http.ResponseWriter) {
	(*w).Header().Set(consts.CONTENT_TYPE, consts.CONTENT_TYPE_JSON)
}

func SetHeader(w *http.ResponseWriter) {
	(*w).Header().Set(consts.CORS_TYPE, "*")
	SetContentTypeJson(w)
}

func SetHeaderAndExit(w *http.ResponseWriter, r *http.Request) bool {
	//(*w).Header().Set(configmap.CONTENT_TYPE, configmap.CONTENT_TYPE_JSON)
	//	(*w).Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")//Origin, Content-Type, X-Auth-Token
	if r.Method == http.MethodOptions {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		(*w).Header().Set("Access-Control-Allow-Credentials", "true")
		(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		(*w).Header().Set("Access-Control-Allow-Headers", "*") //Origin, Content-Type, X-Auth-Token
		(*w).WriteHeader(http.StatusNoContent)
		return true
	}
	(*w).Header().Set(consts.CORS_TYPE, "*")
	SetContentTypeJson(w)
	return false
}

func GetToken(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return "", reserr.BAD_AUTHORIZATION_TOKEN
	}
	strs := strings.Split(authorization, " ")
	if len(strs) < 2 {
		return "", reserr.BAD_AUTHORIZATION_TOKEN
	}
	if strs[0] != "Bearer" {
		return "", reserr.BAD_AUTHORIZATION_TOKEN
	}
	return strs[1], nil
}

func GetVar(r *http.Request, name string) string {
	vars := mux.Vars(r)
	return vars[name]
}

func GetVarRequired(r *http.Request, name string) (string, error) {
	value := GetVar(r, name)
	if value == "" {
		return "", errors.New(name + " not specified")
	}
	return value, nil
}

func Get2Vars(r *http.Request, name1, name2 string) (string, string) {
	vars := mux.Vars(r)
	return vars[name1], vars[name2]
}

func GetVars(r *http.Request, names ...string) *[]string {
	vars := mux.Vars(r)
	result := make([]string, 0)
	for _, name := range names {
		result = append(result, vars[name])
	}
	return &result
}
