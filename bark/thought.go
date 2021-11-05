package bark

import "time"

type Thought struct {
	ThoughtID      string    `json:"thoughtId"`
	UserID         string    `json:"userId"`
	TopicID        string    `json:"topicId"`
	Text           string    `json:"thoughtText"`
	CreatedTime    time.Time `json:"createdTime"`
	LastEditedTime time.Time `json:"lastEditedTime"`
}
