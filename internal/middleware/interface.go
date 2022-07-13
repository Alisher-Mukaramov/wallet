package middleware

import (
	"errors"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/service"
	"go.uber.org/fx"
	"net/http"
)

var (
	invalidHash = errors.New("Некорректный хеш")
	Module      = fx.Provide(newMiddleware)
)

const (
	XuserId = "X-UserId"
	Xdigest = "X-Digest"
)

type Imiddleware interface {
	UserVerify(next http.Handler) http.Handler
	HashVerify(next http.Handler) http.Handler
}

type middlewareParams struct {
	fx.In
	service.Iservice
}

type middleware struct {
	service service.Iservice
}

func newMiddleware(mw middlewareParams) Imiddleware {
	return &middleware{mw.Iservice}
}
