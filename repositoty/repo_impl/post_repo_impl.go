package repoimpl

import (
	"backend-blogtechv2/bananaErr"
	"backend-blogtechv2/db"
	"backend-blogtechv2/log"
	"backend-blogtechv2/model"
	repositoty "backend-blogtechv2/repositoty"
	"context"
	"database/sql"
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

func (p *PostRepoImpl) getTableName(targetTable string) (string, error) {
	var tableName string
	switch targetTable {
	case "technews_posts":
		tableName = "technews_posts"
	case "future_technology_posts":
		tableName = "future_technology_posts"
	case "tutorials_and_tips_posts":
		tableName = "tutorials_and_tips_posts"
	default:
		return "", bananaErr.InvalidTable
	}
	return tableName, nil
}


func (p *PostRepoImpl)SavePost(context context.Context, post model.Post , targetTable string) (model.Post, error) {
	// Xác định bảng mục tiêu dựa vào giá trị targetTable.
	tableName, err := p.getTableName(targetTable)
	if err != nil {
		return post, err
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
	_, err = p.sql.Db.NamedExecContext(context, statement, post)

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

func (p *PostRepoImpl) GetPostByID(context context.Context, postID string, targetTable string) (model.Post, error) {
	var post model.Post

	// Xác định bảng dựa vào giá trị targetTable.
	tableName, err := p.getTableName(targetTable)
	if err != nil {
		return model.Post{}, err
	}

	// Xác định câu lệnh SQL để truy vấn thông tin của bài viết dựa vào postID.
	statement := `
        SELECT * FROM ` + tableName + ` WHERE pid = $1
    `

	// Thực hiện truy vấn và lưu kết quả vào biến post.
	err = p.sql.Db.GetContext(context, &post, statement, postID)
	if err != nil {
		log.Error(err.Error())
		if err == sql.ErrNoRows {
			return post, bananaErr.PostNotFound
		}
		return post, bananaErr.GetPostFail
	}

	return post, nil
}

func (p *PostRepoImpl) GetAllPostsByTable(context context.Context, targetTable string) ([]model.Post, error){
	var posts []model.Post

	// Xác định bảng dựa vào giá trị targetTable.
	tableName, err := p.getTableName(targetTable)
	if err != nil {
		return []model.Post{}, err
	}

	// Xác định câu lệnh SQL để truy vấn thông tin của tất cả bài viết trong bảng.
	statement := `
		SELECT * FROM ` + tableName + `
	`

	// Thực hiện truy vấn và lưu kết quả vào biến posts.
	err = p.sql.Db.SelectContext(context, &posts, statement)
	if err != nil {
		log.Error(err.Error())
		return posts, bananaErr.GetPostFail
	}

	return posts, nil
}

func (p *PostRepoImpl) UpdatePost(context context.Context, post model.Post, targetTable string) (model.Post, error) {
	// Xác định bảng mục tiêu dựa vào giá trị targetTable.
	tableName, err := p.getTableName(targetTable)
	if err != nil {
		return post, err
	}

	// Chuẩn bị câu lệnh SQL để cập nhật bài viết trong bảng mục tiêu.
	statement := `
        UPDATE ` + tableName + `
        SET 
            title = :title,
            summary = :summary,
            author = :author,
            content = :content,
            cover_image = :cover_image,
            content_image = :content_image,
            updated_at = :updated_at
        WHERE pid = :pid
    `

	// Đặt thời gian cập nhật cho bài viết.
	post.UpdatedAt = time.Now()

	// Thực hiện truy vấn UPDATE và kiểm tra lỗi.
	_, err = p.sql.Db.NamedExecContext(context, statement, post)

	if err != nil {
		log.Error(err.Error())
		return post, bananaErr.UpdatePostFail
	}

	return post, nil
}

func (p *PostRepoImpl) DeletePost(context context.Context, postID string, targetTable string) error{
	// Xác định bảng mục tiêu dựa vào giá trị targetTable.
    tableName, err := p.getTableName(targetTable)
    if err != nil {
        return err
    }

	// Chuẩn bị câu lệnh SQL để xóa bài viết từ bảng mục tiêu.
    statement := `
        DELETE FROM ` + tableName + ` WHERE pid = $1
    `
	// Thực hiện truy vấn DELETE và kiểm tra lỗi.
    _, err = p.sql.Db.ExecContext(context, statement, postID)
    if err != nil {
        log.Error(err.Error())
        return bananaErr.DeletePostFail
    }

    return nil
}