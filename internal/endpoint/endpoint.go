package endpoint

import (
	"context"
	"github.com/Alisher-Mukaramov/wallet/internal/handler"
	"github.com/Alisher-Mukaramov/wallet/internal/middleware"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net/http"
	"time"
)

var Module = fx.Options(
	fx.Invoke(
		Entrypoint,
	),
)

type app struct {
	fx.In
	Lifecycle  fx.Lifecycle
	Handler    handler.Ihandler
	Middleware middleware.Imiddleware
}

func Entrypoint(a app) {

	server := mux.NewRouter()

	middlewares := []mux.MiddlewareFunc{
		a.Middleware.UserVerify,
		a.Middleware.HashVerify,
	}

	api := server.PathPrefix("/api").Subrouter()
	api.Use(middlewares...)
	api.HandleFunc("/check-acc.get", a.Handler.CheckAccount()).Methods("GET", "OPTIONS")
	api.HandleFunc("/balance.get", a.Handler.GetBalance()).Methods("GET", "OPTIONS")
	api.HandleFunc("/top-up.put", a.Handler.TopUp()).Methods("PUT", "OPTIONS")
	api.HandleFunc("/transactions.get", a.Handler.GetTransactions()).Methods("GET", "OPTIONS").
		Queries("from", "{from}").
		Queries("to", "{to}")

	srv := http.Server{
		Addr:    ":8070",
		Handler: server,
	}

	a.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go srv.ListenAndServe()

				ticker := time.NewTicker(20 * time.Second)

				go func() {
					a.Handler.TrnConfirmer(ticker)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return srv.Shutdown(ctx)
			},
		},
	)

}
