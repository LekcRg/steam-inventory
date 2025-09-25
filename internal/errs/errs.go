package errs

import "errors"

var ErrNothingMerge = errors.New("nothing to merge")

var (
	ErrAuthValidationReq   = errors.New("auth validation request error")
	ErrInvalidAuth         = errors.New("invalid authentication")
	ErrInvalidSteamID      = errors.New("invalid steam id")
	ErrNotFoundUserSummary = errors.New("user summary not found")
	ErrNotFoundUser        = errors.New("user not found")
)
