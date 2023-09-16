package handler

import (
	"backend-blogtechv2/log"
	"backend-blogtechv2/model"
	"backend-blogtechv2/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	//UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSignUp{}

	// Kiểm tra và liên kết dữ liệu từ yêu cầu đến biến req.
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       "",
	})
}