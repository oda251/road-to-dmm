package object

import (

)

type (
	StatusID = int64
	Status struct {
		ID StatusID `json:"-" db:"id"`
		accountId AccountID `json:"account_id" db:"account_id"`
		status string `json:"status" db:"status"`
		mediaIds int[] `json:"media_ids" db:"media_ids"`
		createAt DateTime `json:"create_at" db:"create_at"`
	}
)