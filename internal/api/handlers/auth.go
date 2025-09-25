package handlers

import (
	"errors"
	"net/http"

	"github.com/LekcRg/steam-inventory/internal/errs"
	"go.uber.org/zap"
)

func (h *Handlers) AuthRedirect(w http.ResponseWriter, r *http.Request) {
	redirectURL, err := h.service.GetAuthRedirectURL()
	if err != nil {
		h.resp.InternalError(w)

		h.log.Error("Generating redirect URL error", zap.Error(err))
		return
	}

	http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
}

func (h *Handlers) AuthValid(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.AuthValid(r.Context(), r.URL.Query())
	if err != nil {
		if errors.Is(err, errs.ErrInvalidAuth) {
			h.resp.Error(w, http.StatusUnauthorized, "Invalid auth try again")

			return
		}

		h.resp.InternalError(w)
		h.log.Error("Auth validation error", zap.Error(err))

		return
	}

	h.resp.JSON(w, http.StatusOK, user)
}
