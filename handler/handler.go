package handler

import "finalproj/service"

type HttpServer struct {
	app service.ServiceInterface
}

func NewHttpServer(app service.ServiceInterface) HttpServer {
	return HttpServer{app: app}
}
