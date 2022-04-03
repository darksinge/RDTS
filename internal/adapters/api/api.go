package api

import (
	"orcinator/internal/domain/saga"
	"orcinator/internal/ports"
)

type App struct {
	api ports.ApiPort
	db  ports.DbPort
}

func (app *App) RegisterService(name string) {
	svc := saga.NewService(name)
	app.db.CreateService(svc)
}
