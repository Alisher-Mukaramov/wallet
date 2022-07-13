package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Alisher-Mukaramov/wallet/internal/models"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/repository"
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/types"
	"net/http"
)

func (s *middleware) HashVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userId := r.Header.Get(XuserId)
		xDigest := r.Header.Get(Xdigest)

		resp := r.Context().Value(types.UserInfo).(models.Response)
		accState := resp.Get().(repository.AccountState)

		if Sha256(userId, []byte(accState.SecretKey)) != xDigest {
			http.Error(w, invalidHash.Error(), http.StatusUnauthorized)
			return
		}

		accState.ClearSecret()

		resp.Set(accState)

		ctx := context.WithValue(r.Context(), types.UserInfo, resp)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Sha256(text string, secret []byte) string {

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(text))
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}
