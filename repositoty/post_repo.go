package repositoty

import (
	"backend-blogtechv2/model" 
	"context"
)

type PostRepo interface {
	SavePost(context context.Context, post model.Post, targetTable string) (model.Post, error)
	// GetPostByID(context context.Context, postID string) (model.Post, error)
	// GetAllPosts(context context.Context) ([]model.Post, error)
	// UpdatePost(context context.Context, post model.Post) (model.Post, error)
	// DeletePost(context context.Context, postID string) error
}
