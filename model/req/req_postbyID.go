package req

type ReqPostID struct {
	PostID string `json:"post_id,omitempty" validate:"required"`
	SelectedLocation string `json:"selected_location,omitempty" validate:"required"` 
}
