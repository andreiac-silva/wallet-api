package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"go.uber.org/zap"

	"wallet-api/api"
	"wallet-api/domain/commands"
)

const (
	errInvalidID      = "Invalid id"
	errInvalidPayload = "Payload could not be parsed"
)

type WalletHandler struct {
	commandHandler eh.CommandHandler
}

func NewWalletHandler(commandHandler eh.CommandHandler) WalletHandler {
	return WalletHandler{
		commandHandler: commandHandler,
	}
}

func (h WalletHandler) Routes(router *chi.Mux) {
	router.Method(http.MethodPost, "/wallets", api.Handler(h.create))
	router.Method(http.MethodPatch, "/wallets/{id}/credit", api.Handler(h.credit))
	router.Method(http.MethodPatch, "/wallets/{id}/debit", api.Handler(h.debit))
}

func (h WalletHandler) create(w http.ResponseWriter, r *http.Request) error {
	request := CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return api.ErrInvalidPayload{Message: errInvalidPayload}
	}
	if err := request.Validate(); err != nil {
		zap.S().Error("invalid create request field(s)", "error", err)
		return err
	}

	newWalletID := uuid.New()
	cmd := commands.NewCreateWalletCommand(newWalletID, request.DocumentNumber)
	if err := h.commandHandler.HandleCommand(context.Background(), cmd); err != nil {
		zap.S().Error("create command could not be performed", "error", err)
		return err
	}

	api.Render.JSON(w, 201, CreateResponse{ID: newWalletID.String()})
	return nil
}

func (h WalletHandler) credit(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return api.ErrInvalidID{Message: errInvalidID}
	}

	request := OperationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return api.ErrInvalidPayload{Message: errInvalidPayload}
	}
	if err := request.Validate(); err != nil {
		zap.S().Error("invalid credit request field(s)", "error", err)
		return api.ErrInvalidAttribute{Message: err.Error()}
	}

	cmd := commands.NewCreditCommand(id, request.Amount, request.Description)
	if err := h.commandHandler.HandleCommand(context.Background(), cmd); err != nil {
		zap.S().Error("invalid credit command", "error", err)
		return err
	}

	api.Render.JSON(w, 201, EmptyResponse)
	return nil
}

func (h WalletHandler) debit(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return api.ErrInvalidID{Message: errInvalidID}
	}

	request := OperationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return api.ErrInvalidPayload{Message: errInvalidPayload}
	}
	if err := request.Validate(); err != nil {
		zap.S().Error("invalid debit request field(s)", "error", err)
		return api.ErrInvalidAttribute{Message: err.Error()}
	}

	cmd := commands.NewDebitCommand(id, request.Amount, request.Description)
	if err := h.commandHandler.HandleCommand(context.Background(), cmd); err != nil {
		zap.S().Error("invalid debit command", "error", err)
		return err
	}

	api.Render.JSON(w, 201, EmptyResponse)
	return nil
}
