package request

type AddMovieRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}