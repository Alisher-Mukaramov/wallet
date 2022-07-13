package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Alisher-Mukaramov/wallet/internal/models"
)

type UserAccInfo struct {
	AccountId         int     `db:"account_id"`
	Saldo             float32 `db:"saldo"`
	UnconfirmedAmount float32 `db:"unconfirmed_amount"`
	MaxLimitSaldo     float32 `db:"max_limit_saldo"`
	CurrencyId        int     `db:"currency_id"`
}

func (repo repository) GetAccInfo(userId, currency string) (r models.Response, err error) {

	accInfo := &UserAccInfo{}

	query := `SELECT acc.id as account_id, 
					 (SELECT COALESCE(SUM(trn.amount),0) 
					              FROM transactions trn 
					              WHERE trn.state_id = 1 and currency_id = c.id) as unconfirmed_amount, 
                     l.max_balance as max_limit_saldo,
                     acc.saldo     as saldo,
                     c.id          as currency_id
              FROM users u 
			  INNER JOIN accounts acc on acc.user_id = u.id
			  INNER JOIN currencies c on acc.currency_id = c.id
			  INNER JOIN account_types at on acc.acc_type_id = at.id
			  INNER JOIN user_types ut on u.user_type_id = ut.id
			  INNER JOIN limits l on l.user_type_id = ut.id
			  WHERE u.id = %s and c.code = '%s'`

	query = fmt.Sprintf(query, userId, currency)

	err = repo.DBInstance().Get(accInfo, query)

	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("%s %w", currency, AccountNotFound)
		return
	}

	if err != nil {
		err = fmt.Errorf("GetAccInfo() -> %w ::: %s", QueryError, err.Error())
		return
	}

	r.Set(*accInfo)

	return
}
