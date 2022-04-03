package api

import (
	"orcinator/internal/core/saga"
	"orcinator/internal/ports"
)

type App struct {
	api ports.ApiPort
	db  ports.IDbPort
}

func (app *App) RegisterService(name string) {
	svc := saga.NewService(name)
	app.db.CreateService(svc)
}
