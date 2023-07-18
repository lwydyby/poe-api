package poe_api

import (
	"fmt"
)

type InvalidToken struct {
	token string
}

func (e *InvalidToken) Error() string {
	return fmt.Sprintf("Invalid token: %s", e.token)
}
