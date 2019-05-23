package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	
	"github.com/abvdasker/blog/dal"
	httplib "github.com/abvdasker/blog/lib/http"
)

const (
	authTokenType = "Bearer"
)

type Auth interface {
	Wrap(httprouter.Handle) httprouter.Handle
}

type auth struct {
	tokensDAL dal.Tokens
	logger *zap.SugaredLogger
}

func NewAuth(tokensDAL dal.Tokens, logger *zap.SugaredLogger) Auth {
	return &auth{
		tokensDAL: tokensDAL,
		logger: logger,
	}
}

func (a *auth) Wrap(handler httprouter.Handle) httprouter.Handle {
	return func(responseWriter http.ResponseWriter, rawRequest *http.Request, params httprouter.Params) {
		ctx := context.Background()

		tokens := rawRequest.Header["Authorization"]
		if len(tokens) < 1 {
			httplib.RespondBadRequest(responseWriter, "missing authorization token")
			return
		}

		tokenHeader := tokens[0]
		tokenStrSlice := strings.Split(tokenHeader, " ")
		if len(tokenStrSlice) < 2 {
			httplib.RespondBadRequest(responseWriter, "invalid token value format")
			return
		}
		
		if tokenStrSlice[0] != authTokenType {
			httplib.RespondUnauthorized(responseWriter, "unauthorized")
			return
		}

		tokenStr := tokenStrSlice[1]
		token, err := a.tokensDAL.ReadByToken(ctx, tokenStr)
		if err != nil {
			a.logger.With(zap.Error(err)).Error("error reading token from db")
			httplib.RespondErr(responseWriter, "internal server error")
			return
		}

		if token == nil {
			httplib.RespondUnauthorized(responseWriter, "unauthorized")
			return
		}

		if token.Token != tokenStr {
			httplib.RespondErr(responseWriter, "internal server error")
			return
		}

		now := time.Now()
		if token.ExpiresAt.Before(now) {
			httplib.RespondUnauthorized(responseWriter, "session has expired")
			return
		}

		handler(responseWriter, rawRequest, params)
	}
}
