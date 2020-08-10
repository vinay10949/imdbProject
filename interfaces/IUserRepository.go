package interfaces

import "imdbProject/models"
type IUserRepository interface {
	//For Adding users
	AddMovie(name string,description string ) (models.SuccessModel, error)
	RateMovie(movieId int32,userId int32, rate int) (models.SuccessModel, error)
	CommentMovie(movieId int32,userId int32,comment string) (models.SuccessModel, error)
	ListMoviesRatedCommented(userId int32) ([]models.UserRatingCommentsRespModel, error)
}
