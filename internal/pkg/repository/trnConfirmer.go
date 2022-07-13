package repository

func (repo repository) TrnConfirmer() {

	query := `UPDATE transactions trn
        	  SET processed_date=now(), state_id=2
              WHERE trn.state_id = 1`

	tx := repo.DBInstance().MustBegin()
	tx.MustExec(query)
	tx.Commit()

}
