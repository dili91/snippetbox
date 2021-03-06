package models

import (
	"errors"
	"time"
)

// ErrNoRecord : no record error
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet : model
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
