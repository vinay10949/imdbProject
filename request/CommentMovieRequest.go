package request

type CommentMovieRequest struct {
	MovieID int32 `json:"movieId"`
	UserID  int32 `json:"userId"`
	Comment  string `json:"comment"`
}