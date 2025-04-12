package http

import "time"

type Post struct {
	ID       int
	Text     string
	Author   string
	Comments int
	Time     time.Time
}
