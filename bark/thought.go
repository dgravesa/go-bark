package bark

import "time"

type Thought struct {
	UserID      string    `json:"userId"`
	Text        string    `json:"thoughtText"`
	CreatedTime time.Time `json:"createdTime"`
}
