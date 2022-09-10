package api

import (
	"errors"
	"net/http"
	"wallet-api/domain"

	"github.com/looplab/eventhorizon"
	renderPkg "github.com/unrolled/render"
)

var Render *renderPkg.Render

func init() {
	Render = renderPkg.New()
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		statusCode := mapApplicationErrToStatusCode(err)
		Render.JSON(w, statusCode, err.Error())
	}
}

func mapApplicationErrToStatusCode(err error) int {
	switch err.(type) {
	case ErrInvalidAttribute, ErrInvalidID, ErrInvalidPayload:
		return http.StatusBadRequest
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrIncompatibleProjectionVersion, domain.ErrInsufficientAmount:
		return http.StatusConflict
	case domain.ErrUnprocessable:
		return http.StatusUnprocessableEntity
	case *eventhorizon.AggregateError:
		if errors.As(err.(*eventhorizon.AggregateError).Err, &domain.ErrNotFound{}) {
			return http.StatusNotFound
		}
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
