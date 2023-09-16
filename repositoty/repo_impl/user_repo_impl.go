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

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repositoty.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

// SaveUser thêm một người dùng mới vào cơ sở dữ liệu.
// Nếu có lỗi xảy ra, nó xử lý và trả về lỗi tương ứng.
func (u *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	// Chuẩn bị câu lệnh SQL để thêm người dùng mới vào bảng "users".
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name , created_at , updated_at )
		VALUES(:user_id, :email, :password, :role, :full_name , :created_at , :updated_at)
	`

	// Đặt thời gian tạo và cập nhật cho người dùng mới.
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Thực hiện truy vấn INSERT và kiểm tra lỗi.
	_, err := u.sql.Db.NamedExecContext(context, statement, user)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, bananaErr.UserConflict
			}
		}
		return user, bananaErr.SignUpFail
	}
	return user, nil
}