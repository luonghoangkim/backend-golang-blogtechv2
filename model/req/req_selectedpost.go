package req

type ReqSelectedPost struct {
	SelectedLocation string `json:"selected_location,omitempty" validate:"required"` 
}
