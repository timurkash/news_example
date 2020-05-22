package main

import (
	"gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/routers"
	"gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/version"
	"gitlab.mcsolutions.ru/find-psy/common/consts"
	"gitlab.mcsolutions.ru/find-psy/configmap"
	"gitlab.mcsolutions.ru/lib/common/config"
	"gitlab.mcsolutions.ru/lib/common/logger"
	"gitlab.mcsolutions.ru/lib/common/server"
	"time"
)

var (
	LOG_LEVEL  = config.GetEnv("LOG_LEVEL", "info")
	SENTRY_DSN = config.GetEnv("SENTRY_DSN", "")
)

func main() {
	logger := logger.NewLogger(LOG_LEVEL, SENTRY_DSN)
	server := server.HttpServer{
		Started:  time.Now(),
		Subgroup: configmap.NEWS_GITLAB_SUBGROUP,
		Name:     configmap.NEWS,
		Port:     configmap.NEWS_PORT,
		BasePath: configmap.NEWS_BASE_PATH,
		Duration: consts.GRSH_SECONDS * time.Second,
		Routes:   routers.GetRoutes(logger),
		Docs:     "./docs",
		Logger:   logger,
		Version:  version.VERSION,
		Revision: version.REVISION,
		Envs:     &version.Envs,
	}
	server.RunGrSh()
}
