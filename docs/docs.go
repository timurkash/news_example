package docs

import (
	"bytes"
	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
	"gitlab.mcsolutions.ru/find-psy/configmap"
	"gitlab.mcsolutions.ru/lib/common/consts"
	"gitlab.mcsolutions.ru/lib/common/json"
	"gitlab.mcsolutions.ru/lib/common/status"
	"net/http"
	"strconv"
)

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

var (
	SwaggerInfo = swaggerInfo{Schemes: []string{}}
	doc         *string
)

type s struct{}

func (s *s) ReadDoc() string {
	if doc == nil {
		resp, err := http.Get(consts.LOCALHOST + strconv.Itoa(configmap.NEWS_PORT) + consts.SWAGGER_DOCS + "swagger.yaml")
		if err != nil {
			return ""
		}
		bytes, err := json.DecodeYamlEncodeByte(resp.Body)
		if err != nil {
			return ""
		}
		str := string(*bytes)
		doc = &str
	}
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			bytes, err := json.EncodeByte(v)
			if err != nil {
				bytes = status.ErrorStatusOut(err)
			}
			return string(*bytes)
		},
	}).Parse(*doc)
	if err != nil {
		return *doc
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return *doc
	}
	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
