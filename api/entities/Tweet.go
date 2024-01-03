package entities

import (
	"github.com/pborman/uuid"
)

type Tweet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New(),
	}

	return &tweet
}
