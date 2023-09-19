package handler

import (
	"backend-blogtechv2/log"
	"backend-blogtechv2/model"
	"backend-blogtechv2/model/req"
	"backend-blogtechv2/repositoty"
	"net/http"
	"time"

	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	PostRepo repositoty.PostRepo
}

func (p *PostHandler) HandlePost(c echo.Context) error {

	req := req.ReqPost{}
	// Kiểm tra và liên kết dữ liệu từ yêu cầu đến biến req.
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Kiểm tra tính hợp lệ của dữ liệu đầu vào
	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	postID, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	post := model.Post{
		PID:          postID.String(),
		Title:        req.Title,
		Summary:      req.Summary,
		Author:       req.Author,
		Content:      req.Content,
		CoverImage:   req.CoverImage,
		ContentImage: req.ContentImage,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Lưu thông tin bài viết vào cơ sở dữ liệu
	selectedLocation := req.SelectedLocation // Lấy giá trị nơi lưu từ yêu cầu
	post, err = p.PostRepo.SavePost(c.Request().Context(), post, selectedLocation)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       post,
	})
}


