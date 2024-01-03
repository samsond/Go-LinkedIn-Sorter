package model

import "time"

type LinkedInStory struct {
	Title      string
	PostedDate time.Time
	Views      int
	Likes      int
}
