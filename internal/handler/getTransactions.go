package handler

import (
	"context"
	"errors"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/repository"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/types"
	"github.com/gorilla/mux"
	"net/http"
)

func (h handler) GetTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sreq := repository.StatementRequest{}

		sreq.DateFrom = mux.Vars(r)["from"]
		sreq.DateTo = mux.Vars(r)["to"]

		ctx := context.WithValue(r.Context(), types.Request, sreq)

		res, err := h.service.Repository().GetTransactions(ctx)
		if err != nil {
			if errors.Is(err, repository.DateFormatError) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		res.ToJson(w)
	}
}
