package sorter

import (
	"github.com/samsond/Go-LinkedIn-Sorter/pkg/model"
	"sort"
	"testing"
	"time"
)

func TestSortByViews(t *testing.T) {

	stories := []model.LinkedInStory{
		{Title: "The Future of Remote Work", PostedDate: time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC), Views: 150},
		{Title: "Breaking Into Tech: A Guide", PostedDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Views: 200},
		{Title: "Top 10 Programming Languages in 2024", PostedDate: time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC), Views: 350},
	}

	expected := []model.LinkedInStory{
		{Title: "Top 10 Programming Languages in 2024", PostedDate: time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC), Views: 350},
		{Title: "Breaking Into Tech: A Guide", PostedDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), Views: 200},
		{Title: "The Future of Remote Work", PostedDate: time.Date(2023, time.January, 2, 0, 0, 0, 0, time.UTC), Views: 150},
	}

	sort.Sort(ByViews(stories))

	for i, story := range stories {
		if story.Views != expected[i].Views {
			t.Errorf("Test failed at index %d: got %v, want %v", i, story, expected[i])
		}
	}
}
