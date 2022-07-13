package middleware

import (
	"context"
	"errors"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/repository"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/types"
	"net/http"
)

var (
	userIsEmpty = errors.New("X-UserId не может быть пустым")
)

func (s *middleware) UserVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userID := r.Header.Get(XuserId)
		if userID == "" {
			http.Error(w, userIsEmpty.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), types.UserID, userID)

		userInfo, err := s.service.Repository().CheckAccount(ctx)

		if err != nil {
			if errors.Is(err, repository.NotExist) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			} else if errors.Is(err, repository.IsLocked) {
				http.Error(w, err.Error(), http.StatusLocked)
				return
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		ctx = context.WithValue(ctx, types.UserInfo, userInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
