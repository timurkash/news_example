package routers

import (
	_ "gitlab.mcsolutions.ru/find-psy/back/admins/news/docs"
	"gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/handlers"
	"gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/handlers/db"
	"gitlab.mcsolutions.ru/lib/common/hh"
	"gitlab.mcsolutions.ru/lib/common/logger"
)

func GetRoutes(logger *logger.Logger) *[]hh.Route {
	routes := handlers.Routes{
		Logger: logger,
		News:   &db.News{},
	}
	return &[]hh.Route{
		//user
		routes.GetRoute(),
		routes.GetRegionRoute(),
		//admin
		routes.GetFullRoute(),
		routes.GetFullRegionRoute(),
		routes.PostRoute(),
		routes.PutRoute(),
		routes.DeleteRoute(),
	}
}
