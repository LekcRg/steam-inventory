package handlers

import (
	"net/http"

	"github.com/LekcRg/steam-inventory/internal/auth"
	"go.uber.org/zap"
)

func (h *Handlers) AuthRedirect(w http.ResponseWriter, r *http.Request) {
	redirectURL, err := auth.GetRedirectURL(h.config.Domain)
	if err != nil {
		h.resp.InternalError(w)

		h.log.Error("Generating redirect URL error", zap.Error(err))
		return
	}

	http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
}

func (h *Handlers) AuthValid(w http.ResponseWriter, r *http.Request) {
	validURL, err := auth.GetValidURL(h.config.Domain, r.URL.Query())
	if err != nil {
		h.resp.InternalError(w)

		h.log.Error("Generating valid URL error", zap.Error(err))
		return
	}

	h.log.Info("url", zap.String("URL", validURL.String()))
	query := r.URL.Query()
	query.Add("URL", validURL.String())
	h.resp.JSON(w, http.StatusOK, query)
	// h.resp.Message(w, http.StatusOK, validURL.String())
}
