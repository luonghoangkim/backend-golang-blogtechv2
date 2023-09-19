package repoimpl

import (
	"backend-blogtechv2/bananaErr"
	"backend-blogtechv2/db"
	"backend-blogtechv2/log"
	"backend-blogtechv2/model"
	"backend-blogtechv2/model/req"
	repositoty "backend-blogtechv2/repositoty"
	"context"
	"database/sql"
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

func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq  req.ReqSignIn) (model.User , error){
	var user = model.User{}
	err := u.sql.Db.GetContext(context , &user ,"SELECT * FROM users WHERE email=$1", loginReq.Email)

	if err != nil { 
		if err == sql.ErrNoRows {
			return user , bananaErr.UserNotFound
		}
		log.Error(err.Error())
		return user , err
	}
	return user, nil
}

func (u *UserRepoImpl)  SelectUserByID(context context.Context, userId string) (model.User , error) {
	var user model.User

	err := u.sql.Db.GetContext(context, &user,
		"SELECT * FROM users WHERE user_id = $1", userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, bananaErr.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {
	sqlStatement := `
		UPDATE users
		SET 
			full_name  = (CASE WHEN LENGTH(:full_name) = 0 THEN full_name ELSE :full_name END),
			email = (CASE WHEN LENGTH(:email) = 0 THEN email ELSE :email END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE user_id    = :user_id
	`

	user.UpdatedAt = time.Now()

	result, err := u.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}
 
	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, bananaErr.UserNotUpdated
	}
	if count == 0 {
		return user, bananaErr.UserNotUpdated
	}

	return user, nil
}