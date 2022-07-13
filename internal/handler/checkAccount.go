package handler

import (
	"github.com/Alisher-Mukaramov/wallet/internal/models"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/types"
	"net/http"
)

func (h handler) CheckAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acc := r.Context().Value(types.UserInfo).(models.Response)
		acc.ToJson(w)
	}
}
