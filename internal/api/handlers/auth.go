package handlers

import (
	"errors"
	"net/http"

	"github.com/LekcRg/steam-inventory/internal/api"
	"github.com/LekcRg/steam-inventory/internal/errs"
	"go.uber.org/zap"
)

// AuthRedirect godoc
// @Summary      Redirect to Steam Auth
// @Description  Returns a 303 redirect to Steam authorization page.
// @Tags         Auth
// @Produce      plain
// @Success      303 {string} string "Redirect to Steam"
// @Failure      500 {object} models.ResponseError
// @Router       /auth [get]
//
// Returns a 303 redirect to Steam authorization page.
func (h *Handlers) AuthRedirect(w http.ResponseWriter, r *http.Request) {
	redirectURL, err := h.service.GetAuthRedirectURL()
	if err != nil {
		h.resp.InternalError(w)

		h.log.Error("Generating redirect URL error", zap.Error(err))
		return
	}

	http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
}

// AuthValid godoc
// @Summary      Steam auth callback
// @Description  Validates Steam auth response, sets session cookie and redirects to /me.
// @Tags         Auth
// @Produce      plain
// @Success      303 {string} string "Redirect to /me"
// @Header       303 {string} Location   "Redirect target (/me)"
// @Header       303 {string} Set-Cookie "sestoken=<token>; HttpOnly; Secure; Path=/; SameSite=Lax"
// @Failure      401 {object} models.ResponseError "Invalid auth"
// @Failure      500 {object} models.ResponseError "Internal error"
// @Router       /auth/valid [get]
//
// Validates Steam auth response, sets session cookie and redirects to /me.
func (h *Handlers) AuthValid(w http.ResponseWriter, r *http.Request) {
	_, session, err := h.service.AuthValid(r.Context(), r.URL.Query())
	if err != nil {
		if errors.Is(err, errs.ErrInvalidAuth) {
			h.resp.Error(w, http.StatusUnauthorized, "Invalid auth try again")

			return
		}

		h.resp.InternalError(w)
		h.log.Error("Auth validation error", zap.Error(err))

		return
	}

	cookies := http.Cookie{
		Name:     "sestoken",
		Value:    session,
		Path:     "/",
		MaxAge:   int(h.config.SessionExpire),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookies)

	http.Redirect(w, r, api.PathMe, http.StatusSeeOther)
}
