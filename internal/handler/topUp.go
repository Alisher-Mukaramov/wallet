package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/repository"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/types"
	"net/http"
)

func (h handler) TopUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		topUp := &repository.TopUp{}

		err := json.NewDecoder(r.Body).Decode(topUp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), types.Request, *topUp)

		res, err := h.service.Repository().TopUp(ctx)
		if err != nil {
			if errors.Is(err, repository.QueryError) {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
				return
			} else if errors.Is(err, repository.BalanceLimitError) {
				http.Error(w, err.Error(), http.StatusNotAcceptable)
				return
			} else if errors.Is(err, repository.AccountNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		res.ToJson(w)
	}
}
