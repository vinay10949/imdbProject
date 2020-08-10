package viewModels


type MovieModel struct {
	Data struct {
		ID          int64 `sql:"primary_key"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Rating      int32  `json:"rating,omitempty"`
	} `json:"data"`

	Status bool `json:"status"`
	ErrMsg   string `json:"errMessage"`
}