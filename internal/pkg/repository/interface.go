package repository

import (
	"context"
	"errors"
	"github.com/Alisher-Mukaramov/wallet/internal/models"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/db"
	"go.uber.org/fx"
)

var (
	QueryError        = errors.New("Ошибка запроса")
	IsLocked          = errors.New("Пользователь заблокирован")
	NotExist          = errors.New("Пользователь не найден")
	AccountNotFound   = errors.New("Аккаунт не найден")
	HasNoAccount      = errors.New("У пользователя нет аккаунта")
	BalanceLimitError = errors.New("Данная сумма превышает допустимого лимита по балансу")
	SysCurrencyError  = errors.New("Кошелек не может работать с данной валютой")
	NoRows            = errors.New("Нет данных")
	DateFormatError   = errors.New("Некорректный формат даты")
	Module            = fx.Provide(newRepository)
)

type Irepository interface {
	CheckAccount(ctx context.Context) (resp models.Response, err error)
	TopUp(ctx context.Context) (r models.Response, err error)
	GetBalance(ctx context.Context) (r models.Response, err error)
	GetTransactions(ctx context.Context) (r models.Response, err error)
	TrnConfirmer()
}

type repoParams struct {
	fx.In
	db.Idb
}

type repository struct {
	db.Idb
}

func newRepository(rp repoParams) Irepository {
	return repository{rp.Idb}
}
