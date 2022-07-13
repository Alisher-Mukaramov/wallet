package internal

import (
	"github.com/Alisher-Mukaramov/wallet/internal/handler"
	"github.com/Alisher-Mukaramov/wallet/internal/middleware"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/db"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/repository"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	db.Module,
	service.Module,
	repository.Module,
	handler.Module,
	middleware.Module,
)
