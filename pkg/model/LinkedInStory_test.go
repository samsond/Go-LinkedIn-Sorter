package model

import (
	"testing"
	"time"
)

func TestLinkedInStoryConstruction(t *testing.T) {
	expectedDate := time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)
	story := LinkedInStory{
		Title:      "The Future of Remote Work",
		PostedDate: expectedDate,
		Views:      150,
		Likes:      50,
	}

	if story.Title != "The Future of Remote Work" {
		t.Errorf("Expected title 'Example Story', got %s", story.Title)
	}

	if !story.PostedDate.Equal(expectedDate) {
		t.Errorf("PostedDate is incorrect, got: %v, want: %v", story.PostedDate, expectedDate)
	}

	if story.Views != 150 {
		t.Errorf("Expected views for 'The Future of Remote Work', got %d", story.Views)
	}

	if story.Likes != 50 {
		t.Errorf("Expected likes for 'The Future of Remote Work', got %d", story.Likes)
	}

}
