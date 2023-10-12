package verify

import "time"

const (
	_IdentityStatusField = "identity_status"
	_WithdrawTotalField  = "withdraw_total"
	_UpdateTimeField     = "update_time"
)

type userVerify struct {
	ID               string    `pg:"id"`
	IdentityVerified bool      `pg:"identity_status"`
	WithdrawTotal    float32   `pg:"withdraw_total"`
	CreateTime       time.Time `pg:"create_time"`
	UpdateTime       time.Time `pg:"update_time"`
}

type identifyInfo struct {
	Type string
	ID   string
}
