package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"threadule/backend/internal/app"
	"threadule/backend/internal/web"
)

func ctxWrapper(appCtx *app.Context, handler web.Handler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		handler(&web.Context{
			Response: writer,
			Request:  request,
			Params:   params,
			AppCtx:   appCtx,
		})
	}
}

type router struct {
	*httprouter.Router
	appCtx *app.Context
}

func (r *router) GET(path string, handler web.Handler) {
	r.Router.GET(path, ctxWrapper(r.appCtx, handler))
}

func (r *router) POST(path string, handler web.Handler) {
	r.Router.POST(path, ctxWrapper(r.appCtx, handler))
}

func (r *router) OPTIONS(path string, handler web.Handler) {
	r.Router.OPTIONS(path, ctxWrapper(r.appCtx, handler))
}

func (r *router) HEAD(path string, handler web.Handler) {
	r.Router.HEAD(path, ctxWrapper(r.appCtx, handler))
}

func (r *router) DELETE(path string, handler web.Handler) {
	r.Router.DELETE(path, ctxWrapper(r.appCtx, handler))
}
