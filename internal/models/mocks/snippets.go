package mocks

import (
	"time"

	"github.com/MohitPanchariya/Snippet-Box/internal/models"
)

var mockSnippet = &models.Snippet{
	Id:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

// This method mocks the models.SnippetModel.Insert method
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 2, nil
}

// This methods mocks the models.SnippetModel.Get method
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// This method mocks the models.SnippetModel.Latest method
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
