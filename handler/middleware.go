package handler

import (
	"net/http"
	"net/http/httputil"

	"go.uber.org/zap"
)

type Middleware interface {
	Wrap(http.Handler) http.Handler
}

type middleware struct {
	logger *zap.SugaredLogger
}

type wrappedHandler struct {
	base   http.Handler
	logger *zap.SugaredLogger
}

func NewMiddleware(logger *zap.SugaredLogger) Middleware {
	return &middleware{
		logger: logger,
	}
}

func (h *middleware) Wrap(handler http.Handler) http.Handler {
	return &wrappedHandler{
		base:   handler,
		logger: h.logger,
	}
}

func (h *wrappedHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	requestDump, err := httputil.DumpRequest(request, false)
	if err != nil {
		h.logger.With(
			zap.Error(err),
		).Error("error logging HTTP request")
	}
	h.logger.With(
		zap.String("request", string(requestDump)),
	).Info("HTTP request")
	h.base.ServeHTTP(responseWriter, request)
}
