package api

import (
	"fmt"
	"net/http"

	db "github.com/SanjaySinghRajpoot/GoBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}

	if !server.validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}

	arg := db.TranferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {

	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account currency does not match!!")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
