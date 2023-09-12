package account

import (
	"net/http"

	"eth-account-creator-api/core/domains/account/models"
	"eth-account-creator-api/internal/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	Service ServiceI
	log     *log.Logger
}

func NewHandler(service ServiceI, logger *log.Logger) *Handler {
	return &Handler{
		Service: service,
		log:     logger,
	}
}

// CreateNewAccount godoc
// @Summary Creates a new Ethereum Account
// @Description Creates a new Ethereum Account returning new address data.
// @Tags Account
// @Produce json
// @Success 200 {object} []models.Account
// @Failure 404 {object} models.ErrorResponse
// @Router /account/create [get].
func (h *Handler) CreateNewAccount(c *gin.Context) {
	account, err := h.Service.CreateAddress()
	if err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	h.log.Zap.Info("CreateNewAccount - Account created")
	c.JSON(http.StatusOK, account)
}

// GetAccountFromPubKey godoc
// @Summary Gets a Ethereum Address by giving it's public key
// @Description Gets a Ethereum Address by passing it's public key through URL.
// @Tags Account
// @Produce json
// @Success 200 {object} []models.Address
// @Failure 404 {object} models.ErrorResponse
// @Router /account/{publicKey} [get].
func (h *Handler) GetAccountFromPubKey(c *gin.Context) {
	publicKey := c.Params.ByName("publicKey")
	address, err := h.Service.FetchAddressFromPubKey(publicKey)
	if err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	h.log.Zap.Info("GetAccountFromPubKey - Account fetched successfuly")
	c.JSON(http.StatusOK, address)
}

// MakeTransaction godoc
// @Summary Make a transaction
// @Description Transfer ETH between accounts
// @Tags Account
// @Accept json
// @Produce json
// @Param Transaction body models.Transaction true "Transaction Model"
// @Success 200 {object} string
// @Failure 404 {object} models.ErrorResponse
// @Router /account/transaction [post].
func (h *Handler) MakeTransaction(c *gin.Context) {
	var transactionPayload models.Transaction
	if err := c.ShouldBindJSON(&transactionPayload); err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	if err := h.Service.SendTransaction(transactionPayload); err != nil {
		h.log.Zap.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, "Transaction Sended")
}
