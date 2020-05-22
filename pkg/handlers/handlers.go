package handlers

import (
	"gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/handlers/db"
	"gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/models"
	"gitlab.mcsolutions.ru/find-psy/common/header"
	"gitlab.mcsolutions.ru/find-psy/common/roles"
	"gitlab.mcsolutions.ru/lib/common/hh"
	"gitlab.mcsolutions.ru/lib/common/json"
	"gitlab.mcsolutions.ru/lib/common/logger"
	"gitlab.mcsolutions.ru/lib/common/paging"
	"net/http"
)

type Routes struct {
	Logger *logger.Logger
	News   *db.News
}

func (ro *Routes) GetRoute() hh.Route {
	return hh.Route{
		Pattern:     "/get",
		Methods:     http.MethodGet,
		HandlerFunc: hh.GetObjHandlerFunc(ro.getFunc, ro.Logger),
	}
}

func (ro *Routes) getFunc(r *http.Request) (interface{}, error) {
	lang, err := header.GetLangRequired(r)
	if err != nil {
		return nil, err
	}
	page, err := header.GetPage(r)
	if err != nil {
		return nil, err
	}
	news, paging_, err := ro.News.Get(lang, "all", page)
	if err != nil {
		return nil, err
	}
	return &models.NewsOut{
		News:      news,
		PagingOut: paging.PagingOut{paging_},
	}, nil
}

func (ro *Routes) GetRegionRoute() hh.Route {
	return hh.Route{
		Pattern:     "/get/{region}",
		Methods:     http.MethodGet,
		HandlerFunc: hh.GetObjHandlerFunc(ro.getRegionFunc, ro.Logger),
	}
}

func (ro *Routes) getRegionFunc(r *http.Request) (interface{}, error) {
	lang, err := header.GetLangRequired(r)
	if err != nil {
		return nil, err
	}
	region, err := header.GetPathRegionRequired(r)
	if err != nil {
		return nil, err
	}
	pages, err := header.GetPage(r)
	if err != nil {
		return nil, err
	}
	news, paging_, err := ro.News.Get(lang, region, pages)
	if err != nil {
		return nil, err
	}
	return &models.NewsOut{
		News:      news,
		PagingOut: paging.PagingOut{Paging: paging_},
	}, nil
}

func (ro *Routes) GetFullRoute() hh.Route {
	return hh.Route{
		Pattern:     "/getFull",
		Methods:     http.MethodGet,
		HandlerFunc: hh.GetObjHandlerFunc(ro.getFullFunc, ro.Logger),
	}
}

func (ro *Routes) getFullFunc(r *http.Request) (interface{}, error) {
	if _, err := header.GetAdminLogin(r); err != nil {
		return nil, err
	}
	if err := header.HasRole(r, roles.NEWS_ADMIN_ROLE); err != nil {
		return nil, err
	}
	lang, err := header.GetLangRequired(r)
	if err != nil {
		return nil, err
	}
	pages, err := header.GetPage(r)
	if err != nil {
		return nil, err
	}
	newsFull, paging_, err := ro.News.GetFull(lang, "all", pages)
	if err != nil {
		return nil, err
	}
	return &models.NewsFullOut{
		News:      newsFull,
		PagingOut: paging.PagingOut{Paging: paging_},
	}, nil
}

func (ro *Routes) GetFullRegionRoute() hh.Route {
	return hh.Route{
		Pattern:     "/getFull/{region}",
		Methods:     http.MethodGet,
		HandlerFunc: hh.GetObjHandlerFunc(ro.getFullRegionFunc, ro.Logger),
	}
}

func (ro *Routes) getFullRegionFunc(r *http.Request) (interface{}, error) {
	if _, err := header.GetAdminLogin(r); err != nil {
		return nil, err
	}
	if err := header.HasRole(r, roles.NEWS_ADMIN_ROLE); err != nil {
		return nil, err
	}
	lang, err := header.GetLangRequired(r)
	if err != nil {
		return nil, err
	}
	region, err := header.GetPathRegionRequired(r)
	if err != nil {
		return nil, err
	}
	pages, err := header.GetPage(r)
	if err != nil {
		return nil, err
	}
	newsFull, paging_, err := ro.News.GetFull(lang, region, pages)
	if err != nil {
		return nil, err
	}
	return &models.NewsFullOut{
		News:      newsFull,
		PagingOut: paging.PagingOut{Paging: paging_},
	}, nil
}

func (ro *Routes) PostRoute() hh.Route {
	return hh.Route{
		Pattern:     "/post",
		Methods:     http.MethodPost,
		HandlerFunc: hh.GetObjHandlerFunc(ro.postFunc, ro.Logger),
	}
}

func (ro *Routes) postFunc(r *http.Request) (interface{}, error) {
	lang, err := header.GetLangRequired(r)
	if err != nil {
		return nil, err
	}
	login, err := header.GetAdminLogin(r)
	if err != nil {
		return nil, err
	}
	if header.HasRole(r, roles.NEWS_ADMIN_ROLE); err != nil {
		return nil, err
	}
	oneNews := &models.OneNews{}
	if err = json.Decode(r.Body, oneNews); err != nil {
		return nil, err
	}
	idOut, err := ro.News.Insert(lang, login, oneNews)
	if err != nil {
		return nil, err
	}
	return idOut, nil
}

func (ro *Routes) PutRoute() hh.Route {
	return hh.Route{
		Pattern:     "/put/{id}",
		Methods:     http.MethodPut,
		HandlerFunc: hh.GetObjHandlerFunc(ro.putFunc, ro.Logger),
	}
}

func (ro *Routes) putFunc(r *http.Request) (interface{}, error) {
	id, err := header.GetPathIdRequired(r)
	if err != nil {
		return nil, err
	}
	if header.HasRole(r, roles.NEWS_ADMIN_ROLE); err != nil {
		return nil, err
	}
	login, err := header.GetAdminLogin(r)
	if err != nil {
		return nil, nil
	}
	oneNews := &models.OneNews{}
	if err = json.Decode(r.Body, oneNews); err != nil {
		return nil, err
	}
	idOut, err := ro.News.Update(id, login, oneNews)
	if err != nil {
		return nil, err
	}
	return idOut, nil
}

func (ro *Routes) DeleteRoute() hh.Route {
	return hh.Route{
		Pattern:     "/delete/{id}",
		Methods:     http.MethodDelete,
		HandlerFunc: hh.GetHandlerFunc(ro.deleteFunc, ro.Logger),
	}
}

func (ro *Routes) deleteFunc(r *http.Request) error {
	id, err := header.GetPathIdRequired(r)
	if err != nil {
		return err
	}
	if header.HasRole(r, roles.NEWS_ADMIN_ROLE); err != nil {
		return err
	}
	if _, err = header.GetAdminLogin(r); err != nil {
		return nil
	}
	return ro.News.Delete(id)
}
