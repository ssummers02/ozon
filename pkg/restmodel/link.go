package restmodel

import "time"

type Link struct {
	ID        int64
	Link      string `json:"link"`
	ShortLink string `json:"short_link"`
	CreatedAt time.Time
}
