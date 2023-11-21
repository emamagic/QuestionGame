package param

import (
	"game/domain"
	"time"
)

type AddToWaitingListRequest struct {
	UserID   uint `json:"user_id"`
	Category domain.Category
}

type AddToWaitingListResponse struct {
	Timeout time.Duration `json:"timeout_in_nanoseconds"`
}
