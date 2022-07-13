package repository

import (
	"context"
	"fmt"
	"github.com/Alisher-Mukaramov/wallet/internal/models"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/types"
)

const (
	stateId, trn_type = 1, 1 // хардкор
	success           = "Транзакция принята на обработку"
)

type TopUpResponse struct {
	Result string `db:"result"`
}

type Transaction struct {
	InsertDate  string  `db:"insert_date"`
	Amount      float32 `db:"amount"`
	DebitAccId  int     `db:"debit_acc_id"`
	CreditAccId int     `db:"credit_acc_id"`
	CurrencyId  int     `db:"currency_id"`
	TrnTypeId   int     `db:"trn_type_id"`
	StateId     int     `db:"state_id"`
}

type TopUp struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

func (repo repository) TopUp(ctx context.Context) (r models.Response, err error) {

	trn := &Transaction{}
	response := TopUpResponse{}

	request := ctx.Value(types.Request).(TopUp)

	userId := fmt.Sprintf("%v", ctx.Value(types.UserID))

	userAccInfo, err := repo.GetAccInfo(userId, request.Currency)
	if err != nil {
		return
	}

	_userAccInfo := userAccInfo.Get().(UserAccInfo)

	techAccId, err := repo.GetTechAccId(_userAccInfo.CurrencyId)
	if err != nil {
		return
	}

	balance := _userAccInfo.Saldo + _userAccInfo.UnconfirmedAmount + request.Amount
	allowedAmount := _userAccInfo.MaxLimitSaldo - (_userAccInfo.Saldo + _userAccInfo.UnconfirmedAmount)

	if balance > _userAccInfo.MaxLimitSaldo {
		err = fmt.Errorf("%w. Доступный лимит для пополнения %.2f", BalanceLimitError, allowedAmount)
		return
	}

	trn.CreditAccId = techAccId
	trn.DebitAccId = _userAccInfo.AccountId
	trn.StateId = stateId
	trn.TrnTypeId = trn_type
	trn.Amount = request.Amount
	trn.CurrencyId = _userAccInfo.CurrencyId
	trn.InsertDate = "now()"

	query := `INSERT INTO transactions(
	          insert_date, amount, debit_acc_id, credit_acc_id, currency_id, trn_type_id, state_id)
	          VALUES (:insert_date, :amount, :debit_acc_id, :credit_acc_id, :currency_id, :trn_type_id, :state_id);`

	tx := repo.DBInstance().MustBegin()
	_, err = tx.NamedExec(query, trn)
	if err != nil {
		err = fmt.Errorf("TopUp() -> %w ::: %s", QueryError, err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()

	response.Result = success
	r.Set(response)

	return
}
