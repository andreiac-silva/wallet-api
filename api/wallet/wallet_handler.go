package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"go.uber.org/zap"

	"wallet-api/api"
	"wallet-api/domain/commands"
)

const (
	id = "id"
)

type WalletHandler struct {
	commandHandler eh.CommandHandler
}

func NewWalletHandler(commandHandler eh.CommandHandler) WalletHandler {
	return WalletHandler{
		commandHandler: commandHandler,
	}
}

func (h WalletHandler) Routes(router *gin.Engine) {
	router.POST("/wallets", h.create)
	router.PATCH("/wallets/:id/credit", h.credit)
	router.PATCH("/wallets/:id/debit", h.debit)
}

func (h WalletHandler) create(c *gin.Context) {
	request := CreateRequest{}
	if err := c.BindJSON(&request); err != nil {
		zap.S().Errorw("invalid payload", "error", err)
		_ = c.Error(api.ErrInvalidPayload{})
		return
	}

	if err := request.Validate(); err != nil {
		zap.S().Errorw("invalid create request field(s)", "error", err)
		_ = c.Error(err)
		return
	}

	newWalletID := uuid.New()
	cmd := commands.NewCreateWalletCommand(newWalletID, request.DocumentNumber)
	if err := h.commandHandler.HandleCommand(context.Background(), cmd); err != nil {
		zap.S().Errorw("create command could not be performed", "error", err)
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, CreateResponse{ID: newWalletID.String()})
}

func (h WalletHandler) credit(c *gin.Context) {
	idStr := c.Params.ByName(id)
	wid, err := uuid.Parse(idStr)
	if err != nil {
		_ = c.Error(api.ErrInvalidID{})
		return
	}

	request := OperationRequest{}
	if err := c.BindJSON(&request); err != nil {
		zap.S().Errorw("invalid payload", "error", err)
		_ = c.Error(api.ErrInvalidPayload{})
		return
	}

	if err := request.Validate(); err != nil {
		zap.S().Errorw("invalid credit request field(s)", "error", err)
		_ = c.Error(err)
		return
	}

	cmd := commands.NewCreditCommand(wid, request.Amount, request.Description)
	if err := h.commandHandler.HandleCommand(context.Background(), cmd); err != nil {
		zap.S().Errorw("invalid credit command", "error", err)
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, nil)
}

func (h WalletHandler) debit(c *gin.Context) {
	idStr := c.Params.ByName(id)
	wid, err := uuid.Parse(idStr)
	if err != nil {
		_ = c.Error(api.ErrInvalidID{})
		return
	}

	request := OperationRequest{}
	if err := c.BindJSON(&request); err != nil {
		zap.S().Errorw("invalid payload", "error", err)
		_ = c.Error(api.ErrInvalidPayload{})
		return
	}

	if err := request.Validate(); err != nil {
		zap.S().Errorw("invalid debit request field(s)", "error", err)
		_ = c.Error(err)
		return
	}
	cmd := commands.NewDebitCommand(wid, request.Amount, request.Description)
	if err := h.commandHandler.HandleCommand(context.Background(), cmd); err != nil {
		zap.S().Errorw("invalid debit command", "error", err)
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, nil)
}
