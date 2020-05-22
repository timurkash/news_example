package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"gitlab.mcsolutions.ru/lib/common/consts"
	"gitlab.mcsolutions.ru/lib/common/env"
	"gitlab.mcsolutions.ru/lib/common/hh"
	"gitlab.mcsolutions.ru/lib/common/logger"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
	"time"
)

type Static struct {
	Route   string
	FileDir string
}

type HttpServer struct {
	Server   *http.Server
	Started  time.Time
	Subgroup string
	Name     string
	Port     int
	portStr  string
	Duration time.Duration
	BasePath string
	Routes   *[]hh.Route
	Docs     string
	Logger   *logger.Logger
	Static   *Static
	Version  string
	Revision string
	Envs     *[]env.Env
}

const PKG = "/pkg"

func (s *HttpServer) initWithRouter() *mux.Router {
	s.portStr = strconv.Itoa(s.Port)
	log.Println("Server started on :" + s.portStr)
	log.Println("Start time:", s.Started)
	return s.getRouter()
}

func (s *HttpServer) Run() {
	handler := s.initWithRouter()
	s.Server = &http.Server{
		Addr:    ":" + s.portStr,
		Handler: handler,
	}
	s.Logger.Zap.Fatal(s.Server.ListenAndServe().Error())
}

func (s *HttpServer) RunGrSh() {
	handler := s.initWithRouter()
	s.Server = &http.Server{
		Addr:         ":" + s.portStr,
		WriteTimeout: s.Duration,
		ReadTimeout:  s.Duration,
		IdleTimeout:  s.Duration * 2,
		Handler:      handler,
	}
	s.listenAndServe()
}

func (s *HttpServer) getRouter() *mux.Router {
	routes := []hh.Route{}
	for _, route := range *s.Routes {
		if s.BasePath != "" {
			route.Pattern = s.BasePath + route.Pattern
		}
		//route.MSName = name
		routes = append(routes, route)
	}
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()
	devopsRouters := hh.DevOpsRouters{
		Started:     &s.Started,
		SubgroupUrl: s.Subgroup,
		Name:        s.Name,
		Version:     s.Version,
		Revision:    s.Revision,
		Envs:        s.Envs,
		Logger:      s.Logger,
		IsReady:     isReady,
	}
	router := hh.NewRouter(devopsRouters.AddTo(&routes))
	if s.Docs != "" {
		router.PathPrefix(consts.SWAGGER_ROUTE).Handler(httpSwagger.Handler(httpSwagger.URL(consts.SWAGGER_ROUTE + "doc.json")))
		router.Handle(consts.SWAGGER_DOCS+"{rest}", http.StripPrefix(consts.SWAGGER_DOCS, http.FileServer(http.Dir("./docs"))))
		log.Println("Swagger started on", "http://localhost:"+s.portStr+consts.SWAGGER_ROUTE, "swagger.yaml on ", "http://localhost:"+s.portStr+consts.SWAGGER_DOCS+"swagger.yaml")
	}
	if s.Static != nil {
		handler := http.StripPrefix(s.Static.Route, http.FileServer(http.Dir(s.Static.FileDir)))
		router.Handle(s.Static.Route+"/{rest}", handler)
		router.Handle(s.Static.Route+"/{rest}/{rest}", handler)
	}
	{
		handler := http.StripPrefix(PKG, http.FileServer(http.Dir("."+PKG)))
		router.Handle(PKG+"/{rest}", handler)
		router.Handle(PKG+"/{rest}/{rest}", handler)
		router.Handle(PKG+"/{rest}/{rest}/{rest}", handler)
		router.Handle(PKG+"/{rest}/{rest}/{rest}/{rest}", handler)
	}
	router.Handle("/metrics", promhttp.Handler())
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	return router
}

func (s *HttpServer) listenAndServe() {
	go func() {
		if err := s.Server.ListenAndServe(); err != nil {
			log.Fatal(err)
			s.Logger.Zap.Fatal(err.Error())
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), s.Duration)
	defer cancel()
	err := s.Server.Shutdown(ctx)
	if err != nil {
		s.Logger.Zap.Info(err.Error())
	}
	log.Println(s.Server.Shutdown(ctx))
	log.Println("shutting down")
	os.Exit(0)
}
