package base

import t "time"

type Model struct {
	ID          uint64
	CreatedAt   t.Time
	UpdatedAt   t.Time
	hiddenField bool
}
