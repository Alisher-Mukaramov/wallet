package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Alisher-Mukaramov/wallet/internal/models"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/types"
)

type AccountState struct {
	IsExist   int    `db:"isexist" json:"isExist"`
	IsActive  bool   `db:"isactive" json:"isActive"`
	Status    string `db:"status" json:"status"`
	SecretKey string `db:"secret" json:"secret,omitempty"`
}

func (a *AccountState) ClearSecret() {
	a.SecretKey = ""
}

func (repo repository) CheckAccount(ctx context.Context) (r models.Response, err error) {

	acc := AccountState{}

	userId := fmt.Sprintf("%v", ctx.Value(types.UserID))

	query := `SELECT count(1) isExist, u.is_active isActive, ut.code status, u.secret_key secret
	          FROM users u
              INNER JOIN user_types ut on ut.id = u.user_type_id
              WHERE u.id = %s 
              GROUP BY u.is_active, ut.code, u.secret_key`

	query = fmt.Sprintf(query, userId)

	err = repo.DBInstance().Get(&acc, query)

	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("CheckAccount() -> %w", NotExist)
		return
	}

	if err != nil {
		err = fmt.Errorf("CheckAccount() -> %w ::: %s", QueryError, err.Error())
		return
	}

	if !acc.IsActive {
		err = fmt.Errorf("%w", IsLocked)
		return
	}

	r.Set(acc)

	return
}
