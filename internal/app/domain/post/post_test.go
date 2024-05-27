package post

import "testing"

func TestNewPost(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		tests := []struct {
			Name          string
			Title         string
			Content       string
			ImageUrl      string
			UserID        string
			ExpectedError error
		}{
			{"Invalid title", "", "A nice content", "https://image.com", "c93f0ec6-70af-4598-b767-dadd1554be3d", ErrInvalidTitle},
			{"Invalid content", "A nice title", "", "https://image.com", "c93f0ec6-70af-4598-b767-dadd1554be3d", ErrInvalidContent},
			{"Invalid image", "A nice title", "A nice content", "htt://image.com", "c93f0ec6-70af-4598-b767-dadd1554be3d", ErrInvalidImageURL},
			{"Invalid user id", "A nice title", "A nice content", "http://image.com", "c93f0ec6-70af-4598-b767-dad", ErrInvalidUserID},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				_, err := NewPost(tt.UserID, tt.Title, tt.Content, tt.ImageUrl)
				if err != tt.ExpectedError {
					t.Errorf("got %s, want %s", err, tt.ExpectedError)
				}
			})
		}
	})

}
