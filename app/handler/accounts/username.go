package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) username(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	accountRepo := h.app.Dao.Account()
	accountWithFollowInfo, err := accountRepo.FetchAccountWithFollowInfoByUsername(r.Context(), username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if accountWithFollowInfo == nil {
		httperror.NotFound(w, nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accountWithFollowInfo); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}