package handler

import (
	"backend-blogtechv2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	//UserRepo repository.UserRepo
}

func (u *PostHandler) HandlePost(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       "",
	})
}

