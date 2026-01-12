package user

import (
	"time"
)

type User struct {
	ID        int
	Age       int
	Name      string
	Username  string
	CreatedAt time.Time
	Premium   bool
}
