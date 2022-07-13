package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Alisher-Mukaramov/wallet/internal/models"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/types"
	"regexp"
)

type Statement struct {
	Date     string `db:"date" json:"date"`
	Currency string `db:"currency" json:"currency"`
	Amount   string `db:"amount" json:"amount"`
	TrnType  string `db:"trn_type" json:"trnType"`
	Status   string `db:"status" json:"status"`
}

type StatementRequest struct {
	DateFrom string
	DateTo   string
	regexp   *regexp.Regexp
}

func (sr StatementRequest) Validate() bool {
	sr.regexp = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if sr.regexp.MatchString(sr.DateFrom) &&
		sr.regexp.MatchString(sr.DateTo) {
		return true
	}
	return false
}

func (repo repository) GetTransactions(ctx context.Context) (r models.Response, err error) {

	statements := &[]Statement{}

	userId := fmt.Sprintf("%v", ctx.Value(types.UserID))

	statementReq := ctx.Value(types.Request).(StatementRequest)

	if !statementReq.Validate() {
		err = fmt.Errorf("GetBalance() -> %w", DateFormatError)
		return
	}

	query := `SELECT trn.insert_date as date,
       				 c.code as currency,
       				 trn.amount as amount,
       				 tt.name as trn_type,
       				 s.code as status
	          FROM transactions trn
    		  INNER JOIN transaction_types tt on trn.trn_type_id = tt.id
    		  INNER JOIN currencies c on trn.currency_id = c.id
    		  INNER JOIN states s on trn.state_id = s.id
    		  INNER JOIN accounts acc on acc.id = trn.debit_acc_id
              INNER JOIN users u on u.id = acc.user_id
    		  WHERE trn.insert_date between '%s' and '%s' and u.id = %s`

	query = fmt.Sprintf(query, statementReq.DateFrom, statementReq.DateTo, userId)

	err = repo.DBInstance().Select(statements, query)

	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("GetBalance() -> %w", NoRows)
		return
	}

	if err != nil {
		err = fmt.Errorf("GetBalance() -> %w ::: %s", QueryError, err.Error())
		return
	}

	r.Set(statements)

	return
}
