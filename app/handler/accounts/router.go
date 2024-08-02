package accounts

import (
	"net/http"

	"yatter-backend-go/app/app"

	"github.com/go-chi/chi"
)

// Implementation of handler
type handler struct {
	app *app.App
	dao *dao.Dao
}

// Create Handler for `/v1/accounts/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	h := &handler{app: app}
	r.Post("/", h.create)
	r.Get("/{username}", h.username)

	return r
}
