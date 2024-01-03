package sorter

import (
	"github.com/samsond/Go-LinkedIn-Sorter/pkg/model"
)

type ByViews []model.LinkedInStory

func (stories ByViews) Len() int {
	return len(stories)
}

func (stories ByViews) Swap(i, j int) {
	stories[i], stories[j] = stories[j], stories[i]
}

func (stories ByViews) Less(i, j int) bool {
	return stories[i].Views > stories[j].Views
}
