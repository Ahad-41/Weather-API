package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddleares = append(mngr.globalMiddleares, middlewares...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func (mngr *Manager) WrapMux(mux *http.ServeMux) http.Handler {
	var h http.Handler = mux

	// Apply global middlewares in reverse order to ensure the first one added is the outermost
	for i := len(mngr.globalMiddleares) - 1; i >= 0; i-- {
		h = mngr.globalMiddleares[i](h)
	}

	return h
}
