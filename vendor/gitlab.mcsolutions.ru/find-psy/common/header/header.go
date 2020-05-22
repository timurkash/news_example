package header

import (
	"errors"
	"gitlab.mcsolutions.ru/lib/common/header"
	"gitlab.mcsolutions.ru/lib/reserr"
	"net/http"
	"strconv"
	"strings"
)

const (
	LOGIN       = "Login"
	ADMIN_LOGIN = "AdminLogin"
	ROLE        = "Role"
	PAGE        = "page"
	LANG        = "lang"
	ID          = "id"
	KEY         = "key"
	REGION      = "region"
)

func GetLoginRequired(r *http.Request) (string, error) {
	login := r.Header.Get(LOGIN)
	if login == "" {
		return "", reserr.LOGIN_NOT_DEFINED
	}
	return login, nil
}

func GetPathLoginRequired(r *http.Request) (string, error) {
	return header.GetVarRequired(r, LOGIN)
}

func GetPathKeyRequired(r *http.Request) (string, error) {
	return header.GetVarRequired(r, KEY)
}

func GetPathRegionRequired(r *http.Request) (string, error) {
	return header.GetVarRequired(r, REGION)
}

func GetLangRequired(r *http.Request) (string, error) {
	lang := r.Header.Get(LANG)
	if lang == "" {
		return "", getErrorNotSpecified(LANG)
	}
	return lang, nil
}

func GetPathIdRequired(r *http.Request) (int64, error) {
	id_, err := header.GetVarRequired(r, ID)
	if err != nil {
		return 0, nil
	}
	id, err := strconv.Atoi(id_)
	if err != nil {
		return 0, nil
	}
	return int64(id), nil
}

func GetAdminLogin(r *http.Request) (string, error) {
	login := r.Header.Get(ADMIN_LOGIN)
	if login == "" {
		return "", reserr.ADMIN_LOGIN_NOT_DEFINED
	}
	return login, nil
}

func HasRole(r *http.Request, roles ...string) error {
	role_ := r.Header.Get(ROLE)
	if role_ == "" {
		return errors.New("role not defined")
	}
	str := ""
	for _, role := range roles {
		if role_ == role {
			str = str + "1"
		} else {
			str = str + "0"
		}
	}
	if strings.Index(str, "1") == -1 {
		return errors.New("has no role")
	}
	return nil
}

func GetPage(r *http.Request) (*[]int64, error) {
	page_ := r.Header.Get(PAGE)
	if page_ == "" {
		return &[]int64{0, 10}, nil
	}
	page := []int64{}
	strs := strings.Split(page_, ",")
	if len(strs) >= 1 {
		from_ := strs[0]
		var from int
		var err error
		if from_ == "" {
			from = 0
		} else {
			from, err = strconv.Atoi(from_)
			if err != nil {
				return nil, err
			}
		}
		page = append(page, int64(from))
		if len(strs) == 1 {
			page = append(page, 0)
		} else {
			count_ := strs[1]
			var count int
			var err error
			if count_ == "" {
				count = 0
			} else {
				count, err = strconv.Atoi(count_)
				if err != nil {
					return nil, err
				}
			}
			page = append(page, int64(count))
		}
	}
	return &page, nil
}

func getErrorNotSpecified(param string) error {
	return errors.New(param + " not specified")
}
