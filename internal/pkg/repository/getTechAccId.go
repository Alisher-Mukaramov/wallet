package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

func (repo repository) GetTechAccId(curId int) (accId int, err error) {

	query := `SELECT DISTINCT id as id
			  FROM accounts acc
              WHERE acc.acc_type_id = 2 AND acc.currency_id = %d`

	query = fmt.Sprintf(query, curId)

	err = repo.DBInstance().Get(&accId, query)

	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("GetTechAccId() -> %w", SysCurrencyError)
		return
	}

	if err != nil {
		err = fmt.Errorf("GetTechAccId() -> %w ::: %s", QueryError, err.Error())
		return
	}

	return
}
