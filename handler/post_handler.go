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

// PostHandler quản lý các yêu cầu liên quan đến bài viết.
type PostHandler struct {
	PostRepo repositoty.PostRepo
}

// HandlePost xử lý yêu cầu tạo bài viết mới.
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

	// Tạo một ID duy nhất cho bài viết mới
	postID, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Tạo một bài viết mới từ dữ liệu yêu cầu
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

func (p *PostHandler) GetPostByID(c echo.Context) error {
	req := req.ReqPostID{}
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

	// Gọi phương thức GetPostByID từ PostRepo để lấy thông tin bài viết theo ID.
	post, err := p.PostRepo.GetPostByID(c.Request().Context(), req.PostID, req.SelectedLocation)
	if err != nil {
		// Xử lý lỗi ở đây nếu cần.
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy thông tin bài viết thành công",
		Data:       post,
	})
}


func (p *PostHandler) GetAllPostsByTable(c echo.Context) error {
	req := req.ReqSelectedPost{}
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

	// Gọi phương thức GetAllPostsByTable từ PostRepo để lấy danh sách bài viết.
	articles, err := p.PostRepo.GetAllPostsByTable(c.Request().Context(), req.SelectedLocation)

	if err != nil {
		// Xử lý lỗi ở đây nếu cần.
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy danh sách bài viết thành công",
		Data:       articles,
	})
}

func (p *PostHandler) UpdatePost(c echo.Context) error {
	req := req.ReqUpdatePost{}
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

	// Tạo một bài viết từ dữ liệu yêu cầu
	post := model.Post{
		PID:          req.PID,
		Title:        req.Title,
		Summary:      req.Summary,
		Author:       req.Author,
		Content:      req.Content,
		CoverImage:   req.CoverImage,
		ContentImage: req.ContentImage,
		UpdatedAt:    time.Now(),
	}

	// Cập nhật thông tin bài viết trong cơ sở dữ liệu
	updatedPost, err := p.PostRepo.UpdatePost(c.Request().Context(), post, req.SelectedLocation)
	if err != nil {
		// Xử lý lỗi ở đây nếu cần.
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Cập nhật bài viết thành công",
		Data:       updatedPost,
	})
}

func (p *PostHandler) DeletePost(c echo.Context) error {
    req := req.ReqPostID{}
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

    // Gọi phương thức DeletePost từ PostRepo để xóa bài viết.
    err := p.PostRepo.DeletePost(c.Request().Context(), req.PostID, req.SelectedLocation)
    if err != nil {
        // Xử lý lỗi ở đây nếu cần.
        return c.JSON(http.StatusInternalServerError, model.Response{
            StatusCode: http.StatusInternalServerError,
            Message:    err.Error(),
            Data:       nil,
        })
    }

    return c.JSON(http.StatusOK, model.Response{
        StatusCode: http.StatusOK,
        Message:    "Xóa bài viết thành công",
        Data:       nil,
    })
}
