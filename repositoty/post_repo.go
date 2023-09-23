package repositoty

import (
	"backend-blogtechv2/model" 
	"context"
)

type PostRepo interface {
	SavePost(context context.Context, post model.Post, targetTable string) (model.Post, error)
	GetPostByID(context context.Context, postID string, targetTable string) (model.Post, error)
	GetAllPostsByTable(context context.Context, targetTable string) ([]model.Post, error)
	UpdatePost(context context.Context, post model.Post ,  targetTable string) (model.Post, error)
	DeletePost(context context.Context, postID string, targetTable string) error
}
