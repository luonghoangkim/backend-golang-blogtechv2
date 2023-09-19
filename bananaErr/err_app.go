package bananaErr

import "errors"

var(
	//user
	UserConflict = errors.New("Người dùng đã tồn tại")
	SignUpFail = errors.New("Đăng ký thất bại")
	UserNotFound = errors.New("Người dùng không tồn tại")
	UserNotUpdated = errors.New("Người dùng chưa cập nhật")

	//post
	PostConflict = errors.New("Bài viết đã tồn tại")
	SavePostFail =  errors.New("Đăng bài thất bại")
	InvalidTable = errors.New("Bảng không hợp lệ")
)