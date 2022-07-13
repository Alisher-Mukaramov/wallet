package main

import (
	"github.com/Alisher-Mukaramov/wallet/internal"
	"github.com/Alisher-Mukaramov/wallet/internal/endpoint"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		internal.Modules,
		endpoint.Module,
	)
	app.Run()
}
