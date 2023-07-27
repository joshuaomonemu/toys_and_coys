package models

import (
	"app/database"
	"app/helper"
	"log"
)

var post = database.CreateClient().Collection("events")

type EventComment struct {
	Username string
	Comment  string
	Time     string
}
type Events struct {
	Username string
	Content  string
	Text     string
	Time     string
	Likes    int
	Comments EventComment
}

func CreateEvent(event *Events) bool {

	rands := helper.Randoms()
	post_val := "post-" + rands

	_, err := post.Doc(post_val).Create(ctx, event)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}
