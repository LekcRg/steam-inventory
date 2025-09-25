package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/LekcRg/steam-inventory/internal/api/middlewares"
	"go.uber.org/zap"
)

func (h *Handlers) UserInfo(w http.ResponseWriter, r *http.Request) {
	steamID, ok := r.Context().Value(middlewares.CtxKeySteamID).(string)
	if !ok {
		h.resp.Error(w, http.StatusInternalServerError, "Failed to get Steam ID from context")
		return
	}

	user, err := h.service.UserInfo(r.Context(), steamID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// or forbidden
			h.resp.Error(w, http.StatusNotFound, "User not found")
			return
		}

		h.resp.Error(w, http.StatusInternalServerError, "Failed to get user info")
		h.log.Error("Get user info failed", zap.Error(err))
		return
	}

	h.resp.JSON(w, http.StatusOK, user)
}
