package controllers

import (
	entities "api/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}
func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusOK, t.tweets)
}
func (t *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, tweet := range t.tweets {
		if tweet.ID == id {
			t.tweets = append(t.tweets[0:idx], t.tweets[idx+1:]...)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"erro": "Tweet not found",
	})
}
