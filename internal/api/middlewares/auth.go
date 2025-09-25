package middlewares

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

type ctxKey string

const CtxKeySteamID ctxKey = "steamid"

func (m *Middlewares) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sestoken, err := r.Cookie("sestoken")
		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				m.log.Error("Middleware auth error", zap.Error(err))
			}

			m.resp.Error(w, http.StatusForbidden, "Forbidden")

			return
		}

		steamID, err := m.cache.GetSession(r.Context(), sestoken.Value)
		if err != nil {
			m.log.Error("Middleware auth error", zap.Error(err))
			m.resp.Error(w, http.StatusForbidden, "Forbidden")

			return
		}

		m.log.Info("auth",
			zap.String("sestoken", sestoken.Value),
			zap.String("steamid", steamID),
		)

		ctx := context.WithValue(r.Context(), CtxKeySteamID, steamID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
