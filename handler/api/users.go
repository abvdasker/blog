package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"

	"github.com/abvdasker/blog/dal"
	"github.com/abvdasker/blog/lib"
	"github.com/abvdasker/blog/model"
)

type Users interface {
	Login() httprouter.Handle
}

type users struct {
	usersDAL  dal.Users
	tokensDAL dal.Tokens
	logger    *zap.SugaredLogger
}

func NewUsers(usersDAL dal.Users, tokensDAL dal.Tokens, logger *zap.SugaredLogger) Users {
	return &users{
		usersDAL:  usersDAL,
		tokensDAL: tokensDAL,
		logger:    logger,
	}
}

func (a *users) Login() httprouter.Handle {
	return a.Handle
}

func (a *users) Handle(responseWriter http.ResponseWriter, rawRequest *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	request, err := parseRequest(rawRequest)
	user, err := a.usersDAL.ReadByUsername(
		ctx,
		request.Username,
	)
	if err != nil {
		a.logger.With(zap.Error(err)).Error("read user by username error")
		respondErr(responseWriter, "error reading user from database")
		return
	}

	if user == nil {
		respondUnauthorized(responseWriter, "incorrect username or password")
		return
	}

	hash := lib.HashPassword64(user.Username, user.Salt, request.Password)
	if !lib.SecureStringsEqual(hash, user.PasswordHash) {
		respondUnauthorized(responseWriter, "incorrect username or password")
		return
	}
	token := model.NewToken(user.ID, user.Username, user.Salt)
	if err := a.tokensDAL.Create(ctx, token); err != nil {
		respondErr(responseWriter, "failed to write token to database")
		return
	}

	data, err := json.Marshal(token)
	if err != nil {
		respondErr(responseWriter, "error serializing token data")
		return
	}
	responseWriter.Write(data)
}

func parseRequest(rawRequest *http.Request) (*model.LoginRequest, error) {
	data, err := ioutil.ReadAll(rawRequest.Body)
	if err != nil {
		return nil, err
	}
	request := new(model.LoginRequest)
	if err = json.Unmarshal(data, request); err != nil {
		return nil, err
	}
	return request, nil
}
