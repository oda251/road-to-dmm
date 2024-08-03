package statuses

import (

)

type (
	postRequest struct {
		status string `json:"status"`
		mediaIds []int `json:"media_ids"`
	}
)

func (h *handler) post(w http.ResponseWriter, r *http.Request) {

}