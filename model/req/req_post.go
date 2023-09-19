package req

type ReqPost struct {
	SelectedLocation string `json:"selected_location,omitempty" validate:"required"`
	Title       string `json:"title,omitempty" validate:"required"`
	Summary     string `json:"summary,omitempty" validate:"required"`
	Author      string `json:"author,omitempty" validate:"required"`
	Content     string `json:"content,omitempty" validate:"required"`
	CoverImage  string `json:"coverImage,omitempty" validate:"required"`
	ContentImage string `json:"contentImage,omitempty" validate:"required"`
}
