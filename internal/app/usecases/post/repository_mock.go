package post

import (
	"github.com/marcelocquadros/blog/internal/app/domain/post"
	"github.com/stretchr/testify/mock"
)

type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) CreatePost(p *post.Post) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockPostRepository) UpdatePost(p *post.Post) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockPostRepository) DeletePost(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPostRepository) FindAllPosts() ([]post.Post, error) {
	args := m.Called()
	return args.Get(0).([]post.Post), args.Error(1)
}

func (m *MockPostRepository) FindPostByID(id string) (*post.Post, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*post.Post), args.Error(1)
}
