package api

import (
	"github.com/go-chi/chi/v5"
)

type Router interface {
	Routes(router *chi.Mux)
}
