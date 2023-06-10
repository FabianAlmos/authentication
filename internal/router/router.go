package router

import (
	"authJWT/internal/middleware"
	"net/http"
)

type route struct {
	URL     string
	handler http.Handler
	m       middleware.Func
}

type Router struct {
	Routes []route
}

func NewRouter() *Router {
	return new(Router)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.Routes {
		if req.URL.Path == route.URL {
			if route.m != nil {
				route.m(route.handler).ServeHTTP(w, req)
			} else {
				route.handler.ServeHTTP(w, req)
			}
		}
	}
}

func (r *Router) AddRoute(URL string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	r.Routes = append(r.Routes, route{
		URL:     URL,
		handler: http.HandlerFunc(f),
		m:       m,
	})
}
