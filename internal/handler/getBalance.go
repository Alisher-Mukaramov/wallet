package handler

import (
	"errors"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/repository"
	"net/http"
)

func (h handler) GetBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		res, err := h.service.Repository().GetBalance(ctx)
		if err != nil {
			if errors.Is(err, repository.QueryError) {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
				return
			} else if errors.Is(err, repository.HasNoAccount) {
				http.Error(w, repository.HasNoAccount.Error(), http.StatusNotFound)
				return
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		res.ToJson(w)
	}
}
