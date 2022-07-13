package handler

import (
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/service"
	"go.uber.org/fx"
	"net/http"
	"time"
)

var (
	Module = fx.Provide(newHandler)
)

type Ihandler interface {
	CheckAccount() http.HandlerFunc
	TopUp() http.HandlerFunc
	GetBalance() http.HandlerFunc
	GetTransactions() http.HandlerFunc
	TrnConfirmer(ticker *time.Ticker)
}

type handlerParams struct {
	fx.In
	service.Iservice
}

type handler struct {
	service service.Iservice
}

func newHandler(hp handlerParams) Ihandler {
	return &handler{hp.Iservice}
}
