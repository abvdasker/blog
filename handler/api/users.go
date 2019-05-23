package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"

	"github.com/abvdasker/blog/dal"
	"github.com/abvdasker/blog/handler/api/middleware"
	"github.com/abvdasker/blog/lib"
	httplib "github.com/abvdasker/blog/lib/http"
	"github.com/abvdasker/blog/model"
)

type Users interface {
	Login() httprouter.Handle
	Create() httprouter.Handle
}

type users struct {
	usersDAL       dal.Users
	tokensDAL      dal.Tokens
	authMiddleware middleware.Auth
	logger         *zap.SugaredLogger
}

func NewUsers(usersDAL dal.Users, tokensDAL dal.Tokens, authMiddleware middleware.Auth, logger *zap.SugaredLogger) Users {
	return &users{
		usersDAL:       usersDAL,
		tokensDAL:      tokensDAL,
		authMiddleware: authMiddleware,
		logger:         logger,
	}
}

func (a *users) Login() httprouter.Handle {
	return a.HandleLogin
}

func (a *users) Create() httprouter.Handle {
	return a.authMiddleware.Wrap(a.HandleCreate)
}

func (a *users) HandleLogin(responseWriter http.ResponseWriter, rawRequest *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	request, err := parseLoginRequest(rawRequest)
	if err != nil {
		httplib.RespondErr(responseWriter, "failed to parse request")
		return
	}

	user, err := a.usersDAL.ReadByUsername(
		ctx,
		request.Username,
	)
	if err != nil {
		a.logger.With(zap.Error(err)).Error("read user by username error")
		httplib.RespondErr(responseWriter, "error reading user from database")
		return
	}

	if user == nil {
		httplib.RespondUnauthorized(responseWriter, "incorrect username or password")
		return
	}

	hash := lib.HashPassword64(user.Username, user.Salt, request.Password)
	if !lib.SecureStringsEqual(hash, user.PasswordHash) {
		httplib.RespondUnauthorized(responseWriter, "incorrect username or password")
		return
	}

	token := model.NewToken(user.UUID, user.Username, user.Salt)
	if err := a.tokensDAL.Create(ctx, token); err != nil {
		a.logger.With(zap.Error(err)).Error("failed to write token to database")
		httplib.RespondErr(responseWriter, "failed to write token to database")
		return
	}

	data, err := json.Marshal(token)
	if err != nil {
		httplib.RespondErr(responseWriter, "error serializing token data")
		return
	}
	responseWriter.Write(data)
}

func (a *users) HandleCreate(responseWriter http.ResponseWriter, rawRequest *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	request, err := parseCreateUserRequest(rawRequest)
	if err != nil {
		httplib.RespondErr(responseWriter, "failed to parse request")
		return
	}

	user := request.ToUser()
	if err := a.usersDAL.Create(ctx, user); err != nil {
		httplib.RespondErr(responseWriter, fmt.Sprintf("could not create user with username %s", user.Username))
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		httplib.RespondErr(responseWriter, "error serializing user data")
		return
	}
	responseWriter.Write(data)
}

func parseLoginRequest(rawRequest *http.Request) (*model.LoginRequest, error) {
	request := new(model.LoginRequest)
	if err := httplib.ParseRequest(rawRequest, request); err != nil {
		return nil, err
	}
	return request, nil
}

func parseCreateUserRequest(rawRequest *http.Request) (*model.CreateUserRequest, error) {
	request := new(model.CreateUserRequest)
	if err := httplib.ParseRequest(rawRequest, request); err != nil {
		return nil, err
	}
	return request, nil
}
