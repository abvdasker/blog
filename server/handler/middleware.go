package handler

import (
	"net/http"
	"net/http/httputil"

	"go.uber.org/zap"
)

type MiddlewareHandler interface {
	http.Handler
}

type middlewareHandler struct {
	logger *zap.SugaredLogger
}

func NewMiddlewareHandler(logger *zap.SugaredLogger) MiddlewareHandler {
	return &middlewareHandler{
		logger: logger,
	}
}

func (h *middlewareHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	requestDump, err := httputil.DumpRequest(request, false)
	if err != nil {
		h.logger.With(
			zap.Error(err),
		).Error("error logging HTTP request")
	}
	h.logger.With(
		zap.String("request", string(requestDump)),
	).Info("HTTP request")
	http.DefaultServeMux.ServeHTTP(responseWriter, request)
}
