package request

type RateMovieRequest struct {
	MovieID int32 `json:"movieId"`
	UserID  int32 `json:"userId"`
	Rating  int `json:"rating"`
}