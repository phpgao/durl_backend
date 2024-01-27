// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/phpgao/durl_backend/internal/biz"
	"github.com/phpgao/durl_backend/internal/conf"
	"github.com/phpgao/durl_backend/internal/data"
	"github.com/phpgao/durl_backend/internal/server"
	"github.com/phpgao/durl_backend/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, app *conf.App, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	leafRepo := data.NewLeafRepo(dataData, logger, app)
	leafUseCase := biz.NewLeafUseCase(leafRepo, logger, app)
	shortUrlRepo := data.NewShortUrlRepo(dataData, logger)
	shortUrlUseCase := biz.NewShortUrlUseCase(shortUrlRepo, leafRepo, logger)
	urlShortenerService := service.NewUrlShortenerService(app, leafUseCase, shortUrlUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, urlShortenerService, logger)
	kratosApp := newApp(logger, httpServer)
	return kratosApp, func() {
		cleanup()
	}, nil
}