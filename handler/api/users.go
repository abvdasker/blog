package api

import (
	"context"
	"encoding/json"
	"fmt"
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
	Create() httprouter.Handle
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
	return a.HandleLogin
}

func (a *users) Create() httprouter.Handle {
	return a.HandleCreate
}

func (a *users) HandleLogin(responseWriter http.ResponseWriter, rawRequest *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	request, err := parseLoginRequest(rawRequest)
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
	token := model.NewToken(user.UUID, user.Username, user.Salt)
	if err := a.tokensDAL.Create(ctx, token); err != nil {
		a.logger.With(zap.Error(err)).Error("failed to write token to database")
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

func (a *users) HandleCreate(responseWriter http.ResponseWriter, rawRequest *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	request, err := parseCreateUserRequest(rawRequest)
	if err != nil {
		respondErr(responseWriter, "failed to parse request")
		return
	}
	user := request.ToUser()
	if err := a.usersDAL.Create(ctx, user); err != nil {
		respondErr(responseWriter, fmt.Sprintf("could not create user with username %s", user.Username))
		return
	}

	return
}

func parseLoginRequest(rawRequest *http.Request) (*model.LoginRequest, error) {
	request := new(model.LoginRequest)
	if err := parseRequest(rawRequest, request); err != nil {
		return nil, err
	}
	return request, nil
}

func parseCreateUserRequest(rawRequest *http.Request) (*model.CreateUserRequest, error) {
	request := new(model.CreateUserRequest)
	if err := parseRequest(rawRequest, request); err != nil {
		return nil, err
	}
	return request, nil
}

func parseRequest(rawRequest *http.Request, request interface{}) error {
	data, err := ioutil.ReadAll(rawRequest.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, request); err != nil {
		return err
	}
	return nil
}
