package api

import (
	"time"

	db "github.com/SanjaySinghRajpoot/GoBank/db/sqlc"
	"github.com/SanjaySinghRajpoot/GoBank/token"
	"github.com/SanjaySinghRajpoot/GoBank/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	SessionID            uuid.UUID    `json:"session_id"`
	AccessToken          string       `json:"access_token"`
	AccessTokenExpireAt  time.Time    `json:"access_token_expire_at"`
	RefreshToken         string       `json:"refresh_token"`
	RefreshTokenExpireAt time.Time    `json:"refresh_token_expire_at"`
	User                 userResponse `json:"user"`
}

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken         string    `json:"access_token"`
	AccessTokenExpireAt time.Time `json:"access_token_expire_at"`
}

type createAccountRequest struct {
	// Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}
