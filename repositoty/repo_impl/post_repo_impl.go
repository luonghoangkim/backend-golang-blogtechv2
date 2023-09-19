package repoimpl

import (
	"backend-blogtechv2/bananaErr"
	"backend-blogtechv2/db"
	"backend-blogtechv2/log"
	"backend-blogtechv2/model"
	repositoty "backend-blogtechv2/repositoty"
	"context"
	"time"

	"github.com/lib/pq"
)

type PostRepoImpl struct {
	sql *db.Sql
}

func NewPostRepo(sql *db.Sql) repositoty.PostRepo {
	return &PostRepoImpl{
		sql: sql,
	}
}

func (p *PostRepoImpl)SavePost(context context.Context, post model.Post , targetTable string) (model.Post, error) {
	// Xác định bảng mục tiêu dựa vào giá trị targetTable.
    var tableName string
    switch targetTable {
    case "technews_posts":
        tableName = "technews_posts"
    case "future_technology_posts":
        tableName = "future_technology_posts"
    case "tutorials_and_tips_posts":
        tableName = "tutorials_and_tips_posts"
    default:
        return post, bananaErr.InvalidTable
    }

    // Chuẩn bị câu lệnh SQL để thêm bài viết mới vào bảng mục tiêu.
    statement := `
        INSERT INTO ` + tableName + ` (pid, title, summary, author, content, cover_image, content_image, created_at, updated_at)
        VALUES (:pid, :title, :summary, :author, :content, :cover_image, :content_image, :created_at, :updated_at)
    `


	// Đặt thời gian tạo và cập nhật cho bài viết mới.
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	// Thực hiện truy vấn INSERT và kiểm tra lỗi.
	_, err := p.sql.Db.NamedExecContext(context, statement, post)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return post, bananaErr.PostConflict
			}
		}
		return post, bananaErr.SavePostFail
	}

	return post, nil
}