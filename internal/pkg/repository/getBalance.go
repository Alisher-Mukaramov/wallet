package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Alisher-Mukaramov/wallet/internal/models"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/types"
)

type Balance struct {
	Account           string  `db:"account" json:"account"`
	Saldo             float32 `db:"saldo" json:"saldo"`
	UnconfirmedAmount float32 `db:"unconfirmed_amount" json:"unconfirmedAmount"`
	Currency          string  `db:"currency" json:"currency"`
}

func (repo repository) GetBalance(ctx context.Context) (r models.Response, err error) {

	balances := &[]Balance{}

	userId := fmt.Sprintf("%v", ctx.Value(types.UserID))

	query := `SELECT acc.account as account, 
					 (SELECT COALESCE(SUM(trn.amount),0) 
					              FROM transactions trn 
					              WHERE trn.state_id = 1 and currency_id = c.id) as unconfirmed_amount, 
                     acc.saldo     as saldo,
                     c.code          as currency
              FROM users u 
			  INNER JOIN accounts acc on acc.user_id = u.id
			  INNER JOIN currencies c on acc.currency_id = c.id
			  INNER JOIN account_types at on at.id = acc.acc_type_id
			  INNER JOIN user_types ut on u.user_type_id = ut.id
			  WHERE u.id = %s`

	query = fmt.Sprintf(query, userId)

	err = repo.DBInstance().Select(balances, query)

	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("GetBalance() -> %w", NoRows)
		return
	}

	if err != nil {
		err = fmt.Errorf("GetBalance() -> %w ::: %s", QueryError, err.Error())
		return
	}

	if len(*balances) == 0 {
		err = fmt.Errorf("GetBalance() -> %w", HasNoAccount)
		return
	}

	r.Set(balances)

	return
}
