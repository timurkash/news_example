package version

import "gitlab.mcsolutions.ru/lib/common/env"

var Envs []env.Env

func init() {
	Envs = append(Envs, env.Env{
		Name:    "POSTGRES_ARGS",
		Default: "postgresql://doadmin@db-postgresql-ams3-60546-do-user-6494983-0.db.ondigitalocean.com:25060/defaultdb?sslmode=require",
		File:    "pkg/handlers/db/pg.go",
		Url:     "http://localhost:3009/pkg/handlers/db/pg.go",
		Type:    "string",
	})
	Envs = append(Envs, env.Env{
		Name:    "LOG_LEVEL",
		Default: "info",
		File:    "pkg/handlers/handlers_test.go",
		Url:     "http://localhost:3009/pkg/handlers/handlers_test.go",
		Type:    "string",
	})
	Envs = append(Envs, env.Env{
		Name:    "SENTRY_DSN",
		Default: "",
		File:    "pkg/handlers/handlers_test.go",
		Url:     "http://localhost:3009/pkg/handlers/handlers_test.go",
		Type:    "string",
	})
}
