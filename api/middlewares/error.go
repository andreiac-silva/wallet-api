package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/looplab/eventhorizon"

	"wallet-api/api"
	"wallet-api/domain"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		c.JSON(mapApplicationErrToStatusCode(err), err.Error())
	}
}

func mapApplicationErrToStatusCode(err error) int {
	switch err.(type) {
	case api.ErrInvalidAttribute, api.ErrInvalidID, api.ErrInvalidPayload:
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
