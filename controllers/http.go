package controllers

import "mygram/services"

type HttpServer struct {
	app services.ServiceInterface
}

func NewHttpServer(app services.ServiceInterface) HttpServer {
	return HttpServer{app: app}
}
